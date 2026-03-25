import { OrderTable } from "../../../components/admin/OrderTable";
import { AppHeader } from "../../../components/shared/AppHeader";

export default function AdminOrdersPage() {
  return (
    <main className="app-shell">
      <AppHeader />
      <OrderTable />
    </main>
  );
}
