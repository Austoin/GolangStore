type ProductAdminTableProps = {
  products: { id: number; name: string; price: number; status: string; stock: number }[];
};

export function ProductAdminTable({ products }: ProductAdminTableProps) {
  return (
    <section className="metric-card">
      <h2>商品管理</h2>
      <ul>
        {products.map((item) => (
          <li key={item.id}>{item.name} / ¥{(item.price / 100).toFixed(2)} / {item.status} / 库存 {item.stock}</li>
        ))}
      </ul>
    </section>
  );
}
