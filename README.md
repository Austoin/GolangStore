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

详细过程见：

- `docs/plans/2026-03-24-mvp-implementation-plan.md`
- `docs/process/stage-00-项目初始化与范围收敛.md`
- `docs/process/stage-01-领域建模与实施阶段拆分.md`
- `docs/process/stage-02-项目骨架创建.md`
- `docs/process/stage-03-公共基础设施与数据库初始化.md`
