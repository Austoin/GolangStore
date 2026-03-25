import { DashboardMetrics } from "../../components/admin/DashboardMetrics";
import { AppHeader } from "../../components/shared/AppHeader";
import { listOrdersHttp, listProductsHttp } from "../../lib/adapters/httpAdapter";

export default async function AdminDashboardPage() {
  const [products, orders] = await Promise.all([listProductsHttp(), listOrdersHttp()]);
  const lowStockCount = products.filter((item) => item.stock <= 5).length;

  return (
    <main className="app-shell">
      <AppHeader />
      <DashboardMetrics totalProducts={products.length} totalOrders={orders.length} lowStockCount={lowStockCount} />
    </main>
  );
}
