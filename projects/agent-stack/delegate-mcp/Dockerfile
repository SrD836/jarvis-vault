FROM python:3.12-alpine

# docker CLI to exec into gateway container (mounts /var/run/docker.sock)
RUN apk add --no-cache docker-cli

WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY delegate_mcp.py .

ENV PORT=38900
EXPOSE 38900

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -q -O - http://127.0.0.1:38900/health || exit 1

CMD ["python3", "delegate_mcp.py"]
