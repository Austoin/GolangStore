import { redirect } from "next/navigation";
import { AppHeader } from "../../../../components/shared/AppHeader";
import { ProductDetailPanel } from "../../../../components/shop/ProductDetailPanel";
import { getProductDetailHttp } from "../../../../lib/adapters/httpAdapter";

type ProductDetailPageProps = {
  params: Promise<{ id: string }>;
};

async function addToCart(formData: FormData) {
  "use server";
  await fetch("http://127.0.0.1:8083/carts", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      user_id: 1,
      product_id: Number(formData.get("product_id")),
      product_name: String(formData.get("product_name")),
      price: Number(formData.get("price")),
      quantity: Number(formData.get("quantity")),
      checked: true,
    }),
    cache: "no-store",
  });
  redirect("/shop/cart");
}

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
        action={
          <form action={addToCart} className="inline-stack">
            <input type="hidden" name="product_id" value={product.id} />
            <input type="hidden" name="product_name" value={product.name} />
            <input type="hidden" name="price" value={product.price} />
            <label htmlFor="detail-qty">数量</label>
            <input id="detail-qty" name="quantity" type="number" min="1" max={product.stock} defaultValue="1" style={{ width: 72, padding: 8 }} />
            <button className="primary-button" type="submit">加入购物车</button>
          </form>
        }
      />
    </main>
  );
}
