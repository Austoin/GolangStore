type OrderTableProps = {
  orders: { orderNo: string; totalAmount: number; status: string; items?: { productName: string; quantity: number }[] }[];
};

export function OrderTable({ orders }: OrderTableProps) {
  return (
    <section className="metric-card">
      <h2 className="section-title">订单管理</h2>
      <table className="panel-table">
        <thead><tr><th>订单号</th><th>金额</th><th>状态</th><th>订单项数</th></tr></thead>
        <tbody>
          {orders.map((item) => (
            <tr key={item.orderNo}><td>{item.orderNo}</td><td>¥{(item.totalAmount / 100).toFixed(2)}</td><td>{item.status}</td><td>{item.items?.length ?? 0}</td></tr>
          ))}
        </tbody>
      </table>
    </section>
  );
}
