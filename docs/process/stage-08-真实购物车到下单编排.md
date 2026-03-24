# 阶段 08：真实购物车到下单编排

## 时间
- 2026-03-24

## 阶段目标
- 让 `/orders/from-cart` 不再依赖请求体直接传购物车项。
- 改为只接收 `user_id`，并由订单服务主动读取购物车已勾选项完成下单。

## 当前计划顺序
1. 在订单应用服务中引入购物车查询依赖
2. 将编排逻辑改为按用户读取已勾选购物车项
3. 收缩 HTTP 请求体，仅保留 `user_id`
4. 补测试并验证入口行为
5. 更新文档、提交并推送

## 当前约束
- 当前阶段仍使用内存购物车仓储和内存订单仓储。
- 当前阶段不接入 MySQL。
- 当前阶段不接入 Redis 库存扣减。

## 当前状态
- 阶段刚开始，`/orders/from-cart` 仍依赖请求体直接传购物车项。

## 当前批次进展

### 已完成内容
- 已调整 `internal/order/service.go`：
  - 注入购物车查询依赖
  - `CreateOrderFromCheckedCartItems(userID)` 改为主动读取已勾选购物车项
- 已调整 `internal/order/handler.go`：
  - `CreateFromCartRequest` 只保留 `user_id`
  - `/orders/from-cart` 不再接收购物车项数组
- 已调整 `cmd/order-service/main.go`：
  - 在运行时装配 `cart.NewMemoryRepository -> cart.NewService -> order.NewService`

### 当前行为范围
- `/orders/from-cart` 现在只接收 `user_id`
- 订单服务会主动读取该用户已勾选购物车项并完成下单
- 空购物车或无已勾选项会返回错误

### 当前限制
- 当前读取的是内存购物车仓储，不是真实持久化购物车
- 当前仍未接入库存校验与扣减
- 当前仍未接入 MySQL 持久化

### 当前验证记录
- 已执行：`go test ./internal/order && go test ./cmd/order-service`
- 结果：通过

### 全量验证与远端同步
- 已执行：`go test ./...`
- 结果：通过
- 已创建提交：`abb1474` `feat: orchestrate cart to order workflow`
- 已推送到远端 `origin/master`
- 已验证：
  - `git status` 工作区干净
  - `git branch -vv` 显示本地 `master` 跟踪 `origin/master`
  - `git ls-remote --heads origin` 显示远端 `master` 指向 `abb1474`
