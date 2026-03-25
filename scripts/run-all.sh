#!/usr/bin/env bash
set -euo pipefail

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
readonly FRONTEND_DIR="$PROJECT_ROOT/frontend"
readonly FRONTEND_LOG="$PROJECT_ROOT/.runtime/frontend.log"
readonly FRONTEND_PID_FILE="$PROJECT_ROOT/.runtime/frontend.pid"

log() {
    echo "[$(date '+%H:%M:%S')] $*" >&2
}

die() {
    log "ERROR: $*"
    exit 1
}

wait_for_http() {
    local url="$1"
    local label="$2"
    local retries=40
    local count=0

    while [[ $count -lt $retries ]]; do
        if curl -fsS "$url" >/dev/null 2>&1; then
            log "HTTP ready: $url"
            return 0
        fi
        ((count+=1))
        log "waiting for $label ($count/$retries): $url"
        sleep 2
    done

    if [[ -f "$FRONTEND_LOG" ]]; then
        log "frontend log tail:"
        tail -n 40 "$FRONTEND_LOG" >&2 || true
    fi

    die "http endpoint did not become ready: $url"
}

stop_frontend() {
    if [[ -f "$FRONTEND_PID_FILE" ]]; then
        local old_pid
        old_pid="$(cat "$FRONTEND_PID_FILE")"
        if kill -0 "$old_pid" >/dev/null 2>&1; then
            log "stopping existing frontend pid=$old_pid"
            kill "$old_pid" >/dev/null 2>&1 || true
            sleep 1
        fi
        rm -f "$FRONTEND_PID_FILE"
    fi

    local output
    output="$(powershell -Command "Get-NetTCPConnection -LocalPort 3000 -State Listen -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess" 2>/dev/null | tr -d '\r' || true)"
    while IFS= read -r pid; do
        [[ -z "$pid" ]] && continue
        log "stopping process on port 3000 pid=$pid"
        powershell -Command "Stop-Process -Id $pid -Force" >/dev/null 2>&1 || true
    done <<< "$output"
}

ensure_frontend_deps() {
    if [[ ! -d "$FRONTEND_DIR/node_modules" ]]; then
        log "installing frontend dependencies"
        npm install --prefix "$FRONTEND_DIR"
    fi
}

start_frontend() {
    mkdir -p "$PROJECT_ROOT/.runtime"
    stop_frontend
    ensure_frontend_deps

    log "starting frontend"
    nohup npm --prefix "$FRONTEND_DIR" run dev -- --hostname 0.0.0.0 --port 3000 >"$FRONTEND_LOG" 2>&1 &
    local pid=$!
    echo "$pid" >"$FRONTEND_PID_FILE"
    log "frontend pid=$pid"
    log "frontend root: http://127.0.0.1:3000/"
    log "frontend shop: http://127.0.0.1:3000/shop"
    log "frontend admin: http://127.0.0.1:3000/admin"
    log "first startup may take longer because Next.js can download SWC"
}

main() {
    command -v bash >/dev/null 2>&1 || die "bash is required"
    command -v curl >/dev/null 2>&1 || die "curl is required"
    command -v npm >/dev/null 2>&1 || die "npm is required"

    bash "$SCRIPT_DIR/start.sh"
    start_frontend

    wait_for_http http://127.0.0.1:3000/ "frontend root"
    wait_for_http http://127.0.0.1:3000/shop "shop page"
    wait_for_http http://127.0.0.1:3000/admin "admin page"

    log "all services are ready"
    log "frontend log: $FRONTEND_LOG"
    log "shop url: http://127.0.0.1:3000/shop"
    log "admin url: http://127.0.0.1:3000/admin"
}

main "$@"
