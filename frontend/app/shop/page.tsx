import { AppHeader } from "../../components/shared/AppHeader";
import { ProductCard } from "../../components/shop/ProductCard";

export default function ShopPage() {
  return (
    <main className="app-shell">
      <AppHeader />
      <section style={{ display: "grid", gridTemplateColumns: "repeat(3, 1fr)", gap: 24 }}>
        <ProductCard name="Phone X" price="¥1999.00" stock="库存 8" />
        <ProductCard name="Cable" price="¥49.00" stock="库存 25" />
        <ProductCard name="Mouse" price="¥129.00" stock="库存 4" />
      </section>
    </main>
  );
}
