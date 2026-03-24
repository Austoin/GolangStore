# Real Cart To Order Orchestration Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Make `/orders/from-cart` create orders by loading checked cart items from the cart service layer using only `user_id` input.

**Architecture:** Keep cart item ownership in `internal/cart` and order creation ownership in `internal/order`. Add orchestration by injecting cart query capability into the order service instead of passing cart items over HTTP request bodies.

**Tech Stack:** Golang, Gin, standard library testing

---

### Task 1: Introduce cart query dependency into order orchestration

**Files:**
- Modify: `internal/order/service.go`
- Modify: `internal/order/service_test.go`

**Step 1: Write the failing test**

Add tests verifying that order orchestration loads checked cart items through a cart dependency using only `user_id`.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the service does not depend on cart querying yet.

**Step 3: Write minimal implementation**

Inject a cart query dependency and create the order from checked items returned by that dependency.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

### Task 2: Shrink `/orders/from-cart` request model

**Files:**
- Modify: `internal/order/handler.go`
- Modify: `internal/order/handler_test.go`

**Step 1: Write the failing test**

Change tests so `/orders/from-cart` only accepts `user_id`.

**Step 2: Run test to verify it fails**

Run: `go test ./internal/order`
Expected: FAIL because the handler still expects cart items in the request body.

**Step 3: Write minimal implementation**

Update the request model and handler to use only `user_id` and delegate cart loading to the service.

**Step 4: Run test to verify it passes**

Run: `go test ./internal/order`
Expected: PASS.

### Task 3: Wire runtime cart dependency in order-service

**Files:**
- Modify: `cmd/order-service/main.go`
- Modify: `cmd/order-service/router_test.go`

**Step 1: Assemble cart service and order service together**

Use in-memory cart repository plus in-memory order repository for runtime wiring.

**Step 2: Verify service entrypoint**

Run: `go test ./cmd/order-service`
Expected: PASS.

### Task 4: Update docs and verify project state

**Files:**
- Modify: `README.md`
- Modify: `docs/process/stage-08-真实购物车到下单编排.md`

**Step 1: Record progress and current limitations**

Update docs with the new orchestration mode.

**Step 2: Run full verification**

Run: `go test ./...`
Expected: PASS.
