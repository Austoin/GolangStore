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
- 已提供 `scripts/start.sh` 一键启动脚本，用于启动 MySQL、Redis 和 order-service
- `cart_items` 初始化表结构已补齐 `product_name`、`price`，与当前仓储/编排模型保持一致
- `scripts/start.sh` 已增加版本号、go 路径、docker daemon 状态等诊断输出，便于排查启动失败原因
- `scripts/start.sh` 已兼容 Git Bash / WSL 对 Go 路径的差异，不再因 shell 环境差异误报 `go is required`
- 已完成一次真实业务验证：插入购物车数据 -> 调用 `/orders/from-cart` -> `orders` 与 `order_items` 成功落库
- 已创建端到端业务完成总计划，目标是补全核心业务逻辑并支持一键启动、一键演示效果
- `internal/product/mysql_repo.go` 与 `cmd/product-service` 已落地，商品查询已具备 MySQL 读取能力
- `internal/cart` 与 `cmd/cart-service` 已落地写接口，购物车新增已具备 MySQL 写入能力
- `internal/order` 已接入库存校验与扣减逻辑，下单前会检查库存并在成功后扣减库存
- `scripts/demo.sh` 已创建，但当前验证表明需要先让 `start.sh` 同时启动 `cart-service`
- 当前已完成端到端闭环验证：`bash scripts/start.sh && bash scripts/demo.sh` 可成功完成商品、购物车、下单、订单落库、库存扣减验证
- `frontend/` Next.js 项目骨架已创建，前后台可视化面板开始进入实现阶段
- `frontend/components/shared` 已建立共享设计系统基础组件
- `frontend/app/shop` 四个前台页面骨架已完成
- `frontend/app/admin` 四个后台页面骨架已完成
- `frontend/lib/adapters` 与 `frontend/lib/mock` 已建立，前台页面已切到 adapter 数据层
- `/shop` 商品列表页已切到真实后端接口 `GET /products`

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
- `docs/process/stage-12-端到端业务闭环收束.md`
- `docs/plans/2026-03-25-end-to-end-business-completion-plan.md`

## 当前运行方式

### 后端演示

```bash
bash scripts/start.sh
bash scripts/demo.sh
```

### 一键启动前后端

```bash
bash scripts/run-all.sh
```

该脚本当前会：

- 启动 MySQL、Redis
- 启动 `product-service`、`cart-service`、`order-service`
- 启动前端 Next.js
- 直接打印前端访问地址
- 等待 `/`、`/shop`、`/admin` 全部就绪

### 一键停止前后端

```bash
bash scripts/stop-all.sh
```

### 前端开发

```bash
cd frontend
npm install
npm run dev
```

前端当前状态：

- 一个 Next.js 项目
- 两套入口：`/shop` 与 `/admin`
- 当前通过 mock adapter 驱动页面
- 后续可替换为真实 HTTP adapter
- 已验证页面入口：
  - `http://127.0.0.1:3000/`
  - `http://127.0.0.1:3000/shop`
  - `http://127.0.0.1:3000/admin`
- 已验证一键启动脚本：`bash scripts/run-all.sh`
- `run-all.sh` 已增强启动期提示，可在前端预热阶段直接看到地址与等待进度

- `/shop/products/[id]` 已切到真实商品详情接口 `GET /products/:id`
- `/shop/cart` 已切到真实购物车接口，并可触发真实下单
- `/shop/orders` 已切到真实订单列表接口 `GET /orders`
- `/admin`、`/admin/products`、`/admin/inventory`、`/admin/orders` 已切到真实后端数据驱动
- 当前前台与后台都已能通过真实后端接口展示核心业务数据
- `product-service` 的商品 JSON 输出已统一为前端可直接消费的小写字段
- 前后台都已重做为统一设计语言的可操作面板
- `/shop` 与 `/admin` 当前都已可基于真实后端数据体验核心功能
- 前台商品卡已支持“查看商品”和“加入购物车”动作
- 后台订单列表已修复稳定 key 展示逻辑
- 购物车页在库存不足等业务失败时会显示页面内错误提示，而不是直接抛 Application error
- 商品卡已支持数量输入，加入购物车时会真实提交所选数量
- `/shop/cart` 已可展示每个商品的数量、小计和总计
