# Cart To Order Orchestration Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build the first application-layer workflow that takes checked cart items for a user and creates a pending order with calculated total amount.

**Architecture:** Keep cart and order responsibilities separate: cart remains the source of selected items, order remains the owner of order creation. Add orchestration at the application layer by using cart query behavior plus order creation behavior without introducing MySQL or Redis yet.

**Tech Stack:** Golang, Gin, standard library testing

---

### Task 1: Add checked cart item filtering

**Files:**
- Modify: `internal/cart/service.go`
- Test: `internal/cart/service_test.go`

**Step 1: Write the failing test**

Add a test verifying that only checked items are returned for order creation input.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/cart`
Expected: FAIL because checked-item filtering does not exist yet.

**Step 3: Write minimal implementation**

Add a method that filters checked cart items after listing by user.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/cart`
Expected: PASS.

### Task 2: Add orchestration request model

**Files:**
- Create: `internal/order/orchestrator.go`
- Test: `internal/order/service_test.go`

**Step 1: Write the failing test**

Add a test verifying that a user ID based orchestration request can be consumed by the order application service.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the orchestration model does not exist yet.

**Step 3: Write minimal implementation**

Add the minimal request model and service dependency shape.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

### Task 3: Orchestrate cart to order creation

**Files:**
- Modify: `internal/order/service.go`
- Modify: `internal/order/service_test.go`

**Step 1: Write the failing test**

Add tests verifying that checked cart items become order items, unchecked items are ignored, empty checked carts fail, and order status becomes pending.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the orchestration method does not exist yet.

**Step 3: Write minimal implementation**

Add an orchestration method that receives checked cart items, rejects empty input, converts items, calculates amount, and creates the order.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

### Task 4: Add orchestration handler

**Files:**
- Modify: `internal/order/handler.go`
- Modify: `internal/order/handler_test.go`

**Step 1: Write the failing test**

Add handler tests for user-ID based cart-to-order creation and empty checked cart failure.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the orchestration handler does not exist yet.

**Step 3: Write minimal implementation**

Add the handler and wire it to the orchestration service.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

### Task 5: Update docs and verify project state

**Files:**
- Modify: `README.md`
- Modify: `docs/process/stage-06-购物车到下单编排.md`

**Step 1: Record stage progress**

Update docs with files, behavior, limitations, and verification results.

**Step 2: Run full verification**

Run: `go test ./...`
Expected: PASS.
