import { redirect } from "next/navigation";
import { AppHeader } from "../../../components/shared/AppHeader";
import { CartSummary } from "../../../components/shop/CartSummary";
import { PrimaryButton } from "../../../components/shared/PrimaryButton";
import { createOrderFromCartHttp, listCartItemsHttp } from "../../../lib/adapters/httpAdapter";

async function submitOrder() {
  "use server";
  await createOrderFromCartHttp(1);
  redirect("/shop/orders");
}

export default async function CartPage() {
  const items = await listCartItemsHttp(1);
  const total = items.reduce((sum: number, item: { price: number; quantity: number }) => sum + item.price * item.quantity, 0);

  return (
    <main className="app-shell">
      <AppHeader />
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
