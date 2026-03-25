import { redirect } from "next/navigation";
import { AppHeader } from "../../../components/shared/AppHeader";
import { CartSummary } from "../../../components/shop/CartSummary";
import { PrimaryButton } from "../../../components/shared/PrimaryButton";
import { createOrderFromCartHttp, listCartItemsHttp } from "../../../lib/adapters/httpAdapter";

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

export default async function CartPage({ searchParams }: CartPageProps) {
  const params = await searchParams;
  const items = await listCartItemsHttp(1);
  const total = items.reduce((sum: number, item: { price: number; quantity: number }) => sum + item.price * item.quantity, 0);

  return (
    <main className="app-shell">
      <AppHeader />
      {params.error ? <section className="metric-card"><p style={{ color: "#a0382b" }}>{decodeURIComponent(params.error)}</p></section> : null}
      <CartSummary
        total={`¥${(total / 100).toFixed(2)}`}
        action={
          <form action={submitOrder}>
            <PrimaryButton type="submit">提交订单</PrimaryButton>
          </form>
        }
      />
    </main>
  );
}
