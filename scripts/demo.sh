#!/usr/bin/env bash
set -euo pipefail

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
readonly MYSQL_CONTAINER="golangstore-mysql"
readonly ORDER_BASE_URL="http://127.0.0.1:8082"
readonly CART_BASE_URL="http://127.0.0.1:8083"

log() {
    echo "[$(date '+%H:%M:%S')] $*" >&2
}

die() {
    log "ERROR: $*"
    exit 1
}

require_cmd() {
    command -v "$1" >/dev/null 2>&1 || die "$1 is required"
}

run_sql() {
    local sql="$1"
    docker exec "$MYSQL_CONTAINER" mysql -uroot -proot golang_store -N -e "$sql"
}

seed_demo_data() {
    log "resetting demo tables"
    run_sql "DELETE FROM order_items; DELETE FROM orders; DELETE FROM cart_items; DELETE FROM product_stocks; DELETE FROM products;"

    log "seeding product and stock"
    run_sql "INSERT INTO products (id, name, description, price, status) VALUES (101, 'phone', 'smart phone', 199900, 1);"
    run_sql "INSERT INTO product_stocks (product_id, stock) VALUES (101, 5);"
}

add_cart_item() {
    log "adding cart item through cart-service"
    curl -fsS -X POST "$CART_BASE_URL/carts" \
        -H "Content-Type: application/json" \
        -d '{"user_id":1,"product_id":101,"product_name":"phone","price":199900,"quantity":2,"checked":true}' >/dev/null
}

create_order() {
    log "creating order from checked cart items"
    curl -fsS -X POST "$ORDER_BASE_URL/orders/from-cart" \
        -H "Content-Type: application/json" \
        -d '{"user_id":1}'
}

verify_stock() {
    local stock
    stock="$(run_sql "SELECT stock FROM product_stocks WHERE product_id = 101;")"
    if [[ "$stock" != "3" ]]; then
        die "expected stock 3 after order creation, got ${stock:-empty}"
    fi
    log "stock verification passed: $stock"
}

verify_order() {
    local order_count
    order_count="$(run_sql "SELECT COUNT(*) FROM orders;")"
    if [[ "$order_count" != "1" ]]; then
        die "expected 1 order, got ${order_count:-empty}"
    fi

    local order_item_count
    order_item_count="$(run_sql "SELECT COUNT(*) FROM order_items;")"
    if [[ "$order_item_count" != "1" ]]; then
        die "expected 1 order item, got ${order_item_count:-empty}"
    fi

    log "order persistence verification passed"
}

main() {
    require_cmd docker
    require_cmd curl

    docker ps --format '{{.Names}}' | grep -qx "$MYSQL_CONTAINER" || die "mysql container is not running"
    curl -fsS "$ORDER_BASE_URL/health" >/dev/null || die "order-service is not healthy"
    curl -fsS "$CART_BASE_URL/health" >/dev/null || die "cart-service is not healthy"

    seed_demo_data
    add_cart_item
    local response
    response="$(create_order)"
    log "order response: $response"
    verify_order
    verify_stock
    printf '%s\n' "$response"
}

main "$@"
