import { PrimaryButton } from "../shared/PrimaryButton";

type CartSummaryProps = {
  total: string;
};

export function CartSummary({ total }: CartSummaryProps) {
  return (
    <section className="metric-card">
      <h2>购物车总计</h2>
      <p>{total}</p>
      <PrimaryButton>提交订单</PrimaryButton>
    </section>
  );
}
