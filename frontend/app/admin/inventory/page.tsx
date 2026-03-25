import { InventoryTable } from "../../../components/admin/InventoryTable";
import { AppHeader } from "../../../components/shared/AppHeader";

export default function AdminInventoryPage() {
  return (
    <main className="app-shell">
      <AppHeader />
      <InventoryTable />
    </main>
  );
}
