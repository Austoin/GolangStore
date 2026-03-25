type OrderTableProps = {
  orders: { orderNo: string; totalAmount: number; status: string; items?: { productName: string; quantity: number }[] }[];
};

export function OrderTable({ orders }: OrderTableProps) {
  return (
    <section className="metric-card">
      <h2>订单管理</h2>
      <ul>
        {orders.map((item) => (
          <li key={item.orderNo}>{item.orderNo} / ¥{(item.totalAmount / 100).toFixed(2)} / {item.status} / 明细 {item.items?.length ?? 0}</li>
        ))}
      </ul>
    </section>
  );
}
