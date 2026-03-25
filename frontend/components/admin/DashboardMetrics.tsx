import { MetricCard } from "../shared/MetricCard";

type DashboardMetricsProps = {
  totalProducts: number;
  totalOrders: number;
  lowStockCount: number;
};

export function DashboardMetrics({ totalProducts, totalOrders, lowStockCount }: DashboardMetricsProps) {
  return (
    <section style={{ display: "grid", gridTemplateColumns: "repeat(4, 1fr)", gap: 20 }}>
      <MetricCard label="商品数" value={String(totalProducts)} note="真实商品" />
      <MetricCard label="订单数" value={String(totalOrders)} note="真实订单" />
      <MetricCard label="低库存" value={String(lowStockCount)} note="库存 <= 5" />
      <MetricCard label="服务状态" value="OK" note="核心服务正常" />
    </section>
  );
}
