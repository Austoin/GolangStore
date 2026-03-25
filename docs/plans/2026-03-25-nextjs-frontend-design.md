# Next.js 前后台一体可视化面板设计

## 目标

为当前 Go 微服务商城项目设计一个可视化前端面板，采用一个 Next.js 项目，同时提供：

- 用户前台体验入口
- 管理后台入口

该前端第一阶段以“界面优先，接口占位”为原则，但数据结构、页面骨架、交互流程都必须与当前后端模型保持一致，确保后续可以平滑切换到真实接口联调。

## 设计原则

### 1. 一个项目，两套入口

采用单一 Next.js 项目，不拆两个独立前端。原因：

- 当前项目阶段更适合快速搭建统一体验
- 可复用设计系统、请求层、类型定义和共享组件
- 后续维护成本更低

入口划分：

- `/shop`：用户前台
- `/admin`：管理后台

### 2. 界面优先，但不脱离真实模型

虽然第一阶段先不强制全量真实联调，但不能使用脱离后端现实的数据结构。所有前端 mock 数据、页面设计、状态卡片和表格都必须围绕当前后端已有领域模型展开：

- `Product`
- `CartItem`
- `Order`
- `OrderItem`
- `Stock`

### 3. 统一设计系统，双重气质

前台和后台不做完全割裂的视觉语言，而是使用同一设计系统下的双重表达：

- 前台：更强调品牌感、商品展示、购买体验
- 后台：更强调密度、控制感、状态信息和操作感

## 站点信息架构

### 用户前台

- `/shop`
  - 商品首页
  - 商品卡片网格
  - 快速加入购物车

- `/shop/products/[id]`
  - 商品详情
  - 价格、库存、状态
  - 加入购物车按钮

- `/shop/cart`
  - 购物车项
  - 数量、小计、总计
  - 去下单按钮

- `/shop/orders`
  - 订单列表
  - 订单状态
  - 订单金额

### 管理后台

- `/admin`
  - 总览 Dashboard
  - 商品总数、订单总数、低库存数量、服务状态

- `/admin/products`
  - 商品管理列表
  - 名称、价格、状态、库存

- `/admin/inventory`
  - 库存面板
  - 低库存高亮

- `/admin/orders`
  - 订单管理页
  - 订单号、用户、金额、状态、订单项数量

## 路由与布局策略

### 顶层布局

根布局提供全局样式、主题变量和顶部入口切换。

建议入口切换方式：

- 顶部 Header 内提供 `Shop` / `Admin` 切换
- 保留项目品牌标识和当前运行状态提示

### 前台布局

前台采用顶部导航 + 内容流布局：

- Header
- 商品主内容区
- 页面底部状态信息（可选）

### 后台布局

后台采用左侧导航 + 右侧内容工作区：

- 左侧导航：Dashboard / Products / Inventory / Orders
- 右侧内容区：表格、详情、状态卡片

## 推荐目录结构

```text
frontend/
  app/
    shop/
      page.tsx
      products/[id]/page.tsx
      cart/page.tsx
      orders/page.tsx
    admin/
      page.tsx
      products/page.tsx
      inventory/page.tsx
      orders/page.tsx
    layout.tsx
    globals.css
  components/
    shared/
    shop/
    admin/
  lib/
    adapters/
    api/
    mock/
    types/
    utils/
  public/
```

## 数据访问策略

### adapter 模式

前端页面不能直接依赖 mock 文件，也不能直接硬编码请求 URL。统一通过 adapter 抽象层访问数据。

例如：

- `listProducts()`
- `getProductDetail(id)`
- `listCartItems(userId)`
- `addCartItem(payload)`
- `createOrderFromCart(userId)`
- `listOrders()`

适配层分两类：

- `mock adapter`：当前阶段默认实现
- `http adapter`：后续真实联调替换实现

## 第一版共享组件范围

### shared

- `AppHeader`
- `StatusBadge`
- `MetricCard`
- `DataTable`
- `EmptyState`
- `SectionTitle`
- `PrimaryButton`

### shop

- `ProductCard`
- `ProductDetailPanel`
- `CartSummary`
- `OrderCard`

### admin

- `InventoryTable`
- `OrderTable`
- `ProductAdminTable`
- `DashboardMetrics`

## 第一版必须包含的交互

只保留四个最核心动作：

1. 浏览商品
2. 加入购物车
3. 从购物车生成订单
4. 在后台查看订单与库存结果

## 视觉方向

### 总体方向

采用统一的现代电商控制台风格，但避免常见模板化后台：

- 前台更偏品牌感、商品陈列感
- 后台更偏控制台感、运维感

### 前台视觉

- 更大的商品卡片
- 清晰价格层级
- 库存标签突出
- CTA 更具引导性

### 后台视觉

- 更紧凑的网格与表格
- 状态色明确
- 指标卡清晰
- 低库存和异常状态高亮

## 第一阶段技术边界

这一版暂不强求：

- 登录鉴权
- SSR 数据抓取优化
- SEO
- 复杂表单校验
- 后台编辑/删除完整流程

优先保证：

- 页面结构完整
- 路由清晰
- 数据模型对齐后端
- 后续可快速切换到真实接口

## 实施顺序建议

1. 初始化 Next.js 项目骨架
2. 搭建全局 layout 和设计变量
3. 完成 `/shop` 四页
4. 完成 `/admin` 四页
5. 接入 mock adapter
6. 预留 HTTP adapter 接口层

## 后续衔接

此设计文档确认后，下一步应生成详细实施计划文档，用于真正开始前端实现：

- `docs/plans/2026-03-25-nextjs-frontend-implementation-plan.md`
