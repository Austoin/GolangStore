import Link from "next/link";
import { ProductAdminTable } from "../../../components/admin/ProductAdminTable";
import { AppHeader } from "../../../components/shared/AppHeader";
import { listProductsHttp } from "../../../lib/adapters/httpAdapter";

export default async function AdminProductsPage() {
  const products = await listProductsHttp();
  return (
    <main className="app-shell">
      <AppHeader />
      <div className="admin-shell">
        <aside className="admin-sidebar">
          <h2>商品指挥台</h2>
          <nav>
            <Link href="/admin">总览</Link>
            <Link href="/admin/products">商品</Link>
            <Link href="/admin/inventory">库存</Link>
            <Link href="/admin/orders">订单</Link>
          </nav>
        </aside>
        <section className="admin-main">
          <ProductAdminTable products={products} />
        </section>
      </div>
    </main>
  );
}
