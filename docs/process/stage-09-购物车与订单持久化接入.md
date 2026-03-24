# 阶段 09：购物车与订单持久化接入

## 时间
- 2026-03-24

## 阶段目标
- 将购物车与订单从内存实现逐步切换到 MySQL 持久化实现。
- 先保证读取购物车、创建订单、查询订单三条核心链路具备真实数据落盘能力。

## 当前计划顺序
1. 为 `cart` 建立 MySQL 仓储
2. 为 `order` 建立 MySQL 仓储
3. 在 `order-service` 中切换运行时装配
4. 验证下单与订单查询主链路
5. 更新文档并逐步推送云端

## 当前约束
- 当前阶段不接入 Redis 库存扣减。
- 当前阶段不做分布式事务。
- 当前阶段优先保证最小真实持久化可用。

## 当前状态
- 阶段刚开始，购物车与订单仍使用内存仓储。

## 本阶段首个提交

### 已完成内容
- 已创建阶段文档：`docs/process/stage-09-购物车与订单持久化接入.md`
- 已创建计划文档：`docs/plans/2026-03-24-cart-order-persistence-plan.md`
- 已同步更新：`README.md`

### 当前验证结果
- 已创建提交：`ce44ba0` `docs: add persistence stage plan`
- 已推送到远端 `origin/master`
- 已验证：
  - `git status` 工作区干净
  - `git branch -vv` 显示本地 `master` 跟踪 `origin/master`
  - `git ls-remote --heads origin` 显示远端 `master` 指向 `ce44ba0`

## 批次 09B：cart MySQL 仓储

### 已完成内容
- 已创建：`internal/cart/mysql_repo.go`
- 已创建：`internal/cart/mysql_repo_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- `MySQLRepository` 使用 Gorm 读取 `cart_items` 表。
- 当前已支持：`ListByUserID(userID)`。
- 已完成数据库行到 `cart.Item` 的字段映射：
  - `user_id`
  - `product_id`
  - `product_name`
  - `price`
  - `quantity`
  - `checked`

### 当前验证结果
- 已执行：`go test ./internal/cart`
- 结果：通过

### 当前限制
- 当前只实现读取，不含写入能力。
- 当前测试基于 sqlite 内存库验证 SQL 映射，不直接依赖真实 MySQL 实例。
