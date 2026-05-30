---
title: "code-reviewer run — 2026-05-29 152234"
type: agent-run
date: 2026-05-29T15:22:34+00:00
agent: "[[agents/code-reviewer]]"
spawned_from: "[[agents/planner]]"
parent_chain: ["main", "planner"]
model: deepseek-chat
provider: deepseek
iter: 1
tokens: {"input_tokens": 73, "output_tokens": 483, "estimated": true}
tool_calls_n: 0
duration_ms: 5801
aborted: false
abort_reason: ""
spawned_children: []
skill: "gsd-code-review"
related:
  - "[[agents/CATALOG]]"
  - "[[agents/auditor]]"
  - "[[agents/code-reviewer]]"
  - "[[agents/debugger]]"
  - "[[agents/tester]]"
tags: ["agent-run", "code-reviewer", "planner-routed"]
# auto-linked 2026-05-29
---


## TL;DR
code-reviewer: Lista de hallazgos con severidad y codigo fix para cada uno. -> OK.

## Task recibida
```
Revisa la funcion Python login(u,p) que construye una query SQL concatenando u y p directamente en un f-string SELECT. Identifica bugs y vulnerabilidades de seguridad (inyeccion SQL, etc.). Lista hallazgos por severidad (Critical, High, Medium, Low) y proporciona el codigo fix para cada uno.
```

## Output
**Hallazgos:**

**Critical: Inyección SQL directa**
- **Descripción:** La función `login(u, p)` construye la query SQL concatenando `u` y `p` directamente en un f-string. Un atacante puede inyectar comandos SQL arbitrarios.
- **Código vulnerable:**
```python
def login(u, p):
    query = f"SELECT * FROM users WHERE username='{u}' AND password='{p}'"
    cursor.execute(query)
```
- **Fix:**
```python
def login(u, p):
    query = "SELECT * FROM users WHERE username = %s AND password = %s"
    cursor.execute(query, (u, p))
```

**High: Almacenamiento de contraseñas en texto plano**
- **Descripción:** La función compara directamente `p` con el valor almacenado en BD. Las contraseñas deben almacenarse con hash (ej. bcrypt) y compararse con `check_password_hash`.
- **Código vulnerable:**
```python
# Asume que password está en texto plano en BD
```
- **Fix:**
```python
import bcrypt

def login(u, p):
    query = "SELECT password_hash FROM users WHERE username = %s"
    cursor.execute(query, (u,))
    result = cursor.fetchone()
    if result and bcrypt.checkpw(p.encode('utf-8'), result[0].encode('utf-8')):
        return True
    return False
```

**Medium: Falta de manejo de errores/excepciones**
- **Descripción:** No hay try/except alrededor de `cursor.execute()`. Un error de BD puede exponer información interna o causar crash.
- **Código vulnerable:**
```python
cursor.execute(query)
```
- **Fix:**
```python
try:
    cursor.execute(query, (u, p))
    result = cursor.fetchone()
except Exception as e:
    log.error(f"Error en login: {e}")
    return False
```

**Low: Sin límite de intentos ni rate limiting**
- **Descripción:** No hay protección contra ataques de fuerza bruta.
- **Fix (opcional):**
```python
from flask_limiter import Limiter

limiter = Limiter(key_func=lambda: request.remote_addr)
@app.route('/login', methods=['POST'])
@limiter.limit("5 per minute")
def login_route():
    # código de login
```

## Human notes
_(no se toca por automatizacion)_
