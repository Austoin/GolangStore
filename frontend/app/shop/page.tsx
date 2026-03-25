import { AppHeader } from "../../components/shared/AppHeader";
import { ProductCard } from "../../components/shop/ProductCard";
import { listProductsHttp } from "../../lib/adapters/httpAdapter";

export default async function ShopPage() {
  const products = await listProductsHttp();

  return (
    <main className="app-shell">
      <AppHeader />
      <section style={{ display: "grid", gridTemplateColumns: "repeat(3, 1fr)", gap: 24 }}>
        {products.map((item) => (
          <ProductCard key={item.id} name={item.name} price={`¥${(item.price / 100).toFixed(2)}`} stock={`库存 ${item.stock}`} />
        ))}
      </section>
    </main>
  );
}
