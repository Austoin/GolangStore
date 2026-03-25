type ProductAdminTableProps = {
  products: { id: number; name: string; price: number; status: string | number; stock: number }[];
  createAction?: React.ReactNode;
};

export function ProductAdminTable({ products, createAction }: ProductAdminTableProps) {
  return (
    <section className="metric-card">
      <h2 className="section-title">商品管理</h2>
      <div style={{ marginBottom: 20 }}>{createAction}</div>
      <table className="panel-table">
        <thead><tr><th>商品</th><th>价格</th><th>状态</th><th>库存</th></tr></thead>
        <tbody>
          {products.map((item) => (
            <tr key={item.id}><td>{item.name}</td><td>¥{(item.price / 100).toFixed(2)}</td><td>{item.status}</td><td>{item.stock}</td></tr>
          ))}
        </tbody>
      </table>
    </section>
  );
}
