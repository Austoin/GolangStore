import Link from "next/link";
import { PrimaryButton } from "../shared/PrimaryButton";
import { StatusBadge } from "../shared/StatusBadge";

type ProductCardProps = {
  id: number;
  name: string;
  price: string;
  stock: string;
  addToCartAction?: React.ReactNode;
};

export function ProductCard({ id, name, price, stock, addToCartAction }: ProductCardProps) {
  return (
    <article className="metric-card">
      <StatusBadge label={stock} tone="success" />
      <h3 className="shop-card-title">{name}</h3>
      <p className="shop-meta">{price}</p>
      <div className="inline-stack">
        <Link href={`/shop/products/${id}`}>查看商品</Link>
      </div>
      <div style={{ marginTop: 12 }}>{addToCartAction ?? <PrimaryButton>加入购物车</PrimaryButton>}</div>
    </article>
  );
}
