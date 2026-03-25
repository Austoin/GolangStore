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
