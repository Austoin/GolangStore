import { AppHeader } from "../../../../components/shared/AppHeader";
import { ProductDetailPanel } from "../../../../components/shop/ProductDetailPanel";
import { getProductDetail } from "../../../../lib/adapters";

export default function ProductDetailPage() {
  const product = getProductDetail(101);

  return (
    <main className="app-shell">
      <AppHeader />
      <ProductDetailPanel
        title={product?.name ?? "Unknown Product"}
        description={product?.description ?? "暂无描述"}
        price={`¥${((product?.price ?? 0) / 100).toFixed(2)}`}
        stock={`库存 ${product?.stock ?? 0}`}
      />
    </main>
  );
}
