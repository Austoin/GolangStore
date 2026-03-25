import { ProductAdminTable } from "../../../components/admin/ProductAdminTable";
import { AppHeader } from "../../../components/shared/AppHeader";
import { listProductsHttp } from "../../../lib/adapters/httpAdapter";

export default async function AdminProductsPage() {
  const products = await listProductsHttp();
  return (
    <main className="app-shell">
      <AppHeader />
      <ProductAdminTable products={products} />
    </main>
  );
}
