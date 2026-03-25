import { AppHeader } from "../../../../components/shared/AppHeader";
import { ProductDetailPanel } from "../../../../components/shop/ProductDetailPanel";

export default function ProductDetailPage() {
  return (
    <main className="app-shell">
      <AppHeader />
      <ProductDetailPanel
        title="Phone X"
        description="面向本地演示环境的旗舰商品详情页。"
        price="¥1999.00"
        stock="库存 8"
      />
    </main>
  );
}
