import Link from "next/link";
import { redirect } from "next/navigation";
import { ProductAdminTable } from "../../../components/admin/ProductAdminTable";
import { AppHeader } from "../../../components/shared/AppHeader";
import { PrimaryButton } from "../../../components/shared/PrimaryButton";
import { listProductsHttp } from "../../../lib/adapters/httpAdapter";

type AdminProductsPageProps = {
  searchParams: Promise<{ keyword?: string }>;
};

async function createProduct(formData: FormData) {
  "use server";
  await fetch("http://127.0.0.1:8081/products", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      name: String(formData.get("name")),
      description: String(formData.get("description")),
      price: Number(formData.get("price")),
      status: Number(formData.get("status")),
      stock: Number(formData.get("stock")),
    }),
    cache: "no-store",
  });
  redirect("/admin/products");
}

export default async function AdminProductsPage({ searchParams }: AdminProductsPageProps) {
  const params = await searchParams;
  const products = await listProductsHttp();
  const keyword = (params.keyword ?? "").trim().toLowerCase();
  const filtered = keyword ? products.filter((item) => item.name.toLowerCase().includes(keyword)) : products;

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
          <section className="metric-card">
            <form className="inline-stack" method="GET">
              <input name="keyword" placeholder="按商品名筛选" defaultValue={params.keyword ?? ""} style={{ padding: 10, width: 220 }} />
              <PrimaryButton type="submit">筛选</PrimaryButton>
            </form>
          </section>
          <ProductAdminTable
            products={filtered}
            createAction={
              <form action={createProduct} className="inline-stack">
                <input name="name" placeholder="商品名" style={{ padding: 10 }} />
                <input name="description" placeholder="描述" style={{ padding: 10 }} />
                <input name="price" type="number" placeholder="价格(分)" style={{ padding: 10, width: 120 }} />
                <input name="stock" type="number" placeholder="库存" style={{ padding: 10, width: 100 }} />
                <input name="status" type="number" defaultValue="1" style={{ padding: 10, width: 80 }} />
                <PrimaryButton type="submit">新增商品</PrimaryButton>
              </form>
            }
          />
        </section>
      </div>
    </main>
  );
}
