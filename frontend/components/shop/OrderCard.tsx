import { StatusBadge } from "../shared/StatusBadge";

type OrderCardProps = {
  orderNo: string;
  amount: string;
  status: string;
  details?: string;
};

export function OrderCard({ orderNo, amount, status, details }: OrderCardProps) {
  return (
    <article className="metric-card" style={{ display: "grid", gap: 12 }}>
      <div className="inline-stack" style={{ justifyContent: "space-between" }}>
        <StatusBadge label={status} tone="warning" />
        <span className="shop-meta">订单号</span>
      </div>
      <h3 className="shop-card-title">{orderNo}</h3>
      <p style={{ margin: 0, fontSize: 28, fontWeight: 800 }}>{amount}</p>
      {details ? <p className="shop-meta">{details}</p> : null}
    </article>
  );
}
