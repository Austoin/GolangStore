# 阶段 03：公共基础设施与数据库初始化

## 时间
- 2026-03-24

## 阶段目标
- 建立公共基础设施包的最小骨架。
- 创建数据库初始化目录与首批 SQL 文件。
- 保持该阶段只处理基础设施，不进入业务 handler、service、repo 实现。

## 当前计划顺序
1. 创建 `pkg/config`
2. 创建 `pkg/mysql`
3. 创建 `pkg/redis`
4. 创建 `deploy/mysql` 初始化目录
5. 创建首批数据库建表脚本

## 当前约束
- 不提前创建业务接口实现。
- 不提前创建复杂公共抽象。
- 先保证配置加载、连接初始化和建表脚本位置稳定。

## 当前状态
- 阶段刚开始，尚未创建公共基础设施代码。
- 项目已具备可编译的服务入口骨架，可承接下一步基础设施接入。

## 本阶段新增进展：Git 远端接入

### 已完成内容
- 已在本地初始化 Git 仓库。
- 已添加远端仓库：`https://github.com/austoin/GolangStore.git`
- 已创建首个初始化提交：`03c95e3`
- 已执行首次推送，并验证本地 `master` 已跟踪远端 `origin/master`

### 验证结果
- `git status` 输出工作区干净。
- `git branch -vv` 显示当前分支跟踪 `origin/master`。
- `git ls-remote --heads origin` 可看到远端 `master` 分支指向提交 `03c95e3`

### 当前注意事项
- 远端返回仓库地址已迁移到大小写规范后的地址：`https://github.com/Austoin/GolangStore.git`
- 当前旧地址仍可成功推送，后续可在合适阶段统一更新远端地址，当前不影响继续开发。

## 批次 9A：配置包初始化

### 本批次目标
- 创建最小配置加载包。
- 通过单元测试固定默认值与环境变量覆盖行为。
- 同步更新 README 和阶段文档，保证工程状态可追踪。

### 本批次已完成
- 已创建：`pkg/config/config.go`
- 已创建：`pkg/config/config_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- 定义了 `App`、`MySQL`、`Redis`、`Config` 结构体。
- 提供 `Load()` 统一加载环境变量。
- 提供 `getEnv()` 作为最小默认值回退逻辑。

### 当前验证目标
- 验证默认值是否正确。
- 验证环境变量覆盖是否生效。

### 当前验证结果
- 已执行：`go test ./pkg/config`
- 结果：通过

### 下一步
- 运行 `pkg/config` 单元测试。
- 测试通过后进入 `pkg/mysql` 初始化。

## 批次 9B：MySQL 包初始化

### 本批次目标
- 创建最小 MySQL 初始化包。
- 先固定 DSN 构建逻辑，再提供 Gorm 打开入口。
- 不在当前批次引入真实数据库连接测试。

### 本批次已完成
- 已创建：`pkg/mysql/mysql.go`
- 已创建：`pkg/mysql/mysql_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- 提供 `BuildDSN(conf config.MySQL) string`。
- 提供 `Open(conf config.MySQL) (*gorm.DB, error)`。
- DSN 中固定 `charset=utf8mb4&parseTime=True&loc=Local`。

### 当前验证目标
- 验证 DSN 构建结果是否符合预期。

### 当前验证结果
- 创建 `pkg/mysql` 后因缺少 Gorm 依赖触发模块错误。
- 已执行：`go mod tidy`
- 已执行：`go test ./pkg/mysql`
- 结果：通过

### 下一步
- 运行 `pkg/mysql` 单元测试。
- 测试通过后进入 `pkg/redis` 初始化。

## 批次 9C：Redis 包初始化

### 本批次目标
- 创建最小 Redis 初始化包。
- 固定地址拼接与客户端构建逻辑。
- 先测试配置映射，不引入真实 Redis 连接测试。

### 本批次已完成
- 已创建：`pkg/redis/redis.go`
- 已创建：`pkg/redis/redis_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- 提供 `BuildAddr(conf config.Redis) string`。
- 提供 `NewClient(conf config.Redis) *redis.Client`。
- 提供 `Ping(ctx, client)` 作为最小连通性检查入口。

### 当前验证目标
- 验证 Redis 地址构建逻辑。
- 验证客户端实例是否正确接收配置。

### 当前验证结果
- 创建 `pkg/redis` 后因缺少 Redis 客户端依赖触发模块错误。
- 已执行：`go mod tidy`
- 已执行：`go test ./pkg/redis`
- 结果：通过

### 下一步
- 运行 `pkg/redis` 单元测试。
- 测试通过后进入 `deploy/mysql` 初始化脚本目录创建。

## 批次 9D：数据库初始化脚本

### 本批次目标
- 创建 MySQL 初始化目录。
- 按已确认的领域模型落首批基础表脚本。
- 同步更新 README 与阶段文档。

### 本批次已完成
- 已创建目录：`deploy/mysql/`
- 已创建脚本：`deploy/mysql/001_init_schema.sql`
- 已同步更新：`README.md`

### 当前脚本内容
- 已建表：
  - `users`
  - `products`
  - `product_stocks`
  - `cart_items`
  - `orders`
  - `order_items`
- 已按当前设计加入最小唯一索引与查询索引。

### 当前设计说明
- 当前阶段未增加外键约束，避免在骨架阶段过早绑定迁移顺序与回滚复杂度。
- 金额字段使用 `BIGINT UNSIGNED`，按分存储，避免浮点误差。
- `product_stocks` 独立成表，便于后续库存逻辑单独优化。

### 当前验证目标
- 验证 SQL 文件结构完整。
- 为下一步接入 Compose 初始化或手工导入提供稳定脚本入口。

### 当前验证结果
- 已读取并检查 `deploy/mysql/001_init_schema.sql`，脚本包含 6 张核心表定义，结构完整。
- 已执行：`go test ./...`
- 结果：
  - `pkg/config` 通过
  - `pkg/mysql` 通过
  - `pkg/redis` 通过
  - 6 个服务入口包可正常编译

### 下一步
- 汇总阶段 03 当前成果。
- 执行一次全量测试并推送到远端。
