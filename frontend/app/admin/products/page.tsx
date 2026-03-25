import { ProductAdminTable } from "../../../components/admin/ProductAdminTable";
import { AppHeader } from "../../../components/shared/AppHeader";

export default function AdminProductsPage() {
  return (
    <main className="app-shell">
      <AppHeader />
      <ProductAdminTable />
    </main>
  );
}
