import { AppHeader } from "../../components/shared/AppHeader";
import { ProductCard } from "../../components/shop/ProductCard";
import { PrimaryButton } from "../../components/shared/PrimaryButton";
import { listProductsHttp } from "../../lib/adapters/httpAdapter";
import { redirect } from "next/navigation";

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

export default async function ShopPage() {
  const products = await listProductsHttp();

  return (
    <main className="app-shell">
      <AppHeader />
      <section className="shop-hero">
        <div className="metric-card">
          <p className="shop-kicker">Modern Commerce Hall</p>
          <h2 className="shop-display">从浏览商品到提交订单，体验一套真正可用的商城前台。</h2>
          <p className="shop-meta">当前商品、购物车、订单均接入真实后端服务。你可以直接选择数量并加入购物车。</p>
        </div>
        <div className="metric-card">
          <h3 className="section-title">立即体验</h3>
          <div className="inline-stack">
            <span>查看商品</span>
            <span>数量控制</span>
            <span>真实加购</span>
            <span>真实下单</span>
          </div>
        </div>
      </section>
      <section className="shop-grid">
        {products.map((item) => (
          <ProductCard
            key={item.id}
            id={item.id}
            name={item.name}
            price={`¥${(item.price / 100).toFixed(2)}`}
            stock={`库存 ${item.stock}`}
            addToCartAction={
              <form action={addToCart} style={{ display: "grid", gap: 10 }}>
                <input type="hidden" name="product_id" value={item.id} />
                <input type="hidden" name="product_name" value={item.name} />
                <input type="hidden" name="price" value={item.price} />
                <div className="inline-stack">
                  <label htmlFor={`qty-${item.id}`}>数量</label>
                  <input id={`qty-${item.id}`} name="quantity" type="number" min="1" max={item.stock} defaultValue="1" style={{ width: 72, padding: 8 }} />
                </div>
                <PrimaryButton type="submit">加入购物车</PrimaryButton>
              </form>
            }
          />
        ))}
      </section>
    </main>
  );
}
