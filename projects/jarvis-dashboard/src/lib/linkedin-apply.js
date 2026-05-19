// LinkedIn Easy Apply orchestrator (Trabajo v2 Fase 5).
//
// Connects to pc-playwright MCP via the C1 tunnel (mcp-pc-relay:8765),
// drives the Easy Apply flow on the user's real Chrome browser on PC,
// uses claudemax to parse form snapshots and match fields against user_profile,
// and asks the user via the embedded chat in /trabajo when a field is unknown.
//
// State machine per application session (persisted to apply-sessions.jsonl):
//   initialising -> navigating -> opening_apply -> filling
//                -> awaiting_user_answer (pause for chat Q&A)
//                -> uploading_cv
//                -> gate2_review (pause for user submit confirmation)
//                -> submitting -> applied | failed
//
// Q&A: when a field can't be mapped from user_profile,
//   1) post message to vault/projects/trabajo chat with metadata.type='auto_apply_question'
//   2) state -> awaiting_user_answer
//   3) when user replies in chat (next user message), parse it as the answer
//   4) persist answer to user_profile.questions_answered[question_hash] for next time
//   5) state -> filling (continue)
//
// Gate 2: when form is ready to submit,
//   1) screenshot via browser_take_screenshot, save to vault/inbox/job-hunting/apply-shots/<sess>.png
//   2) post chat message "Formulario listo. ¿Envío? (responde 'sí' o 'no')"
//   3) state -> gate2_review
//   4) wait for user yes/no
//   5) if yes: browser_click submit; state -> submitting -> applied
//   6) if no: state -> cancelled_by_user

import fs from 'fs/promises';
import path from 'path';
import crypto from 'crypto';
import { McpClient } from './mcp-client.js';
import { readProfile, writeProfile } from './profile.js';
import { callLLM } from './llm-router.js';

const VAULT_ROOT = process.env.VAULT_ROOT || '/vault';
const TOKEN_PATH = process.env.C1_TOKEN_PATH || '/vault/secrets/c1-bridge-token.txt';
const PC_BRIDGE_URL = process.env.PC_BRIDGE_URL || 'http://mcp-pc-relay:8765/servers/playwright/mcp';
const SESSIONS_FILE = path.join(VAULT_ROOT, 'inbox', 'job-hunting', 'apply-sessions.jsonl');
const APPLIED_TODAY = path.join(VAULT_ROOT, 'inbox', 'job-hunting', 'applied-today.json');
const ERRORS_FILE = path.join(VAULT_ROOT, 'inbox', 'job-hunting', 'apply-errors.jsonl');
const SHOTS_DIR = path.join(VAULT_ROOT, 'inbox', 'job-hunting', 'apply-shots');
const CHAT_TRABAJO = path.join(VAULT_ROOT, 'projects', 'trabajo', 'inbox', 'chat.jsonl');
const CHAT_GLOBAL = path.join(VAULT_ROOT, 'inbox', 'chat.jsonl');
const DAILY_LIMIT = parseInt(process.env.LINKEDIN_DAILY_LIMIT || '10', 10);

// In-memory session registry: session_id -> { state, lastActivity, abort }
const _sessions = new Map();

async function loadToken() {
  try { return (await fs.readFile(TOKEN_PATH, 'utf8')).trim(); }
  catch { return null; }
}

async function rateLimitCheck() {
  const today = new Date().toISOString().slice(0, 10);
  let data = { date: today, count: 0 };
  try {
    const raw = await fs.readFile(APPLIED_TODAY, 'utf8');
    const parsed = JSON.parse(raw);
    if (parsed.date === today) data = parsed;
  } catch {}
  return { ok: data.count < DAILY_LIMIT, count: data.count, limit: DAILY_LIMIT, date: today };
}

async function rateLimitIncrement() {
  const today = new Date().toISOString().slice(0, 10);
  let data = { date: today, count: 0 };
  try {
    const raw = await fs.readFile(APPLIED_TODAY, 'utf8');
    const parsed = JSON.parse(raw);
    if (parsed.date === today) data = parsed;
  } catch {}
  data.count += 1;
  data.updated_at = new Date().toISOString();
  await fs.mkdir(path.dirname(APPLIED_TODAY), { recursive: true });
  await fs.writeFile(APPLIED_TODAY, JSON.stringify(data, null, 2), 'utf8');
  return data;
}

async function appendSession(rec) {
  await fs.mkdir(path.dirname(SESSIONS_FILE), { recursive: true });
  await fs.appendFile(SESSIONS_FILE, JSON.stringify(rec) + '\n', 'utf8');
}

async function logError(sessionId, stage, err) {
  await fs.mkdir(path.dirname(ERRORS_FILE), { recursive: true });
  await fs.appendFile(ERRORS_FILE, JSON.stringify({
    ts: new Date().toISOString(),
    session_id: sessionId,
    stage,
    error: String(err && err.message || err),
    stack: err && err.stack
  }) + '\n', 'utf8');
}

async function postChatMessage(text, metadata = {}) {
  // We append directly to project chat AND global chat (same pattern as server.js appendChatMessage).
  // The chat-processor watches the global file; UI filters by project param.
  const msg = {
    id: `auto-${Date.now()}-${crypto.randomBytes(3).toString('hex')}`,
    ts: new Date().toISOString(),
    from: 'job-hunter',
    role: 'assistant',
    text,
    project: 'trabajo',
    agent: 'auto-apply',
    metadata
  };
  await fs.mkdir(path.dirname(CHAT_TRABAJO), { recursive: true });
  await fs.appendFile(CHAT_TRABAJO, JSON.stringify(msg) + '\n', 'utf8');
  await fs.appendFile(CHAT_GLOBAL, JSON.stringify(msg) + '\n', 'utf8');
  return msg;
}

async function waitForUserReply(sinceIso, timeoutMs = 10 * 60 * 1000) {
  // Poll trabajo chat for a user message newer than sinceIso. Return the message.
  const deadline = Date.now() + timeoutMs;
  while (Date.now() < deadline) {
    try {
      const raw = await fs.readFile(CHAT_TRABAJO, 'utf8');
      const lines = raw.split('\n').filter(Boolean);
      for (let i = lines.length - 1; i >= 0; i--) {
        try {
          const m = JSON.parse(lines[i]);
          if (m.role !== 'user') continue;
          if (m.ts <= sinceIso) break;
          return m;
        } catch {}
      }
    } catch {}
    await new Promise(r => setTimeout(r, 2000));
  }
  return null; // timeout
}

// --- MCP wrapper for pc-playwright ---
class PlaywrightMCP {
  constructor() { this.client = null; this.tools = null; }
  async connect() {
    const token = await loadToken();
    if (!token) throw new Error('no C1 bridge token at ' + TOKEN_PATH);
    this.client = new McpClient({
      url: PC_BRIDGE_URL,
      headers: { 'Authorization': `Bearer ${token}` }
    });
    await this.client.initialize();
    this.tools = await this.client.listTools();
    return this.tools;
  }
  async call(name, args = {}) {
    return this.client.callTool(name, args);
  }
  async navigate(url) { return this.call('browser_navigate', { url }); }
  async snapshot() { return this.call('browser_snapshot'); }
  // NOTE: @playwright/mcp uses `target` (snapshot ref OR CSS selector), not `ref`.
  // We accept `ref` in our wrapper signature for readability and translate here.
  async click(element, ref) { return this.call('browser_click', { element, ref }); }
  async type(element, ref, text, slowly = false, submit = false) {
    return this.call('browser_type', { element, target: ref, text, slowly, submit });
  }
  async selectOption(element, ref, values) {
    return this.call('browser_select_option', { element, ref, values });
  }
  async fileUpload(paths) {
    return this.call('browser_file_upload', { paths });
  }
  async screenshot() {
    // browser_take_screenshot returns content as [text, image] where image has base64 data.
    // We deliberately do NOT pass `filename` because that path is interpreted on the PC
    // (Windows filesystem), not on the VPS. We grab the inline base64 and write it server-side.
    const r = await this.call('browser_take_screenshot', { type: 'png' });
    const content = (r && r.content) || [];
    for (const item of content) {
      if (item && item.type === 'image' && item.data) {
        return { mime: item.mimeType || 'image/png', dataB64: item.data };
      }
    }
    return { mime: 'image/png', dataB64: null };
  }
  async waitFor(opts) { return this.call('browser_wait_for', opts); }
  // Runs a JS function in the page. If `ref` is provided, @playwright/mcp resolves it
  // to a Locator and passes it as the first argument. This bypasses ref-staleness
  // because the locator is re-resolved at evaluation time.
  async evaluate(fnStr, ref, element) {
    const args = { function: fnStr };
    if (ref) args.target = ref;
    if (element) args.element = element;
    const r = await this.call('browser_evaluate', args);
    try {
      const txt = (r && r.content && r.content[0] && r.content[0].text) || '';
      const NL = String.fromCharCode(10);
      const RESULT_TAG = '### Result' + NL;
      const startPos = txt.indexOf(RESULT_TAG);
      let blob = txt;
      if (startPos >= 0) {
        const after = txt.slice(startPos + RESULT_TAG.length);
        const endPos = after.indexOf(NL + '###');
        blob = endPos >= 0 ? after.slice(0, endPos) : after;
      }
      blob = blob.trim();
      if (!blob) return null;
      return JSON.parse(blob);
    } catch { return null; }
  }
}

// --- LLM-driven field extraction ---
const FIELD_EXTRACT_PROMPT = `Te paso una representación textual de la pantalla del flujo Easy Apply de LinkedIn (formato MCP browser_snapshot YAML, en español). Tu trabajo: identificar los campos de formulario visibles dentro del modal/dialog Y los botones de avance.

REGLAS DE CLASIFICACIÓN DE BOTONES (orden de prioridad — solo UNO suele estar visible en cada paso):

1. **submit_button_ref** = el botón FINAL que envía la aplicación. Etiquetas típicas en LinkedIn ES/EN (aria-label O texto visible en \`generic\` hijo):
   "Enviar solicitud" · "Enviar aplicación" · "Submit application" · "Send application" · "Enviar"
   SOLO marca como submit si la etiqueta contiene "Enviar"/"Submit"/"Send" + alusión a aplicación/solicitud.

2. **review_button_ref** = botón que va a la pantalla final de revisión antes del submit. Etiquetas típicas (aria-label O texto en \`generic\` hijo):
   "Revisar" · "Review your application" · "Revisar tu solicitud" · "Ir a la pantalla de revisión"

3. **next_button_ref** = botón que avanza al SIGUIENTE paso del wizard (no es el final). En LinkedIn el aria-label suele ser distinto del texto visible — el TEXTO VISIBLE del botón está en un \`generic\` HIJO. Aria-labels reales observados:
   - \`button "Ir al siguiente paso"\` (con hijo \`generic: Siguiente\`) ← MUY COMÚN
   - \`button "Continue to next step"\` (hijo \`generic: Next\` o \`generic: Continue\`)
   - \`button "Continuar al siguiente paso"\` (hijo \`generic: Continuar\`)
   - \`button "Siguiente"\` directo (raro)
   - \`button "Next"\` directo (raro)

   REGLA: si dentro del dialog hay un \`button "..."\` cuyo aria-label O cuyo \`generic\` hijo inmediato contiene una de estas palabras (case-insensitive) **Siguiente · Continuar · Next · Continue · siguiente paso · next step**, ese es \`next_button_ref\`. CASI NUNCA debe ser null cuando hay un dialog abierto sin submit_button_ref ni review_button_ref.

   **NO confundir** con: \`button "Descartar"\` / \`button "Cancelar"\` / \`button "Cerrar"\` (esos son dismiss/cancel, déjalos fuera de cualquier ref).

4. **easy_apply_button_ref** = link/botón "Solicitud sencilla" / "Easy Apply" para ABRIR el modal — solo cuando el modal aún NO se ha abierto. Si el modal ya está abierto (hay un dialog visible), este debe ser null.

5. **file_upload_ref** = ref del control de subida de CV. Etiquetas típicas: "Subir currículum", "Upload resume", "CV*" con type file.

CAMPOS (fields): cada textbox, combobox, listbox, switch, checkbox dentro del modal. Para cada uno:
- label: el texto visible junto al campo (p.ej. "Email*", "Teléfono móvil*", "Código del país*")
- kind: text | email | tel | number | textarea | select | radio | checkbox | date | file | unknown
- current_value: si el snapshot expone un valor ya rellenado (p.ej. \`textbox "Email" "davidgnjunior@gmail.com"\`)
- options: lista de opciones literales si es select/radio/combobox

Devuelve SOLO este JSON sin texto antes ni después (sin fences markdown):

{
  "is_easy_apply": true|false,
  "easy_apply_button_ref": "string|null",
  "submit_button_ref": "string|null",
  "next_button_ref": "string|null",
  "review_button_ref": "string|null",
  "file_upload_ref": "string|null",
  "fields": [{"ref":"string","label":"string","kind":"...","required":true|false,"current_value":"string|null","options":[]}],
  "errors": ["string"],
  "phase_hint": "initial|filling|reviewing|submitted|other"
}

NO inventes refs. SOLO usa refs que aparezcan literalmente en el snapshot (formato \`[ref=eNNN]\` o \`[ref=NNN]\`).
Si una etiqueta de botón puede caer en varias categorías, aplica el orden de prioridad arriba (submit > review > next).
PRESTA ESPECIAL ATENCIÓN al final del snapshot — los botones del modal suelen estar al final del bloque del dialog.`;

async function extractFields(snapshotText) {
  const trimmed = (snapshotText || '').slice(0, 12000);
  const result = await callLLM({
    model: 'claudemax/claude-sonnet-4-6',
    system: FIELD_EXTRACT_PROMPT,
    messages: [{ role: 'user', content: trimmed }],
    maxTokens: 3000
  });
  const text = result.text || result.content || '';
  let jsonText = text.trim();
  const fence = jsonText.match(/```(?:json)?\s*([\s\S]*?)```/);
  if (fence) jsonText = fence[1].trim();
  try { return JSON.parse(jsonText); }
  catch {
    const s = jsonText.indexOf('{'), e = jsonText.lastIndexOf('}');
    if (s >= 0 && e > s) return JSON.parse(jsonText.slice(s, e + 1));
    throw new Error('field extractor returned non-JSON: ' + text.slice(0, 200));
  }
}

// --- Map fields against user profile ---

// Normalize label so "Código del país" and "Código del país *" hash to the same key.
function normalizeLabel(s) {
  return String(s || '')
    .toLowerCase()
    .replace(/\*+/g, '')        // strip required markers
    .replace(/[¿?¡!:]/g, '')    // strip ES/EN punctuation that varies between renders
    .replace(/\s+/g, ' ')      // collapse whitespace
    .trim();
}

function fieldHash(field) {
  return crypto.createHash('sha1')
    .update(`${normalizeLabel(field.label)}|${field.kind}|${(field.options || []).join(',')}`)
    .digest('hex').slice(0, 10);
}

// Set <input>/<textarea> value via the prototype's native setter (the one React/Vue
// listen to) and dispatch input+change+blur events. Bypasses LinkedIn input masks
// that swallow native keystroke events sent by browser_type.
async function setInputViaJS(mcp, label, ref, value) {
  const v = JSON.stringify(String(value == null ? '' : value));
  const fn = `(el) => {
    try {
      if (!el) return { ok: false, error: 'no element' };
      const tag = (el.tagName || '').toLowerCase();
      const proto = tag === 'textarea' ? window.HTMLTextAreaElement.prototype : window.HTMLInputElement.prototype;
      const desc = Object.getOwnPropertyDescriptor(proto, 'value');
      if (!desc || !desc.set) return { ok: false, error: 'no setter' };
      desc.set.call(el, ${v});
      el.dispatchEvent(new Event('input', { bubbles: true }));
      el.dispatchEvent(new Event('change', { bubbles: true }));
      el.dispatchEvent(new Event('blur', { bubbles: true }));
      return { ok: true, actualValue: el.value };
    } catch (e) { return { ok: false, error: String(e && e.message || e) }; }
  }`;
  return await mcp.evaluate(fn, ref, label);
}

// Read back the current value of an input/textarea via JS — used to verify fill success
// when browser_type returned OK but the value got swallowed (input masks etc.).
async function readbackValue(mcp, label, ref) {
  return await mcp.evaluate('(el) => (el && el.value) || null', ref, label);
}

// Deterministic DOM scan: list every required form element in the open modal
// that is currently empty. Bypasses the LLM extractor entirely — useful as a
// safety net at Gate 2 because the LLM sometimes misses fields that ARE in
// the snapshot.
async function scanRequiredEmptyFields(mcp) {
  const fn = `() => {
    const dialog = document.querySelector('div[role="dialog"]') || document.body;
    const sels = dialog.querySelectorAll('input, select, textarea');
    const out = [];
    for (const el of sels) {
      // Skip hidden / disabled
      if (el.type === 'hidden') continue;
      if (el.disabled) continue;
      // (Fase 5d.1: scan hidden screens too — LinkedIn Easy Apply hides inactive form pages)
      const required = el.required ||
        el.getAttribute('aria-required') === 'true' ||
        el.getAttribute('required') !== null;
      if (!required) continue;
      // Look up a human label for this element.
      let label = '';
      try {
        if (el.id) {
          const l = document.querySelector('label[for=\"' + el.id + '\"]');
          if (l) label = (l.innerText || l.textContent || '').trim();
        }
      } catch {}
      if (!label) label = el.getAttribute('aria-label') || el.getAttribute('placeholder') || el.name || '';
      const value = (el.value || '').trim();
      if (value !== '') continue;
      out.push({
        label: label.slice(0, 120),
        tag: el.tagName.toLowerCase(),
        type: el.type || '',
        id: el.id || '',
        name: el.name || ''
      });
    }
    return out;
  }`;
  const r = await mcp.evaluate(fn);
  return Array.isArray(r) ? r : [];
}

// Try to fill a single DOM-scanned field using the native value setter +
// React-compatible event dispatch. Returns true if value landed.
async function jsFillById(mcp, id, value) {
  const v = JSON.stringify(String(value == null ? '' : value));
  const idJs = JSON.stringify(String(id));
  const fn = `() => {
    try {
      const el = document.getElementById(${idJs});
      if (!el) return { ok: false, error: 'no element by id' };
      const tag = (el.tagName || '').toLowerCase();
      const proto = tag === 'textarea' ? window.HTMLTextAreaElement.prototype : window.HTMLInputElement.prototype;
      const desc = Object.getOwnPropertyDescriptor(proto, 'value');
      if (!desc || !desc.set) return { ok: false, error: 'no setter' };
      el.focus();
      desc.set.call(el, ${v});
      el.dispatchEvent(new Event('input', { bubbles: true }));
      el.dispatchEvent(new Event('change', { bubbles: true }));
      el.dispatchEvent(new Event('blur', { bubbles: true }));
      return { ok: true, actualValue: el.value };
    } catch (e) { return { ok: false, error: String(e && e.message || e) }; }
  }`;
  return await mcp.evaluate(fn);
}

function tryMatchFromProfile(field, profile) {
  const label = (field.label || '').toLowerCase();
  const p = profile.personal || {};
  const pref = profile.preferences || {};
  // Direct mappings
  if (/email|correo/.test(label)) return p.email;
  if (/phone|tel[eé]fono|móvil/.test(label)) return p.phone ? String(p.phone).replace(/[^0-9+]/g, '') : null;
  if (/full name|nombre completo|^nombre$/.test(label)) return p.full_name;
  if (/first name|nombre/.test(label)) return p.full_name && p.full_name.split(/\s+/)[0];
  if (/last name|apellid/.test(label)) return p.full_name && p.full_name.split(/\s+/).slice(1).join(' ');
  if (/city|ciudad|location|ubicaci[oó]n/.test(label)) return p.location;
  if (/linkedin/.test(label)) return p.linkedin_url;
  if (/github/.test(label)) return p.github_url;
  if (/portfolio|website|web/.test(label)) return p.portfolio_url;
  if (/salary|salario|expected compensation|pretensi[oó]n/.test(label)) {
    if (pref.salary_min_eur_year) return String(pref.salary_min_eur_year);
  }
  if (/work auth|authoriz|permiso|visa/.test(label)) return p.work_auth;
  // Stored Q&A answers from prior applications
  const h = fieldHash(field);
  if (profile.questions_answered && profile.questions_answered[h]) return profile.questions_answered[h].answer;
  return null;
}

async function persistAnswer(field, answer) {
  const profile = await readProfile();
  profile.questions_answered = profile.questions_answered || {};
  profile.questions_answered[fieldHash(field)] = {
    label: field.label,
    kind: field.kind,
    answer,
    saved_at: new Date().toISOString()
  };
  await writeProfile(profile);
}

// Poll snapshot+extractFields until form content appears (post-click on Easy Apply / Next / Review).
// LinkedIn modals can take 2-4s to render. Returns the first useful extract or {timeout:true}.
async function waitForApplyModalReady(mcp, { maxAttempts = 6, delayMs = 1500 } = {}) {
  let lastFields = null;
  let lastSnapText = '';
  for (let i = 0; i < maxAttempts; i++) {
    await new Promise(r => setTimeout(r, delayMs));
    const snap = await mcp.snapshot();
    const snapText = JSON.stringify(snap?.content || snap || '').slice(0, 18000);
    const fields = await extractFields(snapText);
    lastFields = fields;
    lastSnapText = snapText;
    const hasContent = !!(
      fields.submit_button_ref || fields.review_button_ref || fields.next_button_ref ||
      fields.file_upload_ref || (Array.isArray(fields.fields) && fields.fields.length > 0)
    );
    if (hasContent) return { fields, snapText, attempts: i + 1, timeout: false };
  }
  return { fields: lastFields, snapText: lastSnapText, attempts: maxAttempts, timeout: true };
}

// After clicking the Submit button, poll for any sign that LinkedIn accepted the application.
// Signals (any match passes):
//   - text contains "Solicitud enviada" / "Application sent" / "Tu solicitud se ha enviado" / "se ha enviado" / "Done\b"
//   - presence of a dialog with aria-label matching application-sent semantics
//   - presence of a "Listo"/"Hecho"/"Done" button (post-submit confirmation modal)
//   - URL no longer contains /apply/  (LinkedIn redirects on success)
async function waitForSubmitConfirmation(mcp, { maxAttempts = 8, delayMs = 1500 } = {}) {
  const RX = /Solicitud enviada|Application sent|Tu solicitud se ha enviado|Your application was sent|se ha enviado|application.*sent|aplicaci[oó]n.*enviad|Done\b|Listo\b|Hecho\b/i;
  for (let i = 0; i < maxAttempts; i++) {
    await new Promise(r => setTimeout(r, delayMs));
    let snap = null;
    try { snap = await mcp.snapshot(); } catch (_) { continue; }
    const snapText = JSON.stringify(snap?.content || snap || '').slice(0, 18000);
    if (RX.test(snapText)) return { ok: true, attempts: i + 1, signal: 'text_match' };
    // URL check via evaluate
    try {
      const url = await mcp.evaluate('() => location.href');
      if (url && typeof url === 'string' && !url.includes('/apply/')) {
        return { ok: true, attempts: i + 1, signal: 'url_change', url };
      }
    } catch (_) {}
  }
  return { ok: false, timeout: true };
}

// --- main orchestrator ---
export async function startApplySession({ offerId, offerUrl, offerTitle, company, cvPdfAbsPath }) {
  const sessionId = `apply-${Date.now()}-${crypto.randomBytes(3).toString('hex')}`;
  const session = {
    session_id: sessionId,
    offer_id: offerId,
    offer_url: offerUrl,
    offer_title: offerTitle,
    company,
    cv_pdf: cvPdfAbsPath,
    state: 'initialising',
    started_at: new Date().toISOString(),
    last_update: new Date().toISOString(),
    error: null,
    history: []
  };

  _sessions.set(sessionId, { state: 'initialising', record: session });
  await appendSession(session);

  // Fire-and-forget runner (background)
  runApplySession(sessionId, session).catch(async (e) => {
    await logError(sessionId, 'runner', e);
    session.state = 'failed';
    session.error = String(e.message || e);
    session.last_update = new Date().toISOString();
    await appendSession(session);
    _sessions.set(sessionId, { state: 'failed', record: session });
  });

  return session;
}

export function getSession(sessionId) {
  const s = _sessions.get(sessionId);
  return s ? s.record : null;
}

async function updateState(session, newState, extra = {}) {
  session.state = newState;
  session.last_update = new Date().toISOString();
  Object.assign(session, extra);
  session.history.push({ at: session.last_update, state: newState, ...extra });
  await appendSession(session);
  _sessions.set(session.session_id, { state: newState, record: session });
}

async function runApplySession(sessionId, session) {
  // 1. Rate limit gate
  const rl = await rateLimitCheck();
  if (!rl.ok) {
    await postChatMessage(`⛔ Cuota diaria alcanzada (${rl.count}/${rl.limit}). No se aplicará. Mañana se reinicia.`, { type: 'rate_limit_hit' });
    await updateState(session, 'rate_limited');
    return;
  }

  // 2. MCP connect
  await updateState(session, 'connecting');
  const mcp = new PlaywrightMCP();
  await mcp.connect();

  // 3. Navigate to offer
  await updateState(session, 'navigating');
  await mcp.navigate(session.offer_url);
  await new Promise(r => setTimeout(r, 1500));

  // 4. Snapshot + detect Easy Apply
  let snap = await mcp.snapshot();
  let snapText = JSON.stringify(snap?.content || snap || '').slice(0, 18000);
  let fields = await extractFields(snapText);

  if (!fields.is_easy_apply || !fields.easy_apply_button_ref) {
    await postChatMessage(`⚠️ La oferta «${session.offer_title}» no parece tener Easy Apply. Aborto auto-apply — tendrás que aplicar manualmente.`, { type: 'no_easy_apply' });
    await updateState(session, 'failed', { error: 'no_easy_apply' });
    return;
  }

  // 5. Open Easy Apply modal.
  //
  // EMPIRICAL: clicking the Solicitud sencilla / Easy Apply link via browser_click
  // does NOT reliably open the modal in the user's Chrome (probably because LinkedIn's
  // SPA router needs the URL navigation event to trigger the SDUI Apply flow).
  // Direct navigation to the apply URL works consistently.
  await updateState(session, 'opening_apply');
  const applyUrl = session.offer_url.replace(/\/+$/, '') + '/apply/?openSDUIApplyFlow=true';
  await mcp.navigate(applyUrl);

  // 5b. Wait for the Easy Apply modal to actually render — LinkedIn can take 2-4s.
  const modalWait = await waitForApplyModalReady(mcp);
  if (modalWait.timeout) {
    // Fallback: maybe the SDUI flow URL doesn't apply for this offer — try the
    // classic link click as a last resort.
    await mcp.click('Easy Apply button (fallback)', fields.easy_apply_button_ref);
    const modalWait2 = await waitForApplyModalReady(mcp);
    if (modalWait2.timeout) {
      await postChatMessage(`⚠️ El modal de Easy Apply no abrió para «${session.offer_title}» tras 2 intentos (navegación directa + click). Aborto — cuota no gastada.`, { type: 'modal_did_not_open' });
      await updateState(session, 'failed', { error: 'modal_did_not_open' });
      return;
    }
    Object.assign(modalWait, modalWait2);
  }

  // 6. Iterate form pages
  const profile = await readProfile();
  let safetyMax = 8; // max pages in the Easy Apply wizard
  let filledAnyField = false;
  const fillFailures = [];   // Fase 5c: track text fields that neither type nor JS inject could populate.
  let firstIteration = true;
  while (safetyMax-- > 0) {
    if (firstIteration) {
      // Reuse the snapshot from the modal-wait so we don't burn an extra MCP roundtrip.
      snapText = modalWait.snapText;
      fields = modalWait.fields;
      firstIteration = false;
    } else {
      snap = await mcp.snapshot();
      snapText = JSON.stringify(snap?.content || snap || '').slice(0, 18000);
      fields = await extractFields(snapText);
    }

    // Fill each field
    for (const f of (fields.fields || [])) {
      if (f.kind === 'file' && f.ref === fields.file_upload_ref) {
        await mcp.fileUpload([session.cv_pdf]);
        await updateState(session, 'cv_uploaded', { field: f.label });
        await new Promise(r => setTimeout(r, 800));
        continue;
      }
      let val = tryMatchFromProfile(f, profile);
      if (val == null || val === '') {
        // Ask user
        const sinceIso = new Date().toISOString();
        await postChatMessage(`❓ LinkedIn pregunta: «${f.label}»\n\n¿Qué respondo? (responde aquí, lo recordaré para futuras aplicaciones)`, {
          type: 'auto_apply_question',
          field_hash: fieldHash(f),
          field_label: f.label,
          field_kind: f.kind,
          session_id: sessionId
        });
        await updateState(session, 'awaiting_user_answer', { question_label: f.label });
        const reply = await waitForUserReply(sinceIso, 10 * 60 * 1000);
        if (!reply) {
          await updateState(session, 'failed', { error: 'user did not answer within 10 min' });
          return;
        }
        val = reply.text.trim();
        await persistAnswer(f, val);
        await updateState(session, 'filling', { answered: f.label, with: val });
      }
      if (f.kind === 'select') {
        // Native HTML <select> or LinkedIn combobox.
        try {
          await mcp.selectOption(f.label, f.ref, [val]);
        } catch (e) {
          // Fallback to click + type for custom comboboxes that ignore select_option.
          await mcp.click(f.label, f.ref);
          await new Promise(r => setTimeout(r, 200));
          await mcp.type(f.label, f.ref, String(val));
        }
      } else if (f.kind === 'checkbox' || f.kind === 'radio') {
        if (/^(s[ií]|yes|true|on)$/i.test(String(val).trim())) await mcp.click(f.label, f.ref);
      } else {
        // For text/tel/email/number/textarea: type-first, then readback, then JS fallback.
        // LinkedIn input masks (tel field especially) often swallow native keystrokes
        // sent by browser_type. We verify the value actually landed and inject via JS if not.
        try {
          await mcp.click(f.label, f.ref);
          await new Promise(r => setTimeout(r, 150));
        } catch (e) { /* click on input may fail on some wrappers; type can still work */ }
        let filledVia = 'type';
        try {
          await mcp.type(f.label, f.ref, String(val));
        } catch (e) { /* type may throw on detached node; we'll JS-inject below */ }
        await new Promise(r => setTimeout(r, 200));
        const readback = await readbackValue(mcp, f.label, f.ref);
        if (!readback || String(readback).trim() === '') {
          const js = await setInputViaJS(mcp, f.label, f.ref, String(val));
          if (js && js.ok && js.actualValue != null && String(js.actualValue).trim() !== '') {
            filledVia = 'js_fallback';
          } else {
            // Genuine fill failure — register for the Gate 2 warning.
            fillFailures.push({ label: f.label, ref: f.ref, value_attempted: String(val), error: (js && js.error) || 'readback empty' });
            filledVia = 'failed';
          }
        }
        try { await updateState(session, 'filling', { field: f.label, filled_via: filledVia }); } catch {}
      }
      await new Promise(r => setTimeout(r, 300 + Math.random() * 500));
    }

    // Track whether this iteration actually filled anything (file upload counts).
    if (Array.isArray(fields.fields) && fields.fields.length > 0) filledAnyField = true;

    // CRITICAL: re-snapshot + re-extract now that fields are filled.
    // LinkedIn disables the Next/Submit/Review button until required fields are populated.
    // The `fields` we extracted at the start of this iteration captured the buttons in
    // their DISABLED state (which claudemax correctly excludes from *_button_ref). After
    // filling the required fields, the buttons become enabled — we need a fresh extract.
    if (Array.isArray(fields.fields) && fields.fields.length > 0) {
      await new Promise(r => setTimeout(r, 600)); // let DOM update enabled state
      const freshSnap = await mcp.snapshot();
      const freshText = JSON.stringify(freshSnap?.content || freshSnap || '').slice(0, 18000);
      const freshFields = await extractFields(freshText);
      // Carry forward only the button refs (don't re-fill on the same fields).
      fields.submit_button_ref = freshFields.submit_button_ref || fields.submit_button_ref;
      fields.review_button_ref = freshFields.review_button_ref || fields.review_button_ref;
      fields.next_button_ref = freshFields.next_button_ref || fields.next_button_ref;
      fields.file_upload_ref = freshFields.file_upload_ref || fields.file_upload_ref;
    }

    // Advance: Review (if present) or Next
    if (fields.submit_button_ref) {
      // We've reached the final page; trigger Gate 2.
      break;
    }
    if (fields.review_button_ref) {
      await mcp.click('Review', fields.review_button_ref);
    } else if (fields.next_button_ref) {
      await mcp.click('Next', fields.next_button_ref);
    } else {
      // No advance button found.
      if (!filledAnyField) {
        // Modal never produced fillable content nor an advance button: orchestrator misfired.
        await postChatMessage(`⚠️ El modal de Easy Apply no expuso ningún campo ni botón de avance en «${session.offer_title}». Aborto sin tocar la cuota.`, { type: 'modal_empty_unexpected' });
        await updateState(session, 'failed', { error: 'modal_empty_unexpected' });
        return;
      }
      // Already filled something — likely on final page without explicit submit ref. Break to Gate 2.
      break;
    }
    // Wait for the next wizard step to render before re-snapshotting.
    const stepWait = await waitForApplyModalReady(mcp);
    if (stepWait.timeout) {
      await postChatMessage(`⚠️ Siguiente paso del wizard de Easy Apply no cargó tras 9s en «${session.offer_title}». Aborto.`, { type: 'next_step_timeout' });
      await updateState(session, 'failed', { error: 'next_step_timeout' });
      return;
    }
    // Carry the freshly waited snapshot/fields into the next iteration via the firstIteration mechanism.
    modalWait.snapText = stepWait.snapText;
    modalWait.fields = stepWait.fields;
    firstIteration = true;
  }

  // 7. Gate 2 — screenshot + confirmation
  const shotPath = path.join(SHOTS_DIR, `${sessionId}.png`);
  await fs.mkdir(SHOTS_DIR, { recursive: true });
  const shot = await mcp.screenshot();
  if (shot.dataB64) {
    await fs.writeFile(shotPath, Buffer.from(shot.dataB64, 'base64'));
  }

  // Final-state introspection: detect any required fields still empty so the chat
  // message can warn the user. LinkedIn won't accept the submit if validation fails.
  // We use BOTH the LLM extractor (claudemax) AND a deterministic DOM scan.
  // DOM scan is ground-truth; LLM extractor catches stuff the DOM scan can't.
  let pendingRequired = [];
  let domAutoFilled = [];     // Fase 5d: fields we rescued via DOM scan + JS inject
  try {
    const finalSnap = await mcp.snapshot();
    const finalText = JSON.stringify(finalSnap?.content || finalSnap || '').slice(0, 18000);
    const finalFields = await extractFields(finalText);
    for (const f of (finalFields.fields || [])) {
      if (f.required && (!f.current_value || String(f.current_value).trim() === '')) {
        pendingRequired.push(f.label);
      }
    }
  } catch (_) { /* best-effort warning */ }

  // Fase 5d: deterministic DOM scan for required empty inputs the LLM missed.
  try {
    const DEBUG_FILE = path.join(VAULT_ROOT, 'inbox', 'job-hunting', 'fase5d-debug.jsonl');
    const dlog = async (obj) => {
      try { await fs.appendFile(DEBUG_FILE, JSON.stringify(obj) + '\n'); } catch {}
    };
    const domEmpty = await scanRequiredEmptyFields(mcp);
    await dlog({ ts: new Date().toISOString(), session_id: sessionId, stage: 'scan', domEmpty });
    for (const f of domEmpty) {
      const fakeField = { label: f.label, kind: f.type === 'select-one' ? 'select' : (f.type || 'text'), options: [] };
      const val = tryMatchFromProfile(fakeField, profile);
      await dlog({ ts: new Date().toISOString(), session_id: sessionId, stage: 'rescue-attempt', field: f, matchedVal: val });
      if (val != null && val !== '' && f.id) {
        const r = await jsFillById(mcp, f.id, val);
        await dlog({ ts: new Date().toISOString(), session_id: sessionId, stage: 'jsFillById-result', id: f.id, result: r });
        if (r && r.ok && (r.actualValue || '').trim() !== '') {
          domAutoFilled.push({ label: f.label, value: String(val), source: 'profile' });
          continue;
        }
      }
      if (!pendingRequired.includes(f.label)) pendingRequired.push(f.label);
    }
  } catch (e) {
    try { await fs.appendFile(path.join(VAULT_ROOT, 'inbox', 'job-hunting', 'fase5d-debug.jsonl'),
      JSON.stringify({ ts: new Date().toISOString(), session_id: sessionId, stage: 'ERROR', error: String(e && e.message || e) }) + '\n'); } catch {}
  }

  await updateState(session, 'gate2_review', {
    screenshot: shotPath,
    screenshot_bytes: shot.dataB64 ? Buffer.from(shot.dataB64, 'base64').length : 0,
    pending_required: pendingRequired,
    fill_failures: fillFailures,
    dom_auto_filled: domAutoFilled
  });

  const sinceIso = new Date().toISOString();
  const shotUrl = `/api/trabajo/apply/shot/${sessionId}`;
  // Merge fillFailures (deterministic — what we KNOW failed) into pendingRequired.
  // fillFailures wins over claudemax's snapshot inference because it's direct evidence.
  const knownFailedLabels = fillFailures
    .filter(ff => String(ff.value_attempted || '').trim() !== '')   // skip intentional empties
    .map(ff => ff.label);
  const allWarnLabels = Array.from(new Set([...knownFailedLabels, ...pendingRequired]));
  const warnText = allWarnLabels.length > 0
    ? `\n\n⚠️ **Campos que YO no he podido rellenar** (rellénalos tú manualmente en el navegador antes de aprobar):\n` +
      allWarnLabels.map(l => `   - ${l}`).join('\n')
    : '';
  const rescuedText = domAutoFilled.length > 0
    ? `\n\n🔧 Rescaté ${domAutoFilled.length} campo(s) que el extractor LLM omitió (rellenados via DOM scan): ` +
      domAutoFilled.map(f => `${f.label}=${f.value.slice(0, 30)}`).join(' · ')
    : '';
  await postChatMessage(
    `✅ Formulario para «${session.offer_title}». ` +
    (allWarnLabels.length === 0
      ? 'He rellenado todo y adjuntado el CV.'
      : `He rellenado lo que he podido. Quedan ${allWarnLabels.length} campo(s) por completar a mano.`) +
    `\n\n**Captura del formulario**: [ver screenshot](${shotUrl})${warnText}${rescuedText}\n\n` +
    `¿Envío la aplicación? Responde **\`sí\`** o **\`no\`** (o cualquier mensaje con "no" si quieres cancelar).`,
    { type: 'gate2_confirm', session_id: sessionId, screenshot_url: shotUrl, screenshot_path: path.relative(VAULT_ROOT, shotPath), pending_required: pendingRequired, dom_auto_filled: domAutoFilled }
  );

  const reply = await waitForUserReply(sinceIso, 30 * 60 * 1000);
  if (!reply) {
    await postChatMessage(`⏰ Timeout 30 min sin respuesta. Aborto sin enviar la aplicación.`, { type: 'gate2_timeout' });
    await updateState(session, 'cancelled_timeout');
    return;
  }
  const yes = /^(s[ií]|yes|y|ok|env[ií]a|adelante|go)\b/i.test(reply.text.trim());
  if (!yes) {
    await postChatMessage(`❌ OK, no envío. Aplicación cancelada por ti.`, { type: 'gate2_rejected' });
    await updateState(session, 'cancelled_by_user');
    return;
  }

  // 8. Submit
  await updateState(session, 'submitting');
  // Re-fetch the submit ref (page may have changed)
  snap = await mcp.snapshot();
  snapText = JSON.stringify(snap?.content || snap || '').slice(0, 18000);
  fields = await extractFields(snapText);
  if (!fields.submit_button_ref) {
    await postChatMessage(`⚠️ No encuentro el botón Submit ahora. Aborto. Termina manualmente en la ventana del navegador.`, { type: 'submit_missing' });
    await updateState(session, 'failed', { error: 'submit button not found' });
    return;
  }
  await mcp.click('Submit', fields.submit_button_ref);
  // Fase 5c: verify the submit actually succeeded before incrementing quota.
  const confirm = await waitForSubmitConfirmation(mcp);
  if (!confirm.ok) {
    // Save a screenshot for forensic diagnosis.
    try {
      const failShot = await mcp.screenshot();
      if (failShot.dataB64) {
        const failPath = path.join(SHOTS_DIR, `${sessionId}.submit-fail.png`);
        await fs.writeFile(failPath, Buffer.from(failShot.dataB64, 'base64'));
      }
    } catch (_) {}
    await logError(sessionId, 'submit_unconfirmed', new Error('submit click sent but no confirmation signal within timeout'));
    await postChatMessage(`⚠️ Pulsé Enviar en «${session.offer_title}» pero no he detectado confirmación de LinkedIn en 12s. **No incremento cuota**. Comprueba el navegador para ver si llegó o si quedó algún campo por validar.`, { type: 'submit_unconfirmed' });
    await updateState(session, 'failed', { error: 'submit_unconfirmed' });
    return;
  }
  await rateLimitIncrement();
  await postChatMessage(`🎉 Aplicación enviada a «${session.offer_title}» en ${session.company}. Cuota hoy: actualizada. (confirmado vía ${confirm.signal})`, { type: 'applied', signal: confirm.signal });
  await updateState(session, 'applied', { confirm_signal: confirm.signal });
}

// Read-only helpers exposed for endpoints
export async function listSessions() {
  try {
    const raw = await fs.readFile(SESSIONS_FILE, 'utf8');
    const lines = raw.split('\n').filter(Boolean);
    const byId = new Map();
    for (const ln of lines) {
      try { const r = JSON.parse(ln); if (r.session_id) byId.set(r.session_id, r); } catch {}
    }
    return Array.from(byId.values()).sort((a, b) => (b.last_update || '').localeCompare(a.last_update || ''));
  } catch (e) {
    if (e.code === 'ENOENT') return [];
    throw e;
  }
}

export async function getRateLimit() {
  return rateLimitCheck();
}
