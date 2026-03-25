import { AppHeader } from "../../../components/shared/AppHeader";
import { CartSummary } from "../../../components/shop/CartSummary";

export default function CartPage() {
  return (
    <main className="app-shell">
      <AppHeader />
      <CartSummary total="¥3998.00" />
    </main>
  );
}
