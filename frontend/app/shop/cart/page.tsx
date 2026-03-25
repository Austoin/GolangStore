import { AppHeader } from "../../../components/shared/AppHeader";
import { CartSummary } from "../../../components/shop/CartSummary";
import { listCartItems } from "../../../lib/adapters";

export default function CartPage() {
  const items = listCartItems(1);
  const total = items.reduce((sum, item) => sum + item.price * item.quantity, 0);

  return (
    <main className="app-shell">
      <AppHeader />
      <CartSummary total={`¥${(total / 100).toFixed(2)}`} />
    </main>
  );
}
