# 阶段 00：项目初始化与范围收敛

## 时间
- 2026-03-24

## 阶段目标
- 从 0 开始规划一个基于 Go 微服务架构的智能商城后端。
- 一期先完成最小可交付 MVP，并为后续扩展到高并发秒杀、MCP/LLM 网关、容器化部署保留结构。

## 当时仓库状态
- 仓库初始仅存在空的 `README.md`。
- 尚未创建代码目录、部署清单、数据库脚本、服务骨架与测试结构。

## 已确认范围
- 版本目标：MVP。
- 鉴权策略：无鉴权。
- 架构粒度：4 个核心服务。
- 运行基线：先使用 Docker Compose。

## 已确认架构决策
- 采用单仓多服务组织方式，先保证清晰边界，再逐步补齐基础设施。
- 对人类调用与 AI 调用分离：
  - `gateway` 负责商城 HTTP API 入口。
  - `mcp-gateway` 负责 MCP 协议适配与 LLM Skills 暴露。
- 一期核心领域服务：
  - `user-service`
  - `product-service`
  - `cart-service`
  - `order-service`
- 基础设施使用 MySQL 与 Redis：
  - MySQL 负责业务持久化。
  - Redis 负责热点缓存、库存读优化、秒杀原子扣减。

## 服务职责定义
- `gateway`：统一 API 入口、参数校验、路由转发、响应包装。
- `mcp-gateway`：向 LLM 暴露标准化工具能力，一期仅开放库存查询与订单状态查询。
- `user-service`：保留最小用户实体与地址占位，不做登录鉴权。
- `product-service`：商品列表、商品详情、库存查询、缓存管理。
- `cart-service`：购物车增删改查与购物项整理。
- `order-service`：订单创建、订单查询、订单状态流转、秒杀库存编排。

## 主链路定义
- 普通下单：`gateway -> cart-service -> order-service -> product-service/Redis -> MySQL`
- 秒杀下单：`gateway -> order-service -> Redis Lua -> MySQL`
- AI 调用：`LLM Agent -> mcp-gateway -> product-service/order-service`

## 约束
- 不做无鉴权之外的权限模型。
- 不在一期引入完整消息队列系统。
- 不在一期直接上 Kubernetes 集群交付，先完成 Docker Compose 联调闭环。

## 风险
- 如果一期边界失控，容易过早引入支付、消息队列、监控、后台权限等高成本模块。
- 如果服务边界定义不清，后续 MCP 适配与秒杀链路会耦合在业务服务中。

## 阶段结论
- 已完成 MVP 范围收敛。
- 已明确系统边界、服务职责与主链路。
- 可以进入领域建模与阶段拆分阶段。
