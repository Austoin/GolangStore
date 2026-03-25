type ProductDetailPanelProps = {
  title: string;
  description: string;
  price: string;
  stock: string;
  action?: React.ReactNode;
};

export function ProductDetailPanel({ title, description, price, stock, action }: ProductDetailPanelProps) {
  return (
    <section className="metric-card" style={{ display: "grid", gap: 18 }}>
      <div>
        <p className="shop-kicker">Product Detail</p>
        <h2 className="shop-display" style={{ fontSize: 46 }}>{title}</h2>
      </div>
      <p className="shop-meta">{description}</p>
      <div className="inline-stack" style={{ justifyContent: "space-between" }}>
        <p style={{ fontSize: 34, fontWeight: 800, margin: 0 }}>{price}</p>
        <p className="shop-meta" style={{ margin: 0 }}>{stock}</p>
      </div>
      <div>{action}</div>
    </section>
  );
}
