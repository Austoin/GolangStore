type InventoryTableProps = {
  products: { id: number; name: string; stock: number }[];
};

export function InventoryTable({ products }: InventoryTableProps) {
  return (
    <section className="metric-card">
      <h2 className="section-title">库存面板</h2>
      <table className="panel-table">
        <thead><tr><th>商品</th><th>库存</th><th>状态</th></tr></thead>
        <tbody>
          {products.map((item) => (
            <tr key={item.id}><td>{item.name}</td><td>{item.stock}</td><td>{item.stock <= 5 ? '低库存' : '正常'}</td></tr>
          ))}
        </tbody>
      </table>
    </section>
  );
}
