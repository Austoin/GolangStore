import { PrimaryButton } from "../shared/PrimaryButton";
import { StatusBadge } from "../shared/StatusBadge";

type ProductCardProps = {
  name: string;
  price: string;
  stock: string;
};

export function ProductCard({ name, price, stock }: ProductCardProps) {
  return (
    <article className="metric-card">
      <StatusBadge label={stock} tone="success" />
      <h3>{name}</h3>
      <p>{price}</p>
      <PrimaryButton>加入购物车</PrimaryButton>
    </article>
  );
}
