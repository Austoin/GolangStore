# MVP Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build the first deliverable MVP for a Go microservice-based smart mall backend with Redis-backed inventory, MySQL persistence, and an MCP gateway exposing core mall capabilities to LLM agents.

**Architecture:** Use a monorepo with multiple Go services. Keep the first phase intentionally narrow: gateway, mcp-gateway, product-service, cart-service, order-service, and a minimal user-service placeholder. Use Docker Compose for local integration, MySQL for persistence, and Redis for cache plus seckill stock deduction.

**Tech Stack:** Golang, Gin, Gorm, Redis, MySQL 8.0, Docker Compose, MCP protocol

---

## Status
- This plan is intentionally created early and will be expanded stage by stage.
- The current version records confirmed scope and top-level delivery order.

### Task 1: Establish repository skeleton

**Files:**
- Create: `cmd/`
- Create: `internal/`
- Create: `pkg/`
- Create: `deploy/`
- Create: `configs/`
- Create: `scripts/`
- Create: `test/`

**Step 1: Create root project directories**

Create the monorepo layout for multi-service development.

**Step 2: Add common bootstrap files**

Add `go.work`, root `Makefile`, `.env.example`, and base `docker-compose.yml`.

**Step 3: Verify the structure is consistent**

Run commands to verify every expected directory exists and naming is stable.

### Task 2: Build shared infrastructure layer

**Files:**
- Create: `pkg/config/`
- Create: `pkg/logger/`
- Create: `pkg/mysql/`
- Create: `pkg/redis/`
- Create: `pkg/httpx/`
- Create: `pkg/response/`

**Step 1: Write failing tests for config loading and environment parsing**

**Step 2: Implement minimal config and infrastructure bootstrap**

**Step 3: Run tests and keep shared package stable before service work starts**

### Task 3: Build gateway and service skeletons

**Files:**
- Create: `cmd/gateway/`
- Create: `cmd/mcp-gateway/`
- Create: `cmd/product-service/`
- Create: `cmd/cart-service/`
- Create: `cmd/order-service/`
- Create: `cmd/user-service/`

**Step 1: Add each service entry point and health endpoint**

**Step 2: Add internal package structure for handler/service/repo/model**

**Step 3: Verify each service can start independently**

### Task 4: Implement first business chain

**Files:**
- Modify/Create under `internal/product/`
- Modify/Create under `internal/cart/`
- Modify/Create under `internal/order/`

**Step 1: Implement product query path**

**Step 2: Implement cart CRUD path**

**Step 3: Implement create-order and order-detail path**

**Step 4: Add integration tests for the basic purchase flow**

### Task 5: Implement Redis inventory and seckill path

**Files:**
- Create: `internal/order/seckill/`
- Create: `scripts/lua/stock_deduct.lua`

**Step 1: Write failing tests for atomic stock deduction semantics**

**Step 2: Implement Redis Lua deduction path**

**Step 3: Wire order-service to use Redis-first inventory deduction**

**Step 4: Verify no oversell in concurrent test cases**

### Task 6: Implement MCP gateway MVP

**Files:**
- Create: `internal/mcp/`
- Modify: `cmd/mcp-gateway/`

**Step 1: Define MCP tool contracts for stock query and order status query**

**Step 2: Implement service adapters to call product-service and order-service**

**Step 3: Verify an MCP client can discover and call the tools**

### Task 7: Add Docker Compose delivery baseline

**Files:**
- Create: `deploy/docker-compose/`
- Modify: `docker-compose.yml`
- Create: `deploy/mysql/`

**Step 1: Containerize each service**

**Step 2: Add MySQL and Redis containers**

**Step 3: Verify full local startup and smoke test all core chains**

### Task 8: Expand plan into detailed TDD tasks before coding

**Files:**
- Modify: `docs/plans/2026-03-24-mvp-implementation-plan.md`

**Step 1: Refine every top-level task into file-level TDD steps**

**Step 2: Add exact commands, expected failures, and commit boundaries**

**Step 3: Freeze the execution order before implementation begins**
