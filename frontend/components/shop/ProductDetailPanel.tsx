type ProductDetailPanelProps = {
  title: string;
  description: string;
  price: string;
  stock: string;
  action?: React.ReactNode;
};

export function ProductDetailPanel({ title, description, price, stock, action }: ProductDetailPanelProps) {
  return (
    <section className="metric-card">
      <h2>{title}</h2>
      <p>{description}</p>
      <p>{price}</p>
      <p>{stock}</p>
      <div style={{ marginTop: 16 }}>{action}</div>
    </section>
  );
}
