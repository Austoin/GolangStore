import { AppHeader } from "../../../components/shared/AppHeader";
import { OrderCard } from "../../../components/shop/OrderCard";
import { listOrdersHttp } from "../../../lib/adapters/httpAdapter";

export default async function OrdersPage() {
  const orders = await listOrdersHttp();

  return (
    <main className="app-shell">
      <AppHeader />
      <section style={{ display: "grid", gap: 20 }}>
        {orders.map((item: { orderNo: string; totalAmount: number; status: string }) => (
          <OrderCard key={item.orderNo} orderNo={item.orderNo} amount={`¥${(item.totalAmount / 100).toFixed(2)}`} status={item.status} />
        ))}
      </section>
    </main>
  );
}
