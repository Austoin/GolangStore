# 阶段 04：商品购物车订单主链路

## 时间
- 2026-03-24

## 阶段目标
- 建立商品、购物车、订单三条主链路的最小业务骨架。
- 先从领域模型开始，再逐步补充 repo、service、handler。
- 保持每一步都可测试、可追踪、可提交。

## 当前计划顺序
1. 落商品领域模型
2. 落购物车领域模型
3. 落订单领域模型
4. 创建商品查询最小 repo/service/handler
5. 创建购物车 CRUD 最小 repo/service/handler
6. 创建下单与订单查询最小 repo/service/handler

## 当前约束
- 先不接入真实 HTTP 路由注册。
- 先不接入数据库真实读写。
- 先把领域结构和测试基线搭稳。

## 当前状态
- 阶段刚开始，尚未创建业务领域文件。

## 批次 04A：商品领域模型

### 本批次目标
- 先落商品领域模型。
- 只包含最小字段与最小业务判断行为。
- 同步更新 README 与阶段文档。

### 本批次已完成
- 已创建：`internal/product/model.go`
- 已创建：`internal/product/model_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- 定义 `Product` 结构体。
- 当前包含字段：
  - `ID`
  - `Name`
  - `Description`
  - `Price`
  - `Status`
  - `Stock`
- 提供行为：
  - `IsOnSale()`
  - `HasStock()`

### 当前验证结果
- 已执行：`go test ./internal/product`
- 结果：通过

### 下一步
- 进入购物车领域模型。

## 批次 04B：购物车领域模型

### 本批次目标
- 落购物车领域模型。
- 保持字段最小化，只表达当前购物车项的核心状态。
- 同步更新 README 与阶段文档。

### 本批次已完成
- 已创建：`internal/cart/model.go`
- 已创建：`internal/cart/model_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- 定义 `Item` 结构体。
- 当前包含字段：
  - `UserID`
  - `ProductID`
  - `Quantity`
  - `Checked`
- 提供行为：
  - `IsValidQuantity()`
  - `IsChecked()`

### 当前验证目标
- 验证数量合法性判断。
- 验证勾选状态判断。

### 当前验证结果
- 已执行：`go test ./internal/cart`
- 结果：通过

### 下一步
- 运行 `internal/cart` 测试。

## 批次 04C：订单领域模型

### 本批次目标
- 落订单领域模型。
- 先定义最小订单状态、订单项和订单聚合结构。
- 同步更新 README 与阶段文档。

### 本批次已完成
- 已创建：`internal/order/model.go`
- 已创建：`internal/order/model_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- 定义订单状态常量：
  - `StatusPending`
  - `StatusPaid`
  - `StatusClosed`
- 定义 `Item` 与 `Order` 结构体。
- 提供行为：
  - `IsPending()`
  - `IsPaid()`
  - `HasItems()`

### 当前验证目标
- 验证订单状态判断。
- 验证订单项存在性判断。

### 当前验证结果
- 已执行：`go test ./internal/order`
- 结果：通过

### 下一步
- 运行 `internal/order` 测试。

## 批次 04D：商品查询最小链路

### 本批次目标
- 从商品领域模型进入第一条最小业务链路。
- 先实现 repo、service、handler 的最小串联。
- 当前只使用内存仓储，不接真实数据库。

### 本批次已完成
- 已创建：
  - `internal/product/repo.go`
  - `internal/product/service.go`
  - `internal/product/handler.go`
  - `internal/product/repo_test.go`
  - `internal/product/service_test.go`
  - `internal/product/handler_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- `Repository` 定义商品按 ID 查询能力。
- `MemoryRepository` 提供当前阶段可测试的内存实现。
- `Service` 提供 `GetProduct(id)`。
- `Handler` 提供 `GetByID(ctx)`。
- `Handler` 当前处理：
  - 非法 ID -> `400`
  - 商品不存在 -> `404`
  - 查询成功 -> `200`

### 当前验证目标
- 验证仓储层的查询行为。
- 验证服务层的串联行为。
- 验证处理器层的 HTTP 状态码行为。

### 当前验证结果
- 已执行：`go test ./internal/product`
- 结果：通过

### 下一步
- 运行 `internal/product` 全量测试。

## 批次 04E：购物车查询最小链路

### 本批次目标
- 从购物车领域模型进入最小业务链路。
- 先实现 repo、service、handler 的最小查询串联。
- 当前只使用内存仓储，不接真实数据库。

### 本批次已完成
- 已创建：
  - `internal/cart/repo.go`
  - `internal/cart/service.go`
  - `internal/cart/handler.go`
  - `internal/cart/repo_test.go`
  - `internal/cart/service_test.go`
  - `internal/cart/handler_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- `Repository` 定义按用户 ID 查询购物车项。
- `MemoryRepository` 提供当前阶段可测试的内存实现。
- `Service` 提供 `ListItems(userID)`。
- `Handler` 提供 `ListByUserID(ctx)`。
- `Handler` 当前处理：
  - 非法用户 ID -> `400`
  - 查询成功 -> `200`

### 当前验证目标
- 验证仓储层按用户过滤行为。
- 验证服务层的串联行为。
- 验证处理器层的 HTTP 状态码行为。

### 当前验证结果
- 已执行：`go test ./internal/cart`
- 结果：通过

### 下一步
- 运行 `internal/cart` 全量测试。

## 批次 04F：订单查询最小链路

### 本批次目标
- 从订单领域模型进入最小业务链路。
- 先实现 repo、service、handler 的最小查询串联。
- 当前只使用内存仓储，不接真实数据库，不进入下单创建。

### 本批次已完成
- 已创建：
  - `internal/order/repo.go`
  - `internal/order/service.go`
  - `internal/order/handler.go`
  - `internal/order/repo_test.go`
  - `internal/order/service_test.go`
  - `internal/order/handler_test.go`
- 已同步更新：`README.md`

### 当前实现内容
- `Repository` 定义按订单号查询订单。
- `MemoryRepository` 提供当前阶段可测试的内存实现。
- `Service` 提供 `GetOrder(orderNo)`。
- `Handler` 提供 `GetByOrderNo(ctx)`。
- `Handler` 当前处理：
  - 订单不存在 -> `404`
  - 查询成功 -> `200`

### 当前验证目标
- 验证仓储层按订单号查询行为。
- 验证服务层的串联行为。
- 验证处理器层的 HTTP 状态码行为。

### 当前验证结果
- 已执行：`go test ./internal/order`
- 结果：通过

### 下一步
- 运行 `internal/order` 全量测试。

## 批次 04G：阶段 04 汇总与全量验证

### 本批次目标
- 汇总当前阶段已完成的领域模型与最小查询链路。
- 进行全量测试验证，确认当前阶段可安全提交。

### 当前阶段已完成范围
- 商品领域模型与最小行为测试。
- 购物车领域模型与最小行为测试。
- 订单领域模型与最小行为测试。
- 商品查询最小 repo/service/handler 链路与测试。
- 购物车查询最小 repo/service/handler 链路与测试。
- 订单查询最小 repo/service/handler 链路与测试。

### 当前验证结果
- 已执行：`go test ./...`
- 结果：
  - `internal/product` 通过
  - `internal/cart` 通过
  - `internal/order` 通过
  - `pkg/config` 通过
  - `pkg/mysql` 通过
  - `pkg/redis` 通过
  - 6 个服务入口包可正常编译

### 当前阶段结论
- 阶段 04 当前成果已具备提交条件。
- 下一步进入提交并推送阶段。
