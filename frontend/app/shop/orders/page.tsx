import { AppHeader } from "../../../components/shared/AppHeader";
import { OrderCard } from "../../../components/shop/OrderCard";
import { listOrdersHttp } from "../../../lib/adapters/httpAdapter";

export default async function OrdersPage() {
  const orders = await listOrdersHttp();

  return (
    <main className="app-shell">
      <AppHeader />
      <section className="shop-hero">
        <div className="metric-card">
          <p className="shop-kicker">Orders</p>
          <h2 className="shop-display" style={{ fontSize: 44 }}>查看每一笔订单的金额、状态与商品明细。</h2>
          <p className="shop-meta">订单页已经接入真实后端订单列表接口。</p>
        </div>
        <div className="metric-card">
          <h3 className="section-title">当前订单总数</h3>
          <p style={{ fontSize: 40, fontWeight: 800, margin: 0 }}>{orders.length}</p>
        </div>
      </section>
      <section style={{ display: "grid", gap: 20 }}>
        {orders.map((item: { orderNo: string; totalAmount: number; status: string; items?: { productName: string; quantity: number }[] }) => (
          <OrderCard
            key={item.orderNo}
            orderNo={item.orderNo}
            amount={`¥${(item.totalAmount / 100).toFixed(2)}`}
            status={item.status}
            details={item.items?.map((detail) => `${detail.productName} x${detail.quantity}`).join(' / ')}
          />
        ))}
      </section>
    </main>
  );
}
