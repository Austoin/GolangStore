import { MetricCard } from "../shared/MetricCard";

export function DashboardMetrics() {
  return (
    <section style={{ display: "grid", gridTemplateColumns: "repeat(4, 1fr)", gap: 20 }}>
      <MetricCard label="商品数" value="24" note="在售商品" />
      <MetricCard label="订单数" value="128" note="今日订单" />
      <MetricCard label="低库存" value="3" note="需要关注" />
      <MetricCard label="服务状态" value="OK" note="核心服务正常" />
    </section>
  );
}
