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
