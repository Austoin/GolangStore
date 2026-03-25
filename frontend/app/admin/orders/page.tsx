import Link from "next/link";
import { OrderTable } from "../../../components/admin/OrderTable";
import { AppHeader } from "../../../components/shared/AppHeader";
import { listOrdersHttp } from "../../../lib/adapters/httpAdapter";

export default async function AdminOrdersPage() {
  const orders = await listOrdersHttp();
  return (
    <main className="app-shell">
      <AppHeader />
      <div className="admin-shell">
        <aside className="admin-sidebar">
          <h2>订单指挥台</h2>
          <nav>
            <Link href="/admin">总览</Link>
            <Link href="/admin/products">商品</Link>
            <Link href="/admin/inventory">库存</Link>
            <Link href="/admin/orders">订单</Link>
          </nav>
        </aside>
        <section className="admin-main">
          <OrderTable orders={orders} />
        </section>
      </div>
    </main>
  );
}
