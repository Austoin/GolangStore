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
    <article className="metric-card" style={{ display: "grid", gap: 14 }}>
      <div className="inline-stack" style={{ justifyContent: "space-between" }}>
        <StatusBadge label={stock} tone="success" />
        <Link href={`/shop/products/${id}`}>查看商品</Link>
      </div>
      <div>
        <h3 className="shop-card-title">{name}</h3>
        <p className="shop-meta">高频电商场景下的真实商品展示卡片</p>
      </div>
      <p style={{ fontSize: 28, fontWeight: 800, margin: 0 }}>{price}</p>
      <div>{addToCartAction ?? <PrimaryButton>加入购物车</PrimaryButton>}</div>
    </article>
  );
}
