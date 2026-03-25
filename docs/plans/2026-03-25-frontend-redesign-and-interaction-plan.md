# Frontend Redesign And Interaction Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Redesign the current frontend into a polished, interactive admin-first experience followed by a rebuilt shop experience, with real navigation and core actions usable end to end.

**Architecture:** Keep one Next.js project with `/admin` and `/shop` routes, but rebuild page structure and shared design primitives for a stronger visual identity. Use real backend data where already available and preserve adapter seams so the UI remains maintainable while becoming actually clickable and operational.

**Tech Stack:** Next.js, React, TypeScript, CSS variables, server components, HTTP adapter layer

---

### Task 1: Redesign shared shell and design tokens

**Files:**
- Modify: `frontend/app/globals.css`
- Modify: `frontend/components/shared/AppHeader.tsx`
- Modify: `frontend/components/shared/PrimaryButton.tsx`
- Modify: `frontend/components/shared/StatusBadge.tsx`
- Modify: `frontend/components/shared/MetricCard.tsx`

**Step 1: Write the failing verification condition**

Define the new visual shell requirements and confirm the current shell still reflects the old minimal layout.

**Step 2: Run verification to confirm old shell is still present**

Run: `curl -I http://127.0.0.1:3000/admin`
Expected: PASS for the old page, proving redesign has not happened yet.

**Step 3: Write minimal implementation**

Replace the current shell with the new admin-first visual language and improved shared components.

**Step 4: Run verification to confirm new shell still renders**

Run: `curl -I http://127.0.0.1:3000/admin`
Expected: PASS.

### Task 2: Redesign admin pages with real data panels

**Files:**
- Modify: `frontend/app/admin/page.tsx`
- Modify: `frontend/app/admin/products/page.tsx`
- Modify: `frontend/app/admin/inventory/page.tsx`
- Modify: `frontend/app/admin/orders/page.tsx`
- Modify: `frontend/components/admin/*.tsx`

**Step 1: Write the failing verification condition**

Confirm the current admin pages are still the older placeholder-like layout.

**Step 2: Run verification**

Run: `curl -I http://127.0.0.1:3000/admin/products && curl -I http://127.0.0.1:3000/admin/orders`
Expected: PASS for current pages before redesign.

**Step 3: Write minimal implementation**

Rebuild admin routes into a stronger operating dashboard style while preserving real data sources.

**Step 4: Run verification**

Run: `curl -I http://127.0.0.1:3000/admin && curl -I http://127.0.0.1:3000/admin/inventory`
Expected: PASS.

### Task 3: Redesign shop pages with real interactions

**Files:**
- Modify: `frontend/app/shop/page.tsx`
- Modify: `frontend/app/shop/products/[id]/page.tsx`
- Modify: `frontend/app/shop/cart/page.tsx`
- Modify: `frontend/app/shop/orders/page.tsx`
- Modify: `frontend/components/shop/*.tsx`

**Step 1: Write the failing verification condition**

Confirm the current shop pages are still minimal placeholders.

**Step 2: Run verification**

Run: `curl -I http://127.0.0.1:3000/shop && curl -I http://127.0.0.1:3000/shop/cart`
Expected: PASS for old version.

**Step 3: Write minimal implementation**

Rebuild shop pages so actions are obvious, buttons are clickable, and navigation feels intentional.

**Step 4: Run verification**

Run: `curl -I http://127.0.0.1:3000/shop/products/101 && curl -I http://127.0.0.1:3000/shop/orders`
Expected: PASS.

### Task 4: Complete click-through interactions

**Files:**
- Modify relevant routes/components under `frontend/app/shop/` and `frontend/app/admin/`
- Modify `frontend/lib/adapters/httpAdapter.ts` if needed

**Step 1: Write the failing verification condition**

List the currently non-clickable or incomplete actions.

**Step 2: Run verification**

Verify current routes exist but key flows are not fully interactive.

**Step 3: Write minimal implementation**

Make “查看商品”, “加入购物车”, “提交订单”, and admin page navigation clearly clickable and operational.

**Step 4: Run verification**

Use page access checks and the existing backend demo flow to confirm interactions still function.

### Task 5: Final visual and runtime verification

**Files:**
- Modify: `README.md`
- Modify: `docs/process/stage-12-端到端业务闭环收束.md`

**Step 1: Update documentation**

Record the redesign completion and runtime instructions.

**Step 2: Run verification**

Run: `bash scripts/run-all.sh && bash scripts/demo.sh`
Expected: PASS.

**Step 3: Verify frontend entrypoints**

Run: `curl -I http://127.0.0.1:3000/ && curl -I http://127.0.0.1:3000/shop && curl -I http://127.0.0.1:3000/admin`
Expected: PASS.
