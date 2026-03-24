# 阶段 07：order-service 路由接线

## 时间
- 2026-03-24

## 阶段目标
- 将当前 `internal/order` 的查询与下单能力接入 `cmd/order-service` 运行入口。
- 让订单服务从“内部包可测试”推进到“服务进程可访问”。

## 当前计划顺序
1. 在 `cmd/order-service` 装配内存仓储与 service
2. 挂载订单查询路由
3. 挂载创建订单路由
4. 挂载购物车到下单编排路由
5. 验证入口行为并更新文档

## 当前约束
- 当前阶段仍使用内存仓储。
- 当前阶段不接入 MySQL、Redis。
- 当前阶段只做订单服务入口，不动其他服务入口。

## 当前状态
- 阶段刚开始，`cmd/order-service` 仍只有 `/health`。

## 当前批次进展

### 已完成内容
- 已创建：`cmd/order-service/router.go`
- 已创建：`cmd/order-service/router_test.go`
- 已更新：`cmd/order-service/main.go`

### 当前已接线路由
- `GET /health`
- `GET /orders/:orderNo`
- `POST /orders`
- `POST /orders/from-cart`

### 当前实现说明
- `cmd/order-service` 当前通过 `order.NewMemoryRepository(nil)` 装配内存仓储。
- `cmd/order-service` 只负责装配 `repo -> service -> handler -> router`。
- 业务逻辑仍留在 `internal/order`，未污染入口层。

### 当前验证记录
- 已执行：`go test ./cmd/order-service`
- 结果：通过
