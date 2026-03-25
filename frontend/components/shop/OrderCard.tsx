import { StatusBadge } from "../shared/StatusBadge";

type OrderCardProps = {
  orderNo: string;
  amount: string;
  status: string;
};

export function OrderCard({ orderNo, amount, status }: OrderCardProps) {
  return (
    <article className="metric-card">
      <StatusBadge label={status} tone="warning" />
      <h3>{orderNo}</h3>
      <p>{amount}</p>
    </article>
  );
}
