import { AppHeader } from "../../../components/shared/AppHeader";
import { OrderCard } from "../../../components/shop/OrderCard";

export default function OrdersPage() {
  return (
    <main className="app-shell">
      <AppHeader />
      <section style={{ display: "grid", gap: 20 }}>
        <OrderCard orderNo="O1001" amount="¥3998.00" status="待支付" />
      </section>
    </main>
  );
}
