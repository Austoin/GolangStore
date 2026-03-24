# Cart And Order Persistence Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Replace in-memory cart and order storage in the order workflow with MySQL-backed repositories.

**Architecture:** Keep domain logic in `internal/cart` and `internal/order`, and add MySQL repository implementations that use the existing schema under `deploy/mysql/001_init_schema.sql`. Wire these repositories into `cmd/order-service` only after package-level behavior is verified.

**Tech Stack:** Golang, Gorm, MySQL, Gin, standard library testing

---

### Task 1: Add cart MySQL repository

**Files:**
- Create: `internal/cart/mysql_repo.go`
- Test: `internal/cart/mysql_repo_test.go`

**Step 1: Write the failing test**

Add tests for mapping `cart_items` rows to cart domain items and filtering checked rows by user.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/cart`
Expected: FAIL because the MySQL repository does not exist yet.

**Step 3: Write minimal implementation**

Create a Gorm-based repository implementation that returns checked cart items for a user.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/cart`
Expected: PASS.

### Task 2: Add order MySQL repository

**Files:**
- Create: `internal/order/mysql_repo.go`
- Test: `internal/order/mysql_repo_test.go`

**Step 1: Write the failing test**

Add tests for persisting an order and loading it by `order_no`.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the MySQL repository does not exist yet.

**Step 3: Write minimal implementation**

Create a Gorm-based repository implementation for order write and order query.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

### Task 3: Wire MySQL repositories into order-service

**Files:**
- Modify: `cmd/order-service/main.go`
- Modify: `cmd/order-service/router_test.go`

**Step 1: Assemble MySQL-backed cart and order repositories**

Use `pkg/config` and `pkg/mysql` to bootstrap the runtime dependencies.

**Step 2: Run order-service verification**

Run: `go test ./cmd/order-service`
Expected: PASS.

### Task 4: Update docs and verify project state

**Files:**
- Modify: `README.md`
- Modify: `docs/process/stage-09-购物车与订单持久化接入.md`

**Step 1: Record persistence progress**

Update docs with repository status and current limitations.

**Step 2: Run full verification**

Run: `go test ./...`
Expected: PASS.
