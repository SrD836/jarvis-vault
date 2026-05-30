---
title: "C1 Tunnel Doctor"
type: runbook
related:
  - "[[agents/apier]]"
  - "[[agents/c1-doctor]]"
  - "[[agents/debugger]]"
  - "[[agents/monitor]]"
  - "[[projects/trabajo-v2]]"
  - "[[projects/trabajo/project]]"
  - "[[skills/external/gsd-debug]]"
tags: [runbook, c1, browser-automation, infra]
# auto-linked 2026-05-20
---


# C1 Tunnel — Runbook

El **puente C1** conecta el VPS con el navegador Chrome del PC para que el sistema pueda controlar el navegador remotamente (LinkedIn auto-apply, scraping autenticado, etc.). La cadena tiene 4 piezas que deben estar todas vivas:

```
[PC]                                      [VPS]
Chrome :9222 (CDP)  <--  pc-playwright (Node MCP server, escucha en 127.0.0.1:8765)
                                      ^
                                      |  autossh -R 8765:localhost:8765 agent@88.198.168.61
                                      v
                              VPS:127.0.0.1:8765  <-- mcp-pc-relay (Docker, autossh -L 0.0.0.0:8765:127.0.0.1:8765)
                                                                                ^
                                                                                |
                                                                       MCP clients (gateway, dashboard) via container :8765
```

## Diagnóstico instantáneo

Desde cualquier máquina con curl al dashboard:

```bash
curl -s http://VPS:3000/api/c1/health | python3 -m json.tool
```

O directamente desde el VPS:

```bash
python3 /home/agent/agent-stack/hermes/c1_tunnel_doctor.py
```

El reporte clasifica el fallo en uno de 3 estados:

| Síntoma | Significa | Arreglo |
|---|---|---|
| `mcp_pc_relay_container: ok=false` | Container relay caído en VPS | `docker restart mcp-pc-relay` o `curl http://VPS:3000/api/c1/health?repair=1` |
| `vps_listener_8765: ok=false (Connection refused)` | El PC no tiene su autossh reverso activo | Acción **en el PC**: arrancar el autossh (sección abajo) |
| `mcp_handshake: timed out` o `non-JSON` | Listener arriba pero pc-playwright no responde | Acción **en el PC**: arrancar/reiniciar pc-playwright MCP |

## Acciones en el PC (Windows)

### 1) Verificar que Chrome esté abierto con CDP

Chrome necesita estar abierto con el flag `--remote-debugging-port=9222` apuntando al perfil dedicado `C:\jarvis\chrome-cdp-profile`. Comando de arranque típico:

```powershell
& "C:\Program Files\Google\Chrome\Application\chrome.exe" `
  --remote-debugging-port=9222 `
  --user-data-dir="C:\jarvis\chrome-cdp-profile" `
  --no-first-run --no-default-browser-check
```

Comprobar que está vivo:

```powershell
Test-NetConnection -ComputerName localhost -Port 9222
# Debe mostrar TcpTestSucceeded : True
```

### 2) Arrancar `pc-playwright` MCP (Node)

El servicio escucha en `127.0.0.1:8765` y traduce JSON-RPC ↔ CDP. El comando real depende de cómo se instaló pero será algo como:

```powershell
cd C:\jarvis\pc-playwright
node server.js --port 8765 --cdp-endpoint http://localhost:9222
```

O via npm script:

```powershell
cd C:\jarvis\pc-playwright
npm start
```

Comprobar handshake básico:

```powershell
$payload = @{ jsonrpc='2.0'; id=1; method='initialize'; params=@{ protocolVersion='2024-11-05'; capabilities=@{}; clientInfo=@{ name='test'; version='1.0' } } } | ConvertTo-Json -Compress
$tcp = New-Object System.Net.Sockets.TcpClient('127.0.0.1', 8765)
$stream = $tcp.GetStream()
$bytes = [Text.Encoding]::UTF8.GetBytes("$payload`n")
$stream.Write($bytes, 0, $bytes.Length)
$reader = New-Object IO.StreamReader($stream)
$reader.ReadLine()
# Debe devolver un JSON con jsonrpc:'2.0' y un campo result
$tcp.Close()
```

### 3) Arrancar/restaurar el autossh reverso

Túnel reverso desde PC al VPS. Si no existe ya como servicio, este comando lo inicia:

```powershell
autossh -M 0 `
  -o ServerAliveInterval=30 -o ServerAliveCountMax=3 `
  -o ExitOnForwardFailure=yes `
  -R 8765:localhost:8765 `
  -i C:\jarvis\keys\id_ed25519 `
  agent@88.198.168.61 -N
```

Comprobar desde el VPS que el listener subió:

```bash
ss -tlnp | grep 8765   # Debe haber LISTEN en 127.0.0.1:8765
```

## Health-check automatizado

El dashboard expone `/api/c1/health` y debería consultarlo:

- **Antes de iniciar cualquier auto-apply** (Fase 5): si `ok=false`, abortar y postear `next_action` al chat de /trabajo.
- **Periódicamente** (cron 5 min recomendado) para alertar pérdidas de conexión silenciosas.

## Histórico de incidencias

| Fecha | Síntoma | Causa | Resolución |
|---|---|---|---|
| 2026-05-19 | handshake timeout / listener refused | pc-playwright + autossh caídos en PC tras reboot | Manual restart en PC |
