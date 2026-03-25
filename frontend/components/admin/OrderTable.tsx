type OrderTableProps = {
  orders: { id?: number; orderNo?: string; totalAmount: number; status: string; items?: { productName: string; quantity: number }[] }[];
};

export function OrderTable({ orders }: OrderTableProps) {
  return (
    <section className="metric-card" style={{ display: "grid", gap: 18 }}>
      <div className="inline-stack" style={{ justifyContent: "space-between" }}>
        <div>
          <p className="eyebrow">Orders</p>
          <h2 className="section-title" style={{ margin: 0 }}>订单管理</h2>
        </div>
        <p className="shop-meta">展示订单号、金额、状态和明细摘要</p>
      </div>
      <table className="panel-table">
        <thead><tr><th>订单号</th><th>金额</th><th>状态</th><th>订单项数</th><th>明细摘要</th></tr></thead>
        <tbody>
          {orders.map((item, index) => (
            <tr key={item.orderNo ?? item.id ?? index}>
              <td>{item.orderNo ?? `未命名订单-${index + 1}`}</td>
              <td>¥{(item.totalAmount / 100).toFixed(2)}</td>
              <td>{item.status}</td>
              <td>{item.items?.length ?? 0}</td>
              <td>{item.items?.map((detail) => `${detail.productName} x${detail.quantity}`).join(' / ') || '无明细'}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </section>
  );
}
