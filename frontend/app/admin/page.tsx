import Link from "next/link";
import { DashboardMetrics } from "../../components/admin/DashboardMetrics";
import { AppHeader } from "../../components/shared/AppHeader";
import { listOrdersHttp, listProductsHttp } from "../../lib/adapters/httpAdapter";

export default async function AdminDashboardPage() {
  const [products, orders] = await Promise.all([listProductsHttp(), listOrdersHttp()]);
  const lowStockCount = products.filter((item) => item.stock <= 5).length;

  return (
    <main className="app-shell">
      <AppHeader />
      <div className="admin-shell">
        <aside className="admin-sidebar">
          <h2>运营指挥台</h2>
          <nav>
            <Link href="/admin">总览</Link>
            <Link href="/admin/products">商品</Link>
            <Link href="/admin/inventory">库存</Link>
            <Link href="/admin/orders">订单</Link>
          </nav>
        </aside>
        <section className="admin-main">
          <DashboardMetrics totalProducts={products.length} totalOrders={orders.length} lowStockCount={lowStockCount} />
          <section className="metric-card">
            <h2 className="section-title">实时摘要</h2>
            <div className="inline-stack">
              <p>商品总数：{products.length}</p>
              <p>订单总数：{orders.length}</p>
              <p>低库存：{lowStockCount}</p>
            </div>
          </section>
        </section>
      </div>
    </main>
  );
}
