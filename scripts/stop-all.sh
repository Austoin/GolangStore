#!/usr/bin/env bash
set -euo pipefail

stop_port() {
    local port="$1"
    local output
    output="$(powershell -Command "Get-NetTCPConnection -LocalPort $port -State Listen -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess" 2>/dev/null | tr -d '\r' || true)"
    while IFS= read -r pid; do
        [[ -z "$pid" ]] && continue
        powershell -Command "Stop-Process -Id $pid -Force" >/dev/null 2>&1 || true
    done <<< "$output"
}

stop_port 3000
stop_port 8081
stop_port 8082
stop_port 8083

docker compose stop mysql redis >/dev/null 2>&1 || true
