import { AppHeader } from "../../components/shared/AppHeader";
import { ProductCard } from "../../components/shop/ProductCard";
import { listProducts } from "../../lib/adapters";

export default function ShopPage() {
  const products = listProducts();

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
