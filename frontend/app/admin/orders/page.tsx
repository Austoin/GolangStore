import Link from "next/link";
import { OrderTable } from "../../../components/admin/OrderTable";
import { AppHeader } from "../../../components/shared/AppHeader";
import { PrimaryButton } from "../../../components/shared/PrimaryButton";
import { listOrdersHttp } from "../../../lib/adapters/httpAdapter";

type AdminOrdersPageProps = {
  searchParams: Promise<{ keyword?: string }>;
};

export default async function AdminOrdersPage({ searchParams }: AdminOrdersPageProps) {
  const params = await searchParams;
  const orders = await listOrdersHttp();
  const keyword = (params.keyword ?? "").trim().toLowerCase();
  const filtered = keyword ? orders.filter((item: { orderNo?: string }) => (item.orderNo ?? "").toLowerCase().includes(keyword)) : orders;

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
          <section className="metric-card">
            <form className="inline-stack" method="GET">
              <input name="keyword" placeholder="按订单号筛选" defaultValue={params.keyword ?? ""} style={{ padding: 10, width: 220 }} />
              <PrimaryButton type="submit">筛选</PrimaryButton>
            </form>
          </section>
          <OrderTable orders={filtered} />
        </section>
      </div>
    </main>
  );
}
