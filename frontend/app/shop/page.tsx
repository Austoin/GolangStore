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
      quantity: 1,
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
          <p className="shop-kicker">Shop Experience</p>
          <h2 className="shop-display">浏览商品、加入购物车、立即下单。</h2>
          <p className="shop-meta">当前前台已接入真实商品、购物车和下单接口。</p>
        </div>
        <div className="metric-card">
          <h3 className="section-title">当前可体验动作</h3>
          <div className="inline-stack">
            <span>查看商品</span>
            <span>加入购物车</span>
            <span>提交订单</span>
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
              <form action={addToCart}>
                <input type="hidden" name="product_id" value={item.id} />
                <input type="hidden" name="product_name" value={item.name} />
                <input type="hidden" name="price" value={item.price} />
                <PrimaryButton type="submit">加入购物车</PrimaryButton>
              </form>
            }
          />
        ))}
      </section>
    </main>
  );
}
