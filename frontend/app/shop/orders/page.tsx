import { AppHeader } from "../../../components/shared/AppHeader";
import { OrderCard } from "../../../components/shop/OrderCard";
import { listOrders } from "../../../lib/adapters";

export default function OrdersPage() {
  const orders = listOrders();

  return (
    <main className="app-shell">
      <AppHeader />
      <section style={{ display: "grid", gap: 20 }}>
        {orders.map((item) => (
          <OrderCard key={item.orderNo} orderNo={item.orderNo} amount={`¥${(item.totalAmount / 100).toFixed(2)}`} status={item.status} />
        ))}
      </section>
    </main>
  );
}
