import { PrimaryButton } from "../shared/PrimaryButton";

type ProductDetailPanelProps = {
  title: string;
  description: string;
  price: string;
  stock: string;
};

export function ProductDetailPanel({ title, description, price, stock }: ProductDetailPanelProps) {
  return (
    <section className="metric-card">
      <h2>{title}</h2>
      <p>{description}</p>
      <p>{price}</p>
      <p>{stock}</p>
      <PrimaryButton>加入购物车</PrimaryButton>
    </section>
  );
}
