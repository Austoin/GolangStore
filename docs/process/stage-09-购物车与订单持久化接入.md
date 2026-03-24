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

### 本批次提交与远端同步
- 已创建提交：`7b6a805` `feat: add cart mysql repository`
- 已推送到远端 `origin/master`
- 已验证：
  - `git status` 工作区干净
  - `git branch -vv` 显示本地 `master` 跟踪 `origin/master`
  - `git ls-remote --heads origin` 显示远端 `master` 指向 `7b6a805`

## 批次 09C：order MySQL 仓储

### 已完成内容
- 已创建：`internal/order/mysql_repo.go`
- 已创建：`internal/order/mysql_repo_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- `MySQLRepository` 使用 Gorm 读写：
  - `orders`
  - `order_items`
- 当前已支持：
  - `Create(order)`
  - `GetByOrderNo(orderNo)`
- 已完成数据库行到领域模型的映射：
  - 订单主表 -> `Order`
  - 订单项表 -> `[]Item`

### 当前验证结果
- 已执行：`go test ./internal/order`
- 结果：通过

### 当前限制
- 当前写入未包事务，后续接入真实生产链路时需要补事务控制。
- 当前测试基于 sqlite 内存库验证仓储行为，不直接依赖真实 MySQL 实例。

### 本批次提交与远端同步
- 已创建提交：`7c70ee7` `feat: add order mysql repository`
- 已推送到远端 `origin/master`
- 已验证：
  - `git status` 工作区干净
  - `git branch -vv` 显示本地 `master` 跟踪 `origin/master`
  - `git ls-remote --heads origin` 显示远端 `master` 指向 `7c70ee7`

## 批次 09D：切换 order-service 到 MySQL 装配

### 已完成内容
- 已更新：`cmd/order-service/main.go`
- 已更新：`cmd/order-service/router_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- `cmd/order-service` 启动时改为：
  - 读取 `config.Load()`
  - 使用 `pkg/mysql.Open()` 打开数据库连接
  - 装配 `cart.NewMySQLRepository(db)`
  - 装配 `order.NewMySQLRepository(db)`
- 已将运行时装配拆分为：
  - `buildRuntimeDependencies(conf)`
  - `buildHandler(db)`
- 测试通过 `buildHandler(db)` 避免直接依赖真实 MySQL 连接。

### 当前验证结果
- 已执行：`go test ./cmd/order-service`
- 结果：通过

### 当前限制
- 当前服务启动仍直接 `panic(err)`，后续可再补更友好的启动错误处理。
- 当前未自动执行建表/迁移，默认依赖数据库已有初始化脚本。

### 本批次提交与远端同步
- 已创建提交：`1a73534` `feat: wire order service to mysql`
- 已推送到远端 `origin/master`
- 已验证：
  - `git status` 工作区干净
  - `git branch -vv` 显示本地 `master` 跟踪 `origin/master`
  - `git ls-remote --heads origin` 显示远端 `master` 指向 `1a73534`

## 阶段 10：一键启动验证脚本

### 已完成内容
- 已创建：`.gitignore`
- 已创建：`scripts/start.sh`
- 已更新：`docker-compose.yml`
- 已同步更新：`README.md`

### 当前实现内容
- `scripts/start.sh` 会执行：
  - 启动 `mysql`、`redis`
  - 等待中间件就绪
  - 后台启动 `order-service`
  - 检查 `http://127.0.0.1:8082/health`
- 运行时日志写入：`.runtime/order-service.log`
- 运行时 PID 写入：`.runtime/order-service.pid`

### 当前验证结果
- 已执行：`bash -n scripts/start.sh`
- 已执行：`bash scripts/start.sh`
- 当前环境结果：失败，错误为 `docker daemon is not running`

### 当前结论
- 脚本语法正确。
- 脚本在 Docker daemon 不可用时会给出明确错误，而不是静默失败。
- 当前无法完成真实启动，是环境中的 Docker daemon 未启动，不是脚本解析错误。

## 阶段 10：启动脚本诊断增强

### 已完成内容
- 已更新：`scripts/start.sh`
- 已同步更新：`README.md`

### 当前增强点
- 启动时输出脚本版本号。
- 输出当前 `docker`、`curl`、`go` 的实际路径。
- 输出 Docker daemon 检查结果。

### 当前目的
- 避免误跑旧脚本却无法辨认。
- 让启动失败时能快速区分：
  - 脚本版本问题
  - `go` 路径问题
  - Docker daemon 未启动

## 阶段 10：Go 路径兼容修复

### 当前问题
- 在不同 `bash` 环境中，`go` 的可执行路径解析不同。
- 导致脚本在某些环境里误报 `ERROR: go is required`。

### 当前修复
- 已为 `scripts/start.sh` 增加 `resolve_go_bin()`。
- 当前优先级为：
  1. 当前 shell 中的 `go`
  2. `/mnt/c/Program Files/Go/bin/go.exe`
  3. `/c/Program Files/Go/bin/go.exe`

### 当前验证结果
- 已执行：`bash -n scripts/start.sh && bash scripts/start.sh`
- 结果：
  - 不再报 `go is required`
  - 已成功进入 Docker Compose 启动流程
  - 当前命令超时终止于镜像拉取阶段，不是 Go 路径问题

## 阶段 10：持久化模式一致性修复

### 当前问题
- `internal/cart` 的 MySQL 仓储与购物车到下单编排逻辑依赖 `cart_items.product_name`、`cart_items.price`。
- 但初始化 SQL 中的 `cart_items` 表尚未包含这两个字段。

### 当前修复
- 已更新 `deploy/mysql/001_init_schema.sql`。
- 已为 `cart_items` 增加：
  - `product_name`
  - `price`

### 当前影响
- 初始化数据库后，购物车读取模型与订单编排模型将不再发生字段缺失。
- 该修复避免了运行时读取购物车时的结构不一致问题。

### 本批次提交与远端同步
- 已创建提交：`227d2b4` `fix: align cart schema with order workflow`
- 已推送到远端 `origin/master`
- 已验证：
  - `git status` 工作区干净
  - `git branch -vv` 显示本地 `master` 跟踪 `origin/master`
  - `git ls-remote --heads origin` 显示远端 `master` 指向 `227d2b4`
