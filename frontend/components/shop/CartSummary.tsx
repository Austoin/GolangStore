import { PrimaryButton } from "../shared/PrimaryButton";

type CartSummaryProps = {
  total: string;
  action?: React.ReactNode;
};

export function CartSummary({ total, action }: CartSummaryProps) {
  return (
    <section className="metric-card">
      <h2>购物车总计</h2>
      <p>{total}</p>
      {action ?? <PrimaryButton>提交订单</PrimaryButton>}
    </section>
  );
}
