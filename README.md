## GolangStore

基于 Go 微服务架构的智能商城后端项目。

当前阶段仅完成项目规划与工程骨架初始化，后续将按以下顺序推进：

1. 仓库骨架与公共基础设施
2. 商品、购物车、订单主链路
3. Redis 库存与秒杀链路
4. MCP 网关与 LLM Skills
5. Docker Compose 联调与交付整理

当前已完成：

- Git 仓库初始化并推送到 GitHub
- 6 个服务入口骨架与健康检查接口
- `pkg/config` 基础配置加载能力与单元测试
- `pkg/mysql` DSN 构建与 Gorm 初始化入口
- `pkg/redis` 地址构建、客户端初始化与最小测试
- `deploy/mysql/001_init_schema.sql` 首批基础表初始化脚本
- `internal/product` 商品领域模型与最小行为测试
- `internal/cart` 购物车领域模型与最小行为测试
- `internal/order` 订单领域模型与最小行为测试
- `internal/product` 商品查询最小 repo/service/handler 链路与测试
- `internal/cart` 购物车查询最小 repo/service/handler 链路与测试
- `internal/order` 订单查询最小 repo/service/handler 链路与测试
- 当前全量验证命令 `go test ./...` 已通过
- `阶段 05` 计划文档已创建，准备进入下单创建与状态流转实现
- `internal/order` 已具备最小下单创建流程：输入模型、金额计算、购物车转订单项、创建订单 handler
- `阶段 06` 计划文档已创建，准备进入购物车到下单编排实现
- `internal/cart` 已具备已勾选购物车项提取能力，`internal/order` 已具备购物车到下单编排入口
- `阶段 07` 计划文档已创建，准备接通 `order-service` 运行入口路由
- `cmd/order-service` 已接通订单查询、创建订单、购物车到下单编排路由
- `阶段 08` 计划文档已创建，准备改造成“订单服务主动读取购物车”的真实编排
- `/orders/from-cart` 已收缩为只接收 `user_id`，并由 `order-service` 主动读取已勾选购物车项完成下单
- `阶段 09` 计划文档已创建，准备把购物车与订单从内存实现切换到 MySQL 持久化
- `internal/cart/mysql_repo.go` 已落地，购物车读取已具备 MySQL 仓储实现
- `internal/order/mysql_repo.go` 已落地，订单创建与订单查询已具备 MySQL 仓储实现
- `cmd/order-service` 运行时装配已切换到 MySQL 仓储实现
- `cart_items` 初始化表结构已补齐 `product_name`、`price`，与当前仓储/编排模型保持一致

详细过程见：

- `docs/plans/2026-03-24-mvp-implementation-plan.md`
- `docs/process/stage-00-项目初始化与范围收敛.md`
- `docs/process/stage-01-领域建模与实施阶段拆分.md`
- `docs/process/stage-02-项目骨架创建.md`
- `docs/process/stage-03-公共基础设施与数据库初始化.md`
- `docs/process/stage-04-商品购物车订单主链路.md`
- `docs/process/stage-05-下单创建与状态流转.md`
- `docs/plans/2026-03-24-order-creation-stage-plan.md`
- `docs/process/stage-06-购物车到下单编排.md`
- `docs/plans/2026-03-24-cart-to-order-orchestration-plan.md`
- `docs/process/stage-07-order-service路由接线.md`
- `docs/plans/2026-03-24-order-service-routing-plan.md`
- `docs/process/stage-08-真实购物车到下单编排.md`
- `docs/plans/2026-03-24-real-cart-to-order-plan.md`
- `docs/process/stage-09-购物车与订单持久化接入.md`
- `docs/plans/2026-03-24-cart-order-persistence-plan.md`
