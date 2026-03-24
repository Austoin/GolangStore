#!/usr/bin/env bash
set -euo pipefail

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
readonly RUNTIME_DIR="$PROJECT_ROOT/.runtime"
readonly PRODUCT_LOG="$RUNTIME_DIR/product-service.log"
readonly CART_LOG="$RUNTIME_DIR/cart-service.log"
readonly ORDER_LOG="$RUNTIME_DIR/order-service.log"
readonly PRODUCT_PID_FILE="$RUNTIME_DIR/product-service.pid"
readonly CART_PID_FILE="$RUNTIME_DIR/cart-service.pid"
readonly ORDER_PID_FILE="$RUNTIME_DIR/order-service.pid"
readonly SCRIPT_VERSION="2026-03-25-stage10"
GO_BIN=""

log() {
    echo "[$(date '+%H:%M:%S')] $*" >&2
}

die() {
    log "ERROR: $*"
    exit 1
}

resolve_go_bin() {
    if command -v go >/dev/null 2>&1; then
        GO_BIN="$(command -v go)"
        return 0
    fi

    if [[ -x "/mnt/c/Program Files/Go/bin/go.exe" ]]; then
        GO_BIN="/mnt/c/Program Files/Go/bin/go.exe"
        return 0
    fi

    if [[ -x "/c/Program Files/Go/bin/go.exe" ]]; then
        GO_BIN="/c/Program Files/Go/bin/go.exe"
        return 0
    fi

    return 1
}

wait_for_compose_health() {
    local service="$1"
    local retries=30
    local count=0

    while [[ $count -lt $retries ]]; do
        local status
        status=$(docker inspect --format '{{if .State.Health}}{{.State.Health.Status}}{{else}}{{.State.Status}}{{end}}' "$service" 2>/dev/null || true)
        if [[ "$status" == "healthy" || "$status" == "running" ]]; then
            log "$service is ready: $status"
            return 0
        fi
        ((count+=1))
        sleep 2
    done

    die "service did not become ready: $service"
}

wait_for_http() {
    local url="$1"
    local retries=30
    local count=0

    while [[ $count -lt $retries ]]; do
        if curl -fsS "$url" >/dev/null 2>&1; then
            log "HTTP ready: $url"
            return 0
        fi
        ((count+=1))
        sleep 2
    done

    die "http endpoint did not become ready: $url"
}

stop_existing_process() {
    local pid_file="$1"
    local name="$2"

    if [[ -f "$pid_file" ]]; then
        local old_pid
        old_pid="$(cat "$pid_file")"
        if kill -0 "$old_pid" >/dev/null 2>&1; then
            log "stopping existing $name pid=$old_pid"
            kill "$old_pid" >/dev/null 2>&1 || true
            sleep 1
        fi
        rm -f "$pid_file"
    fi
}

stop_process_on_port() {
    local port="$1"
    local output

    output="$(powershell -Command "Get-NetTCPConnection -LocalPort $port -State Listen -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess" 2>/dev/null | tr -d '\r' || true)"
    if [[ -z "$output" ]]; then
        return 0
    fi

    while IFS= read -r pid; do
        [[ -z "$pid" ]] && continue
        log "stopping process on port $port pid=$pid"
        powershell -Command "Stop-Process -Id $pid -Force" >/dev/null 2>&1 || true
    done <<< "$output"

    sleep 1
}

start_service() {
    local name="$1"
    local entry="$2"
    local log_file="$3"
    local pid_file="$4"
    local port="$5"

    mkdir -p "$RUNTIME_DIR"

    stop_existing_process "$pid_file" "$name"
    stop_process_on_port "$port"

    log "starting $name with: $GO_BIN run ./$entry"
    nohup "$GO_BIN" run "./$entry" >"$log_file" 2>&1 &
    local pid=$!
    echo "$pid" >"$pid_file"
    log "$name pid=$pid"
}

main() {
    log "script version: $SCRIPT_VERSION"
    log "project root: $PROJECT_ROOT"

    command -v docker >/dev/null 2>&1 || die "docker is required"
    command -v curl >/dev/null 2>&1 || die "curl is required"
    resolve_go_bin || die "go is required"

    log "docker path: $(command -v docker)"
    log "curl path: $(command -v curl)"
    log "go path: $GO_BIN"

    docker info >/dev/null 2>&1 || die "docker daemon is not running"
    log "docker daemon check: ok"

    log "starting mysql and redis with docker compose"
    docker compose up -d mysql redis

    wait_for_compose_health golangstore-mysql
    wait_for_compose_health golangstore-redis

    start_service product-service cmd/product-service "$PRODUCT_LOG" "$PRODUCT_PID_FILE" 8081
    start_service cart-service cmd/cart-service "$CART_LOG" "$CART_PID_FILE" 8083
    start_service order-service cmd/order-service "$ORDER_LOG" "$ORDER_PID_FILE" 8082

    wait_for_http http://127.0.0.1:8081/health
    wait_for_http http://127.0.0.1:8083/health
    wait_for_http http://127.0.0.1:8082/health

    log "startup complete"
    log "product-service log: $PRODUCT_LOG"
    log "cart-service log: $CART_LOG"
    log "order-service log: $ORDER_LOG"
    curl -fsS http://127.0.0.1:8081/health
    printf '\n'
    curl -fsS http://127.0.0.1:8083/health
    printf '\n'
    curl -fsS http://127.0.0.1:8082/health
}

main "$@"
