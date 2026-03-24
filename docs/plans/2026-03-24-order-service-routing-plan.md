# Order Service Routing Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Wire the existing order query and creation workflows into the running `order-service` HTTP entrypoint.

**Architecture:** Keep `internal/order` as the owner of order business logic, and make `cmd/order-service` only responsible for assembling the in-memory repository, service, and Gin routes. This stage intentionally avoids database and Redis integration so the HTTP surface can be verified independently.

**Tech Stack:** Golang, Gin, standard library testing

---

### Task 1: Add router construction helper

**Files:**
- Create: `cmd/order-service/router.go`
- Test: `cmd/order-service/router_test.go`

**Step 1: Write the failing test**

Add tests verifying that the order-service router responds to `/health` and mounts the order endpoints.

**Step 2: Run test to verify it fails**

Run: `go test ./cmd/order-service`
Expected: FAIL because the router helper does not exist yet.

**Step 3: Write minimal implementation**

Create a router factory that accepts an `order.Handler` and mounts the required routes.

**Step 4: Run test to verify it passes**

Run: `go test ./cmd/order-service`
Expected: PASS.

### Task 2: Wire order-service main entrypoint

**Files:**
- Modify: `cmd/order-service/main.go`
- Modify: `cmd/order-service/router.go`

**Step 1: Assemble memory repository and order handler**

Use existing in-memory order repository and service.

**Step 2: Start Gin using the router helper**

Keep current port `8082` unchanged.

**Step 3: Run package verification**

Run: `go test ./cmd/order-service`
Expected: PASS.

### Task 3: Update docs and verify full project state

**Files:**
- Modify: `README.md`
- Modify: `docs/process/stage-07-order-service路由接线.md`

**Step 1: Record route wiring progress**

Update stage doc and README with the newly exposed order-service routes.

**Step 2: Run full verification**

Run: `go test ./...`
Expected: PASS.
