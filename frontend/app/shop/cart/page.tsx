import { redirect } from "next/navigation";
import { AppHeader } from "../../../components/shared/AppHeader";
import { CartSummary } from "../../../components/shop/CartSummary";
import { PrimaryButton } from "../../../components/shared/PrimaryButton";
import { createOrderFromCartHttp, deleteCartItemHttp, listCartItemsHttp, setCartItemCheckedHttp } from "../../../lib/adapters/httpAdapter";

type CartPageProps = {
  searchParams: Promise<{ error?: string }>;
};

async function submitOrder() {
  "use server";
  try {
    await createOrderFromCartHttp(1);
    redirect("/shop/orders");
  } catch (error) {
    const message = error instanceof Error ? error.message : "提交订单失败";
    redirect(`/shop/cart?error=${encodeURIComponent(message)}`);
  }
}

async function deleteItem(formData: FormData) {
  "use server";
  await deleteCartItemHttp(1, Number(formData.get("product_id")));
  redirect("/shop/cart");
}

async function toggleChecked(formData: FormData) {
  "use server";
  await setCartItemCheckedHttp(1, Number(formData.get("product_id")), formData.get("checked") === "true");
  redirect("/shop/cart");
}

export default async function CartPage({ searchParams }: CartPageProps) {
  const params = await searchParams;
  const items = await listCartItemsHttp(1);
  const total = items.filter((item: { checked: boolean }) => item.checked).reduce((sum: number, item: { price: number; quantity: number }) => sum + item.price * item.quantity, 0);

  return (
    <main className="app-shell">
      <AppHeader />
      <section className="shop-hero">
        <div className="metric-card">
          <p className="shop-kicker">Cart</p>
          <h2 className="shop-display" style={{ fontSize: 44 }}>管理购物车商品、勾选状态与最终订单金额。</h2>
          <p className="shop-meta">你可以删除商品、切换勾选状态，并直接提交真实订单。</p>
        </div>
        <CartSummary
          total={`¥${(total / 100).toFixed(2)}`}
          action={
            <form action={submitOrder}>
              <PrimaryButton type="submit">提交订单</PrimaryButton>
            </form>
          }
        />
      </section>
      {params.error ? <section className="metric-card"><p style={{ color: "#a0382b" }}>{decodeURIComponent(params.error)}</p></section> : null}
      <section className="metric-card">
        <h2 className="section-title">购物车商品</h2>
        <table className="panel-table">
          <thead><tr><th>勾选</th><th>商品</th><th>单价</th><th>数量</th><th>小计</th><th>操作</th></tr></thead>
          <tbody>
            {items.map((item: { productId: number; productName: string; price: number; quantity: number; checked: boolean }) => (
              <tr key={item.productId}>
                <td>
                  <form action={toggleChecked}>
                    <input type="hidden" name="product_id" value={item.productId} />
                    <input type="hidden" name="checked" value={String(!item.checked)} />
                    <button type="submit">{item.checked ? '已选' : '未选'}</button>
                  </form>
                </td>
                <td>{item.productName}</td>
                <td>¥{(item.price / 100).toFixed(2)}</td>
                <td>{item.quantity}</td>
                <td>¥{((item.price * item.quantity) / 100).toFixed(2)}</td>
                <td>
                  <form action={deleteItem}>
                    <input type="hidden" name="product_id" value={item.productId} />
                    <button type="submit">删除</button>
                  </form>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </section>
    </main>
  );
}
