# End To End Business Completion Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Complete the remaining core mall backend business logic so the project can be started with one command and verified with a repeatable end-to-end demo flow.

**Architecture:** Build on the existing monorepo structure and keep business logic in `internal/*` packages. Finish the real persistence and orchestration paths first, then expose them via service entrypoints, and finally add a reproducible demo script that proves the system works without manual debugging.

**Tech Stack:** Golang, Gin, Gorm, MySQL, Redis, Docker Compose, Bash

---

### Task 1: Complete product persistence and query path

**Files:**
- Create: `internal/product/mysql_repo.go`
- Create: `internal/product/mysql_repo_test.go`
- Modify: `internal/product/repo.go`
- Modify: `internal/product/service.go`
- Modify: `internal/product/handler.go`
- Modify: `cmd/product-service/main.go`
- Test: `internal/product/*.go`
- Test: `cmd/product-service/*.go`

**Step 1: Write the failing tests**

Add tests for loading products and stock from MySQL-backed rows.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/product && go test ./cmd/product-service`
Expected: FAIL because the MySQL repository and runtime wiring do not exist yet.

**Step 3: Write minimal implementation**

Add product MySQL repository and wire `product-service` to use it.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/product && go test ./cmd/product-service`
Expected: PASS.

**Step 5: Commit**

```bash
git add internal/product cmd/product-service README.md docs/process
git commit -m "feat: add product mysql query workflow"
```

### Task 2: Complete cart write workflow

**Files:**
- Modify: `internal/cart/repo.go`
- Modify: `internal/cart/service.go`
- Modify: `internal/cart/handler.go`
- Modify: `internal/cart/mysql_repo.go`
- Modify: `internal/cart/*_test.go`
- Modify: `cmd/cart-service/main.go`

**Step 1: Write the failing tests**

Add tests for add-to-cart and update-cart behavior with MySQL persistence.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/cart && go test ./cmd/cart-service`
Expected: FAIL because cart write APIs do not exist yet.

**Step 3: Write minimal implementation**

Add create/update behavior to the cart repository, service, and handler.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/cart && go test ./cmd/cart-service`
Expected: PASS.

**Step 5: Commit**

```bash
git add internal/cart cmd/cart-service README.md docs/process
git commit -m "feat: add cart write workflow"
```

### Task 3: Add stock validation and deduction in order creation

**Files:**
- Modify: `internal/order/service.go`
- Modify: `internal/order/service_test.go`
- Create: `internal/product/stock_repo.go` or extend product repo
- Modify: `internal/product/*_test.go`

**Step 1: Write the failing tests**

Add tests for successful stock deduction and failure on insufficient stock.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order && go test ./internal/product`
Expected: FAIL because stock validation/deduction is not implemented yet.

**Step 3: Write minimal implementation**

Add stock checking and stock deduction in the order creation path.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order && go test ./internal/product`
Expected: PASS.

**Step 5: Commit**

```bash
git add internal/order internal/product README.md docs/process
git commit -m "feat: add stock validation and deduction"
```

### Task 4: Complete service runtime wiring for product and cart services

**Files:**
- Modify: `cmd/product-service/main.go`
- Modify: `cmd/cart-service/main.go`
- Create/Modify: `cmd/product-service/router.go`
- Create/Modify: `cmd/cart-service/router.go`
- Create/Modify tests under `cmd/product-service/` and `cmd/cart-service/`

**Step 1: Write the failing tests**

Add router tests for the real service routes.

**Step 2: Run test to verify it fails**

Run: `go test ./cmd/product-service && go test ./cmd/cart-service`
Expected: FAIL because the runtime wiring is not complete yet.

**Step 3: Write minimal implementation**

Wire product and cart runtime dependencies using MySQL-backed repositories.

**Step 4: Run test to verify it passes**

Run: `go test ./cmd/product-service && go test ./cmd/cart-service`
Expected: PASS.

**Step 5: Commit**

```bash
git add cmd/product-service cmd/cart-service README.md docs/process
git commit -m "feat: wire product and cart services to mysql"
```

### Task 5: Add one-command demo verification script

**Files:**
- Create: `scripts/demo.sh`
- Modify: `scripts/start.sh`
- Modify: `.gitignore`

**Step 1: Write the failing verification flow**

Define the exact demo flow: seed data, add cart item, create order, query order, verify stock changed.

**Step 2: Run script to verify it fails**

Run: `bash scripts/demo.sh`
Expected: FAIL because the script does not exist yet.

**Step 3: Write minimal implementation**

Add the script and keep the flow deterministic and reproducible.

**Step 4: Run script to verify it passes**

Run: `bash scripts/demo.sh`
Expected: PASS with clear output proving end-to-end behavior.

**Step 5: Commit**

```bash
git add scripts .gitignore README.md docs/process
git commit -m "feat: add end-to-end demo script"
```

### Task 6: Final full verification and documentation closure

**Files:**
- Modify: `README.md`
- Modify: `docs/process/stage-12-端到端业务闭环收束.md`

**Step 1: Update project status docs**

Record what is runnable, what is verified, and the exact commands.

**Step 2: Run full verification**

Run: `go test ./...`
Expected: PASS.

**Step 3: Run startup and demo verification**

Run: `bash scripts/start.sh && bash scripts/demo.sh`
Expected: PASS.

**Step 4: Commit**

```bash
git add README.md docs/process/stage-12-端到端业务闭环收束.md
git commit -m "docs: finalize end-to-end business closure"
```
