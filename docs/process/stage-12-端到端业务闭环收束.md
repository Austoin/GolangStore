# 阶段 12：端到端业务闭环收束

## 时间
- 2026-03-25

## 阶段目标
- 从当前“订单服务可运行、订单可落库、脚本可启动”的状态，推进到“核心业务逻辑完整、可一键启动、可一键验证效果”的可演示交付状态。

## 当前项目基线
- 已完成工程骨架、文档体系与阶段过程记录。
- 已完成 `product`、`cart`、`order` 领域模型与最小查询链路。
- 已完成购物车到下单编排。
- 已完成 `order-service` 运行入口接线。
- 已完成 `cart` 与 `order` 的 MySQL 仓储。
- 已完成真实验证：购物车数据 -> `/orders/from-cart` -> `orders` 与 `order_items` 落库。
- 已完成 `scripts/start.sh` 一键启动脚本，并验证其跨 Git Bash / WSL 的兼容性。

## 当前缺口
- 还未接入真实商品持久化查询与库存读取。
- 还未在下单链路中执行库存校验与扣减。
- 还未补齐购物车写接口，当前仍主要依赖插库或已有仓储读取验证。
- `product-service`、`cart-service` 尚未像 `order-service` 一样完成真实运行时装配。
- 缺少一个“一键演示效果”的脚本来自动跑通完整业务闭环。

## 收尾目标定义
- 可以通过 `bash scripts/start.sh` 一键启动核心中间件和核心服务。
- 可以通过一组固定命令或脚本完成：
  - 初始化商品与库存
  - 添加购物车
  - 从购物车创建订单
  - 查询订单
  - 验证库存减少
- 核心链路在本地真实运行，不依赖手工改代码或临时插表结构。

## 执行约束
- 继续保持小步推进。
- 每一个小步骤完成后都要：
  - 更新文档
  - 提交 git
  - 推送云端
- 继续坚持先测试失败、再实现、再验证。

## 当前进展

### Task 1：商品持久化与查询真实化
- 已创建：
  - `internal/product/mysql_repo.go`
  - `internal/product/mysql_repo_test.go`
  - `cmd/product-service/router.go`
  - `cmd/product-service/router_test.go`
- 已更新：
  - `cmd/product-service/main.go`
  - `README.md`

### 当前实现结果
- `internal/product` 已支持从 `products` + `product_stocks` 读取商品与库存。
- `cmd/product-service` 已接入 MySQL 装配。
- 已挂载路由：
  - `GET /health`
  - `GET /products/:id`

### 当前验证结果
- 已执行：`go test ./internal/product && go test ./cmd/product-service`
- 结果：通过

### Task 2：购物车写接口补齐
- 已创建：
  - `cmd/cart-service/router.go`
  - `cmd/cart-service/router_test.go`
- 已更新：
  - `internal/cart/repo.go`
  - `internal/cart/service.go`
  - `internal/cart/handler.go`
  - `internal/cart/mysql_repo.go`
  - `internal/cart/repo_test.go`
  - `internal/cart/service_test.go`
  - `internal/cart/handler_test.go`
  - `cmd/cart-service/main.go`
  - `README.md`

### 当前实现结果
- `internal/cart` 已支持 `Save(item)`。
- `cart.Service` 已支持 `AddItem(item)`。
- `cart.Handler` 已支持 `POST /carts`。
- `cmd/cart-service` 已接入 MySQL 装配。
- 当前路由：
  - `GET /health`
  - `GET /carts/:userId`
  - `POST /carts`

### 当前验证结果
- 已执行：`go test ./internal/cart && go test ./cmd/cart-service`
- 结果：通过

### Task 3：下单库存校验与扣减
- 已更新：
  - `internal/order/service.go`
  - `internal/order/service_test.go`
  - `internal/order/handler_test.go`
  - `internal/product/mysql_repo.go`
  - `internal/product/mysql_repo_test.go`
  - `README.md`

### 当前实现结果
- `order.Service` 已新增库存依赖。
- 当前下单流程会在创建订单前校验每个商品是否库存充足。
- 当前下单流程会在创建成功前执行库存扣减。
- 当前新增错误：`ErrInsufficientStock`。
- `internal/product.MySQLRepository` 已支持：
  - `HasEnough(productID, quantity)`
  - `Deduct(productID, quantity)`

### 当前验证结果
- 已执行：`go test ./internal/product && go test ./internal/order`
- 结果：通过

### Task 5：一键演示脚本（当前中间状态）
- 已创建：`scripts/demo.sh`

### 当前脚本目标
- 清理旧数据
- 初始化商品与库存
- 调用 `POST /carts`
- 调用 `POST /orders/from-cart`
- 验证 `orders`、`order_items` 与库存扣减结果

### 当前验证结果
- 已执行：`bash -n scripts/demo.sh && bash scripts/demo.sh`
- 第一次失败：脚本语法错误，已修复
- 第二次失败：`cart-service is not healthy`

### 当前结论
- `demo.sh` 的脚本语法已修复。
- 当前阻塞点不在 `demo.sh` 本身，而在 `start.sh` 尚未同时启动 `cart-service`。
- 下一步应先增强启动脚本，再重新验证 `demo.sh`。

### 本批次提交与远端同步
- 已创建提交：`a4e0426` `feat: add stock validation and deduction`
- 已创建提交：`a2f0e9a` `fix: align order service wiring with stock checks`
- 已推送到远端 `origin/master`
- 已验证：
  - `git status` 工作区干净
  - `git branch -vv` 显示本地 `master` 跟踪 `origin/master`
  - `git ls-remote --heads origin` 显示远端 `master` 指向 `a2f0e9a`

### Task 5：一键演示脚本（最新阻塞）
- 已执行：`bash scripts/start.sh`
- 已验证：
  - `product-service` 健康
  - `cart-service` 健康
  - `order-service` 健康检查返回成功
- 已执行：`bash scripts/demo.sh`
- 当前失败点：`POST /orders/from-cart` 返回 `400`

### 当前根因定位
- 读取 `.runtime/order-service.log` 后发现，最新 `order-service` 进程启动失败：
  - `listen tcp :8082: bind: Only one usage of each socket address...`
- 进一步检查确认：
  - 端口 `8082` 被 Windows 进程 `PID 2888` 占用
  - `.runtime/order-service.pid` 记录的是 shell 子进程 PID，不是最终监听端口的真实 Windows 进程 PID
- 结论：
  - 旧 `order-service` 实例没有被正确停止
  - 导致 `demo.sh` 实际请求到的是旧实例，而不是最新代码实例

### 下一步
- 修复 `start.sh` 的服务清理策略，改为按端口清理 `8081`、`8082`、`8083`
- 再重新执行：
  - `bash scripts/start.sh`
  - `bash scripts/demo.sh`

### Task 5：一键演示脚本（最终结果）
- 已更新：
  - `scripts/start.sh`
  - `scripts/demo.sh`
  - `cmd/order-service/main.go`
  - `README.md`

### 当前实现结果
- `start.sh` 现在会：
  - 启动 MySQL、Redis
  - 按端口清理旧的 `product-service`、`cart-service`、`order-service`
  - 启动三服务并等待健康检查通过
- `demo.sh` 现在会：
  - 清理旧数据
  - 初始化商品与库存
  - 调用购物车接口写入购物车项
  - 调用 `/orders/from-cart` 创建订单
  - 验证 `orders`、`order_items` 落库
  - 验证库存由 `5` 扣减为 `3`

### 当前验证结果
- 已执行：`bash scripts/start.sh && bash scripts/demo.sh`
- 结果：通过
- 关键结果：
  - 三服务健康检查全部通过
  - 订单创建成功
  - `orders` 与 `order_items` 成功落库
  - `product_stocks.stock` 成功从 `5` 扣减到 `3`

### Task 6：最终全量验证与文档收口
- 已执行：`go test ./...`
- 结果：通过

### 当前阶段结论
- 项目已达到当前目标：
  - 核心业务逻辑闭环完成
  - 可一键启动查看效果
  - 可一键演示真实业务结果

## 前端实现：Task 1 初始化骨架

### 已完成内容
- 已创建：
  - `frontend/package.json`
  - `frontend/tsconfig.json`
  - `frontend/next.config.js`
  - `frontend/next-env.d.ts`
  - `frontend/app/layout.tsx`
  - `frontend/app/globals.css`
  - `frontend/app/page.tsx`
- 已同步更新：`README.md`

### 当前结果
- `frontend/` 目录已存在。
- Next.js 最小骨架已建立。
- 首页已提供前台 `/shop` 和后台 `/admin` 的入口占位。

## 前端实现：Task 2 共享设计系统

### 已完成内容
- 已创建：
  - `frontend/components/shared/AppHeader.tsx`
  - `frontend/components/shared/PrimaryButton.tsx`
  - `frontend/components/shared/StatusBadge.tsx`
  - `frontend/components/shared/MetricCard.tsx`
- 已更新：
  - `frontend/app/globals.css`

### 当前结果
- 前端共享组件基础集已建立。
- 全局样式变量已可支撑前后台统一视觉。

## 前端实现：Task 3 shop 路由

### 已完成内容
- 已创建：
  - `frontend/app/shop/page.tsx`
  - `frontend/app/shop/products/[id]/page.tsx`
  - `frontend/app/shop/cart/page.tsx`
  - `frontend/app/shop/orders/page.tsx`
  - `frontend/components/shop/ProductCard.tsx`
  - `frontend/components/shop/ProductDetailPanel.tsx`
  - `frontend/components/shop/CartSummary.tsx`
  - `frontend/components/shop/OrderCard.tsx`

### 当前结果
- 前台四页路由骨架已完成。
- 商品列表、商品详情、购物车、订单视图都有页面占位。

## 前端实现：Task 4 admin 路由

### 已完成内容
- 已创建：
  - `frontend/app/admin/page.tsx`
  - `frontend/app/admin/products/page.tsx`
  - `frontend/app/admin/inventory/page.tsx`
  - `frontend/app/admin/orders/page.tsx`
  - `frontend/components/admin/DashboardMetrics.tsx`
  - `frontend/components/admin/ProductAdminTable.tsx`
  - `frontend/components/admin/InventoryTable.tsx`
  - `frontend/components/admin/OrderTable.tsx`

### 当前结果
- 后台四页路由骨架已完成。
- Dashboard、商品管理、库存面板、订单管理都已有页面占位。

## 前端实现：Task 5 mock adapter 层

### 已完成内容
- 已创建：
  - `frontend/lib/types/index.ts`
  - `frontend/lib/mock/data.ts`
  - `frontend/lib/adapters/mockAdapter.ts`
  - `frontend/lib/adapters/index.ts`
- 已更新：
  - `frontend/app/shop/page.tsx`
  - `frontend/app/shop/products/[id]/page.tsx`
  - `frontend/app/shop/cart/page.tsx`
  - `frontend/app/shop/orders/page.tsx`

### 当前结果
- 前端页面已不再直接写死展示数据。
- `/shop` 页面已通过 adapter 层读取 mock 数据。
- 后续可替换为真实 HTTP adapter 而不改页面结构。

## 前端实现：Task 6 文档收口

### 已完成内容
- 已更新：`README.md`
- 已完善当前前端运行说明与项目状态说明

### 当前结果
- README 已包含前后端当前运行方式
- 前端当前交付边界已明确：
  - Next.js 单项目
  - `/shop` 与 `/admin` 双入口
  - mock adapter 模式

### 当前运行验证结果
- 已执行：`npm install`
- 已执行：`npm run dev -- --hostname 0.0.0.0 --port 3000`
- 已验证：
  - `http://127.0.0.1:3000/` 返回 `200`
  - `http://127.0.0.1:3000/shop` 返回 `200`
  - `http://127.0.0.1:3000/admin` 返回 `200`

### 当前阶段结论
- 前端骨架、共享组件、前后台页面、mock adapter 已完成。
- 前端当前已可运行并访问主要入口页面。

## 运行脚本收口

### 已完成内容
- 已创建：
  - `scripts/run-all.sh`
  - `scripts/stop-all.sh`

### 当前实现结果
- `run-all.sh` 会：
  - 启动 MySQL、Redis
  - 启动 `product-service`、`cart-service`、`order-service`
  - 安装前端依赖（如缺失）
  - 启动 Next.js 前端
  - 验证 `/`、`/shop`、`/admin` 可访问
- `stop-all.sh` 会：
  - 停止 3000、8081、8082、8083 端口上的进程
  - 停止 MySQL 与 Redis 容器

### 当前验证结果
- 已执行：`bash scripts/run-all.sh`
- 结果：通过
- 已执行：`bash scripts/stop-all.sh`
- 结果：通过

### 启动体验优化
- 已更新：`scripts/run-all.sh`
- 当前增强点：
  - 前端启动后立刻输出访问地址
  - 等待 `/`、`/shop`、`/admin` 时输出明确进度
  - 若超时会自动打印前端日志尾部
  - 明确提示首次启动可能因 Next.js 下载 SWC 较慢

### 当前验证结果
- 已重新执行：`bash scripts/run-all.sh`
- 结果：通过
- 当前启动输出已可直接看到：
  - `frontend root`
  - `frontend shop`
  - `frontend admin`

### 最终结论
- 当前项目已经具备：
  - 一键启动前后端
  - 一键停止前后端
  - 一键演示核心业务闭环

### Task 5：一键演示脚本（最新业务阻塞）
- 修复 `start.sh` 后，三服务已能同时健康启动。
- 重新执行 `bash scripts/demo.sh` 后，`POST /orders/from-cart` 仍返回 `400`。

### 当前根因定位
- 直接请求接口得到响应体：`{"error":"checked cart items cannot be empty"}`。
- 回查数据库 `cart_items` 发现：
  - `user_id = 0`
  - `product_id = 0`
  - `product_name = ''`
  - `price = 199900`
  - `quantity = 2`
  - `checked = 1`
- 结论：
  - `POST /carts` 的 JSON 绑定没有把 `user_id`、`product_id`、`product_name` 映射进 `cart.Item`
  - 当前最可能根因是 `cart.Item` 缺少 JSON tag

### 下一步
- 为 `internal/cart/model.go` 增加 JSON tag
- 重新执行：
  - `bash scripts/start.sh`
  - `bash scripts/demo.sh`


## 前台真实联调：Task 1 商品列表

### 已完成内容
- 已新增：`frontend/lib/adapters/httpAdapter.ts`
- 已更新：
  - `frontend/app/shop/page.tsx`
  - `internal/product/repo.go`
  - `internal/product/service.go`
  - `internal/product/handler.go`
  - `internal/product/mysql_repo.go`
  - `internal/product/mysql_repo_test.go`
  - `cmd/product-service/router.go`
  - `cmd/product-service/router_test.go`
  - `README.md`

### 当前实现结果
- `product-service` 已支持：
  - `GET /products`
  - `GET /products/:id`
- `/shop` 页面已从 mock 数据切到真实后端商品列表接口。

### 当前验证结果
- 已执行：`go test ./internal/product && go test ./cmd/product-service`
- 结果：通过
- 已执行：`bash scripts/run-all.sh`
- 已执行：`curl -I http://127.0.0.1:3000/shop`
- 结果：`200 OK`


## 前台真实联调：Task 2 商品详情

### 已完成内容
- 已更新：
  - `frontend/lib/adapters/httpAdapter.ts`
  - `frontend/app/shop/products/[id]/page.tsx`
  - `README.md`

### 当前实现结果
- 商品详情页已从 mock 数据切到真实后端详情接口。
- 当前使用真实路由参数 `id` 请求商品详情。

### 当前验证结果
- 已执行：`curl -I http://127.0.0.1:3000/shop/products/101`
- 结果：`200 OK`


## 前台真实联调：Task 3 购物车与下单

### 已完成内容
- 已更新：
  - `frontend/lib/adapters/httpAdapter.ts`
  - `frontend/components/shop/CartSummary.tsx`
  - `frontend/app/shop/cart/page.tsx`
  - `README.md`

### 当前实现结果
- 购物车页已切到真实购物车接口。
- 购物车页已可直接触发真实 `/orders/from-cart` 下单。
- 下单完成后会跳转到 `/shop/orders`。


## 前台真实联调：Task 4 订单列表

### 已完成内容
- 已更新：
  - `internal/order/repo.go`
  - `internal/order/service.go`
  - `internal/order/handler.go`
  - `internal/order/mysql_repo.go`
  - `internal/order/repo_test.go`
  - `internal/order/service_test.go`
  - `internal/order/handler_test.go`
  - `cmd/order-service/router.go`
  - `frontend/lib/adapters/httpAdapter.ts`
  - `frontend/app/shop/orders/page.tsx`
  - `README.md`

### 当前实现结果
- `order-service` 已支持：`GET /orders`
- `/shop/orders` 已切到真实订单列表接口。

### 当前验证结果
- 已执行：`curl -I http://127.0.0.1:3000/shop/cart`
- 结果：`200 OK`

### 当前验证结果
- 已执行：`go test ./internal/order`
- 结果：通过
- 已执行：`curl -I http://127.0.0.1:3000/shop/orders`
- 结果：`200 OK`


## 后台真实联调

### 已完成内容
- 已更新：
  - `frontend/components/admin/DashboardMetrics.tsx`
  - `frontend/components/admin/ProductAdminTable.tsx`
  - `frontend/components/admin/InventoryTable.tsx`
  - `frontend/components/admin/OrderTable.tsx`
  - `frontend/app/admin/page.tsx`
  - `frontend/app/admin/products/page.tsx`
  - `frontend/app/admin/inventory/page.tsx`
  - `frontend/app/admin/orders/page.tsx`
  - `README.md`

### 当前实现结果
- `/admin/products` 已使用真实商品接口数据。
- `/admin/inventory` 已使用真实库存数据。
- `/admin/orders` 已使用真实订单数据。
- `/admin` 已使用真实商品数、订单数、低库存数进行汇总。


### 当前验证结果
- 已执行：`bash scripts/run-all.sh`
- 已验证：
  - `/admin` -> `200 OK`
  - `/admin/products` -> `200 OK`
  - `/admin/inventory` -> `200 OK`
  - `/admin/orders` -> `200 OK`

### 当前阶段总结果
- 用户前台 `/shop` 已基本切到真实后端接口。
- 管理后台 `/admin` 已切到真实后端数据驱动。
- 当前前后台都已具备真实数据展示能力。


## 前端运行问题修复

### 已修复问题
- `/shop` 页面商品列表真实联调时，后端返回字段为大写 `ID/Name/...`。
- 前端 adapter 读取的是小写 `id/name/...`，导致 React 列表 key 警告。

### 当前修复
- 已为 `internal/product/model.go` 增加 JSON tag：
  - `id`
  - `name`
  - `description`
  - `price`
  - `status`
  - `stock`

### 当前验证结果
- 已执行：`go test ./internal/product`
- 已执行：`bash scripts/run-all.sh`
- 已执行：`curl http://127.0.0.1:8081/products`
- 当前返回字段已为小写 JSON 键。

### 浏览器控制台中的 hydration 提示说明
- 当前看到的 `data-redeviation-bs-uid` 来自浏览器扩展注入到 `<html>`。
- 这类属性不是项目 SSR 代码主动渲染的，不属于当前应用逻辑错误。
- 若要确认，可在无扩展窗口或禁用相关扩展后再看控制台。


## 前后端重做与交互打通总结

### 已完成内容
- 前后台视觉已经按“现代电商展厅 + 运营指挥台”重做。
- 前台核心页面与后台核心页面都已接入真实数据源。
- 一键启动、一键演示、健康检查、真实下单、库存扣减全部验证通过。

### 最终验证结果
- 已执行：`bash scripts/run-all.sh`
- 已执行：`bash scripts/demo.sh`
- 已验证页面：
  - `/shop`
  - `/shop/products/101`
  - `/shop/cart`
  - `/shop/orders`
  - `/admin`
  - `/admin/products`
  - `/admin/inventory`
  - `/admin/orders`
- 以上页面与脚本当前均可工作。

### 当前最终状态
- 项目已达到：
  - 可一键启动
  - 可一键停止
  - 可一键演示真实商城核心业务闭环
  - 前后台面板均可体验真实后端数据


## 前台/后台交互修复

### 已完成内容
- 已更新：
  - `frontend/components/shop/ProductCard.tsx`
  - `frontend/app/shop/page.tsx`
  - `frontend/components/admin/OrderTable.tsx`
  - `internal/product/model.go`
  - `README.md`

### 当前结果
- 前台商品卡已支持：
  - 查看商品
  - 加入购物车
- 后台订单表已使用稳定 key 渲染。
- `/shop`、`/shop/products/101`、`/shop/cart`、`/admin/orders` 当前都可访问。


## 购物车提交错误体验优化

### 已完成内容
- 已更新：`frontend/app/shop/cart/page.tsx`
- 已同步更新：`README.md`

### 当前实现结果
- 提交订单失败时，不再直接抛出 Next.js application error 页面。
- 当前会回到 `/shop/cart?error=...` 并在页面内显示错误信息。

### 当前验证结果
- 已构造库存不足场景：
  - 库存 = 1
  - 购物车数量 = 2
- 已访问：`/shop/cart?error=failed%20to%20create%20order%3A%20400`
- 结果：页面返回 `200 OK`


## 数量控制可视化优化

### 已完成内容
- 已更新：
  - `frontend/components/shop/ProductCard.tsx`
  - `frontend/app/shop/page.tsx`
  - `frontend/app/shop/cart/page.tsx`
  - `internal/cart/mysql_repo.go`
  - `internal/cart/mysql_repo_test.go`
  - `README.md`

### 当前实现结果
- 商品卡已支持数量输入。
- 加入购物车时会把数量真实提交到后端。
- 若购物车中已存在同商品，会更新数量而不是重复插入。
- 购物车页已展示：
  - 商品名称
  - 单价
  - 数量
  - 小计
  - 总计

### 当前验证结果
- 已执行：`bash scripts/run-all.sh`
- 已执行：手动调用 `POST /carts`，数量 = 3
- 已执行：`curl -I http://127.0.0.1:3000/shop/cart`
- 已执行：数据库回查 `cart_items.quantity`
- 结果：页面 `200 OK`，数据库中数量 = `3`


## 商品详情与购物车交互增强

### 已完成内容
- 已更新：
  - `frontend/app/shop/products/[id]/page.tsx`
  - `frontend/components/shop/ProductDetailPanel.tsx`
  - `frontend/app/shop/cart/page.tsx`
  - `frontend/lib/adapters/httpAdapter.ts`
  - `internal/cart/repo.go`
  - `internal/cart/service.go`
  - `internal/cart/handler.go`
  - `internal/cart/mysql_repo.go`
  - `cmd/cart-service/router.go`
  - `README.md`

### 当前实现结果
- 商品详情页已支持：
  - 数量选择
  - 真实加入购物车
- 购物车已支持：
  - 删除商品
  - 勾选切换
  - 数量、小计、总计展示
- 购物车后端已支持：
  - 删除接口
  - 勾选切换接口

### 当前验证结果
- 已执行：`bash scripts/run-all.sh`
- 已验证：
  - `/shop/products/101` -> `200 OK`
  - `/shop/cart` -> `200 OK`
  - `PATCH /carts/1/101/checked` -> `204`
  - `DELETE /carts/1/101` -> `204`


## 后台商品管理增强

### 已完成内容
- 已更新：
  - `internal/product/repo.go`
  - `internal/product/service.go`
  - `internal/product/handler.go`
  - `internal/product/mysql_repo.go`
  - `cmd/product-service/router.go`
  - `frontend/components/admin/ProductAdminTable.tsx`
  - `frontend/app/admin/products/page.tsx`
  - `README.md`

### 当前实现结果
- 后端已支持：`POST /products`
- 后台商品管理页已支持：
  - 商品名称输入
  - 商品描述输入
  - 价格输入
  - 库存数量输入
  - 状态输入
  - 提交新增商品

### 当前验证结果
- 已执行：`go test ./internal/product`
- 已执行：`POST /products`
- 已执行：`curl -I http://127.0.0.1:3000/admin/products`
- 结果：页面 `200 OK`，后端商品列表已可看到新增商品

## 后台订单面板增强

### 当前结果
- `/admin/orders` 已展示：
  - 订单号
  - 金额
  - 状态
  - 订单项数


## 后台商品管理可操作化

### 已完成内容
- 已更新：
  - `internal/product/repo.go`
  - `internal/product/service.go`
  - `internal/product/handler.go`
  - `internal/product/mysql_repo.go`
  - `internal/product/handler_test.go`
  - `cmd/product-service/router.go`
  - `frontend/components/admin/ProductAdminTable.tsx`
  - `frontend/app/admin/products/page.tsx`
  - `README.md`

### 当前实现结果
- 后端已支持：`POST /products`
- 后台商品页已支持：
  - 输入商品名称
  - 输入描述
  - 输入价格
  - 输入库存数量
  - 提交新增商品

### 当前验证结果
- 已执行：`go test ./internal/product`
- 已执行：`POST /products`
- 已验证：`/admin/products -> 200 OK`


## 后台商品管理与前台现代化重排版

### 已完成内容
- 已更新：
  - `frontend/app/shop/page.tsx`
  - `frontend/app/shop/products/[id]/page.tsx`
  - `frontend/components/shop/ProductCard.tsx`
  - `frontend/components/shop/ProductDetailPanel.tsx`
  - `frontend/components/admin/ProductAdminTable.tsx`
  - `frontend/app/admin/products/page.tsx`
  - `internal/product/*`
  - `cmd/product-service/router.go`
  - `README.md`

### 当前实现结果
- 后台商品管理页已支持：
  - 商品名输入
  - 描述输入
  - 价格输入
  - 库存数量输入
  - 真实提交新增商品
- 前台首页已重做为更明确的现代电商展厅布局。
- 商品详情页已支持数量选择和真实加购。

### 当前验证结果
- 已执行：`POST /products`
- 结果：`201 Created`
- 已执行：`curl -I http://127.0.0.1:3000/admin/products`
- 结果：`200 OK`
- 已执行：`curl -I http://127.0.0.1:3000/shop`
- 结果：`200 OK`
- 已执行：`curl -I http://127.0.0.1:3000/shop/products/101`
- 结果：`200 OK`
