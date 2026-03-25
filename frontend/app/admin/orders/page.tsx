import { OrderTable } from "../../../components/admin/OrderTable";
import { AppHeader } from "../../../components/shared/AppHeader";
import { listOrdersHttp } from "../../../lib/adapters/httpAdapter";

export default async function AdminOrdersPage() {
  const orders = await listOrdersHttp();
  return (
    <main className="app-shell">
      <AppHeader />
      <OrderTable orders={orders} />
    </main>
  );
}
