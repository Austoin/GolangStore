# Order Creation Stage Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build the first minimal order creation workflow that converts cart items into order items, calculates total amount, and initializes order status for the smart mall backend.

**Architecture:** Keep the implementation inside the existing `internal/order` and `internal/cart` boundaries. Use in-memory repositories and deterministic domain logic first so behavior can be tested without database coupling. Only after the domain and handler flow are stable should persistence be added in a later stage.

**Tech Stack:** Golang, Gin, standard library testing

---

### Task 1: Define order creation input model

**Files:**
- Modify: `internal/order/model.go`
- Test: `internal/order/model_test.go`

**Step 1: Write the failing test**

Add tests to verify a create-order request with a user ID and one or more cart items can be represented by the order domain model.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the create-order input model does not exist yet.

**Step 3: Write minimal implementation**

Add a minimal request/input struct for order creation in `internal/order/model.go`.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS for the new model tests.

**Step 5: Commit**

```bash
git add internal/order/model.go internal/order/model_test.go
git commit -m "feat: add order creation input model"
```

### Task 2: Add total amount calculation

**Files:**
- Modify: `internal/order/model.go`
- Test: `internal/order/model_test.go`

**Step 1: Write the failing test**

Add tests for total amount calculation using order items with integer price-in-cents and quantity.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the calculation method does not exist yet.

**Step 3: Write minimal implementation**

Add a deterministic method that sums `price * quantity` across all order items.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

**Step 5: Commit**

```bash
git add internal/order/model.go internal/order/model_test.go
git commit -m "feat: add order amount calculation"
```

### Task 3: Convert cart items into order items

**Files:**
- Modify: `internal/order/service.go`
- Create: `internal/order/create.go`
- Test: `internal/order/service_test.go`

**Step 1: Write the failing test**

Add tests for converting cart items into order items while preserving product ID, product name, price, and quantity.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the conversion method does not exist yet.

**Step 3: Write minimal implementation**

Add a minimal conversion function and keep it free of database dependencies.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

**Step 5: Commit**

```bash
git add internal/order/create.go internal/order/service.go internal/order/service_test.go
git commit -m "feat: convert cart items into order items"
```

### Task 4: Create in-memory order creation service

**Files:**
- Modify: `internal/order/repo.go`
- Modify: `internal/order/service.go`
- Test: `internal/order/repo_test.go`
- Test: `internal/order/service_test.go`

**Step 1: Write the failing test**

Add tests for creating an order in memory, assigning an order number, setting `StatusPending`, and calculating total amount.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because create-order behavior does not exist yet.

**Step 3: Write minimal implementation**

Add create support to the in-memory repository and service.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

**Step 5: Commit**

```bash
git add internal/order/repo.go internal/order/repo_test.go internal/order/service.go internal/order/service_test.go
git commit -m "feat: add in-memory order creation workflow"
```

### Task 5: Add create-order handler

**Files:**
- Modify: `internal/order/handler.go`
- Test: `internal/order/handler_test.go`

**Step 1: Write the failing test**

Add handler tests for a valid create-order request and an invalid empty-items request.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the create-order handler does not exist yet.

**Step 3: Write minimal implementation**

Add a POST handler that binds request JSON, validates minimal constraints, calls the service, and returns the created order.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

**Step 5: Commit**

```bash
git add internal/order/handler.go internal/order/handler_test.go
git commit -m "feat: add create-order handler"
```

### Task 6: Update documentation and verify full project state

**Files:**
- Modify: `README.md`
- Modify: `docs/process/stage-05-下单创建与状态流转.md`

**Step 1: Update stage document with completed batch details**

Record files, behavior, verification commands, and current limitations.

**Step 2: Update README current progress section**

Add the order creation workflow status.

**Step 3: Run full verification**

Run: `go test ./...`
Expected: PASS.

**Step 4: Commit**

```bash
git add README.md docs/process/stage-05-下单创建与状态流转.md
git commit -m "docs: record order creation stage progress"
```
