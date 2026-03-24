# 阶段 05：下单创建与状态流转

## 时间
- 2026-03-24

## 阶段目标
- 从现有查询链路推进到最小下单闭环。
- 实现购物车项转订单项、总金额计算、订单状态初始化。
- 先完成内存级下单流程，再考虑接数据库持久化。

## 当前计划顺序
1. 定义下单输入模型
2. 定义金额计算与订单项转换规则
3. 实现内存级订单创建服务
4. 实现最小创建订单 handler
5. 补测试并做全量验证

## 当前约束
- 当前阶段不接入真实 MySQL 持久化。
- 当前阶段不接入 Redis 扣减。
- 当前阶段不做支付状态流转，只初始化为 `pending`。

## 当前状态
- 阶段刚开始，尚未创建下单创建相关代码。

## 当前批次进展

### 已完成内容
- 已扩展 `internal/order/model.go`：
  - `CreateRequest`
  - `CreateItemRequest`
  - `CalculateTotalAmount()`
- 已创建 `internal/order/create.go`，提供购物车项转订单项逻辑。
- 已扩展 `internal/order/service.go`：
  - `ConvertCartItems(...)`
  - `CreateOrder(...)`
- 已扩展 `internal/order/repo.go`：
  - `Create(order)`
- 已扩展 `internal/order/handler.go`：
  - `Create(ctx)`

### 当前行为范围
- 可接收最小下单请求。
- 可将下单项转换为订单项。
- 可计算订单总金额。
- 可生成内存级订单并初始化状态为 `pending`。
- 可通过 HTTP POST 创建订单。

### 当前限制
- 仍未接入真实 MySQL 持久化。
- 当前订单号生成规则为最小可运行方案，不是最终生产规则。
- 当前未接入 Redis 库存扣减。
- 当前未实现支付后的状态流转。

### 当前验证记录
- 已多次执行：`go test ./internal/order`
- 已验证通过：
  - 下单输入模型
  - 金额计算
  - 购物车到订单项转换
  - 内存级订单创建
  - 创建订单 handler
