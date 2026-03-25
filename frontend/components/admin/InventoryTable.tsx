type InventoryTableProps = {
  products: { id: number; name: string; stock: number }[];
};

export function InventoryTable({ products }: InventoryTableProps) {
  return (
    <section className="metric-card">
      <h2>库存面板</h2>
      <ul>
        {products.map((item) => (
          <li key={item.id}>{item.name} / 库存 {item.stock}</li>
        ))}
      </ul>
    </section>
  );
}
