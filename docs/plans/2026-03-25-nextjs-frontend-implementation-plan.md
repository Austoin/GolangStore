# Next.js Frontend Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build a single Next.js frontend project with both shop and admin experiences so the user can visually explore and operate the mall backend through a polished interface.

**Architecture:** Use one frontend project with route groups for `/shop` and `/admin`. Keep layout, shared UI, and data access reusable through `components/`, `lib/types/`, and `lib/adapters/`. Start with mock adapters so the UI can be completed quickly, then leave an adapter seam for later real backend integration.

**Tech Stack:** Next.js, React, TypeScript, CSS modules or global CSS variables, mock adapter layer

---

### Task 1: Initialize frontend project skeleton

**Files:**
- Create: `frontend/package.json`
- Create: `frontend/tsconfig.json`
- Create: `frontend/next.config.js`
- Create: `frontend/app/layout.tsx`
- Create: `frontend/app/globals.css`
- Create: `frontend/app/page.tsx`

**Step 1: Write the failing verification condition**

Define the expected project structure and confirm the frontend directory does not exist yet.

**Step 2: Run verification to confirm missing structure**

Run: `test -d frontend`
Expected: FAIL because the frontend project has not been created yet.

**Step 3: Write minimal implementation**

Create the Next.js project skeleton and base layout files.

**Step 4: Run verification to confirm structure exists**

Run: `test -f frontend/package.json && test -f frontend/app/layout.tsx`
Expected: PASS.

**Step 5: Commit**

```bash
git add frontend docs/process README.md
git commit -m "feat: initialize nextjs frontend skeleton"
```

### Task 2: Build shared design system shell

**Files:**
- Create: `frontend/components/shared/AppHeader.tsx`
- Create: `frontend/components/shared/PrimaryButton.tsx`
- Create: `frontend/components/shared/StatusBadge.tsx`
- Create: `frontend/components/shared/MetricCard.tsx`
- Modify: `frontend/app/globals.css`

**Step 1: Write the failing verification condition**

Define required shared components and verify they do not exist yet.

**Step 2: Run verification to confirm missing files**

Run: `test -f frontend/components/shared/AppHeader.tsx`
Expected: FAIL.

**Step 3: Write minimal implementation**

Create the shared design primitives and global CSS variables.

**Step 4: Run verification to confirm files exist**

Run: `test -f frontend/components/shared/AppHeader.tsx && test -f frontend/components/shared/PrimaryButton.tsx`
Expected: PASS.

**Step 5: Commit**

```bash
git add frontend/components frontend/app/globals.css
git commit -m "feat: add frontend shared design system"
```

### Task 3: Implement shop routes

**Files:**
- Create: `frontend/app/shop/page.tsx`
- Create: `frontend/app/shop/products/[id]/page.tsx`
- Create: `frontend/app/shop/cart/page.tsx`
- Create: `frontend/app/shop/orders/page.tsx`
- Create: `frontend/components/shop/ProductCard.tsx`
- Create: `frontend/components/shop/ProductDetailPanel.tsx`
- Create: `frontend/components/shop/CartSummary.tsx`
- Create: `frontend/components/shop/OrderCard.tsx`

**Step 1: Write the failing verification condition**

Define the required shop routes and confirm they do not exist yet.

**Step 2: Run verification to confirm missing routes**

Run: `test -f frontend/app/shop/page.tsx`
Expected: FAIL.

**Step 3: Write minimal implementation**

Create the four shop routes and their presentation components.

**Step 4: Run verification to confirm route files exist**

Run: `test -f frontend/app/shop/page.tsx && test -f frontend/app/shop/cart/page.tsx`
Expected: PASS.

**Step 5: Commit**

```bash
git add frontend/app/shop frontend/components/shop
git commit -m "feat: add shop frontend routes"
```

### Task 4: Implement admin routes

**Files:**
- Create: `frontend/app/admin/page.tsx`
- Create: `frontend/app/admin/products/page.tsx`
- Create: `frontend/app/admin/inventory/page.tsx`
- Create: `frontend/app/admin/orders/page.tsx`
- Create: `frontend/components/admin/ProductAdminTable.tsx`
- Create: `frontend/components/admin/InventoryTable.tsx`
- Create: `frontend/components/admin/OrderTable.tsx`
- Create: `frontend/components/admin/DashboardMetrics.tsx`

**Step 1: Write the failing verification condition**

Define the required admin routes and confirm they do not exist yet.

**Step 2: Run verification to confirm missing routes**

Run: `test -f frontend/app/admin/page.tsx`
Expected: FAIL.

**Step 3: Write minimal implementation**

Create the admin routes and admin-specific components.

**Step 4: Run verification to confirm route files exist**

Run: `test -f frontend/app/admin/page.tsx && test -f frontend/app/admin/orders/page.tsx`
Expected: PASS.

**Step 5: Commit**

```bash
git add frontend/app/admin frontend/components/admin
git commit -m "feat: add admin frontend routes"
```

### Task 5: Add mock data adapter layer

**Files:**
- Create: `frontend/lib/types/index.ts`
- Create: `frontend/lib/mock/data.ts`
- Create: `frontend/lib/adapters/mockAdapter.ts`
- Create: `frontend/lib/adapters/index.ts`
- Modify: relevant shop/admin pages to consume adapter functions

**Step 1: Write the failing verification condition**

Define the adapter layer files and confirm they do not exist yet.

**Step 2: Run verification to confirm missing files**

Run: `test -f frontend/lib/adapters/mockAdapter.ts`
Expected: FAIL.

**Step 3: Write minimal implementation**

Create the frontend types, mock data, and adapter facade used by shop and admin pages.

**Step 4: Run verification to confirm adapter files exist**

Run: `test -f frontend/lib/adapters/mockAdapter.ts && test -f frontend/lib/types/index.ts`
Expected: PASS.

**Step 5: Commit**

```bash
git add frontend/lib frontend/app
git commit -m "feat: add frontend mock adapter layer"
```

### Task 6: Add frontend run instructions and close documentation loop

**Files:**
- Modify: `README.md`
- Modify: `docs/process/stage-12-端到端业务闭环收束.md`

**Step 1: Record frontend implementation status**

Update docs with the frontend structure, routes, and current adapter mode.

**Step 2: Add frontend local run commands**

Document how to run the Next.js frontend locally.

**Step 3: Commit**

```bash
git add README.md docs/process/stage-12-端到端业务闭环收束.md
git commit -m "docs: record frontend implementation status"
```
