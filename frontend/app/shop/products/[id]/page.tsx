import { AppHeader } from "../../../../components/shared/AppHeader";
import { ProductDetailPanel } from "../../../../components/shop/ProductDetailPanel";
import { getProductDetailHttp } from "../../../../lib/adapters/httpAdapter";

type ProductDetailPageProps = {
  params: Promise<{ id: string }>;
};

export default async function ProductDetailPage({ params }: ProductDetailPageProps) {
  const { id } = await params;
  const product = await getProductDetailHttp(Number(id));

  return (
    <main className="app-shell">
      <AppHeader />
      <ProductDetailPanel
        title={product.name}
        description={product.description}
        price={`¥${(product.price / 100).toFixed(2)}`}
        stock={`库存 ${product.stock}`}
      />
    </main>
  );
}
