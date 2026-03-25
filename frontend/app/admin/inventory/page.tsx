import { InventoryTable } from "../../../components/admin/InventoryTable";
import { AppHeader } from "../../../components/shared/AppHeader";
import { listProductsHttp } from "../../../lib/adapters/httpAdapter";

export default async function AdminInventoryPage() {
  const products = await listProductsHttp();
  return (
    <main className="app-shell">
      <AppHeader />
      <InventoryTable products={products} />
    </main>
  );
}
