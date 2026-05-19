// ============================================================
// PATCH: Añadir a server.js — endpoint GET /api/bot/pnl
// Insertar ANTES de app.listen(...)
// Si fs ya está requerido arriba, quitar el require duplicado.
// ============================================================

const path = require('path');

const TRADING_DIR = '/vault/inbox/trading';

app.get('/api/bot/pnl', (req, res) => {
  try {
    // Portfolio
    let portfolio = { bankroll: 10000, capital_comprometido: 0, free_capital: 10000 };
    try {
      const pfPath = path.join(TRADING_DIR, 'portfolio.json');
      if (fs.existsSync(pfPath)) {
        portfolio = JSON.parse(fs.readFileSync(pfPath, 'utf-8'));
      }
    } catch (e) { /* defaults */ }

    // Active trades
    let trades_active = [];
    try {
      const actPath = path.join(TRADING_DIR, 'active.jsonl');
      if (fs.existsSync(actPath)) {
        trades_active = fs.readFileSync(actPath, 'utf-8')
          .split('\n').filter(Boolean).map(l => JSON.parse(l));
      }
    } catch (e) { trades_active = []; }

    // Closed trades
    let trades_closed = [];
    try {
      const clPath = path.join(TRADING_DIR, 'closed.jsonl');
      if (fs.existsSync(clPath)) {
        trades_closed = fs.readFileSync(clPath, 'utf-8')
          .split('\n').filter(Boolean).map(l => JSON.parse(l));
      }
    } catch (e) { trades_closed = []; }

    // Rejections
    let rejections = [];
    try {
      const rjPath = path.join(TRADING_DIR, 'rejections.jsonl');
      if (fs.existsSync(rjPath)) {
        rejections = fs.readFileSync(rjPath, 'utf-8')
          .split('\n').filter(Boolean).map(l => JSON.parse(l));
      }
    } catch (e) { rejections = []; }

    // Blocked (vetoed candidates)
    let blockages = [];
    try {
      const blPath = path.join(TRADING_DIR, 'blocked.jsonl');
      if (fs.existsSync(blPath)) {
        blockages = fs.readFileSync(blPath, 'utf-8')
          .split('\n').filter(Boolean).map(l => JSON.parse(l));
      }
    } catch (e) { blockages = []; }

    const pnl_total = trades_closed.reduce((s, t) => s + (t.pnl || 0), 0);
    const pnl_percentage = portfolio.bankroll > 0 ? (pnl_total / portfolio.bankroll) * 100 : 0;
    const closed_count = trades_closed.length;
    const win_rate = closed_count > 0
      ? trades_closed.filter(t => (t.pnl || 0) > 0).length / closed_count
      : 0;

    res.json({
      pnl_total,
      pnl_percentage,
      win_rate,
      active_count: trades_active.length,
      closed_count,
      total_trades: trades_active.length + closed_count,
      bankroll: portfolio.bankroll || 10000,
      free_capital: portfolio.free_capital || 10000,
      capital_comprometido: portfolio.capital_comprometido || 0,
      trades_active,
      trades_closed,
      rejections,
      blockages
    });
  } catch (err) {
    console.error('/api/bot/pnl error:', err);
    res.status(500).json({ error: 'Failed to load bot P&L data' });
  }
});
