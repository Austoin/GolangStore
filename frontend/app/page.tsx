import Link from "next/link";

export default function HomePage() {
  return (
    <main style={{ padding: 48 }}>
      <h1>GolangStore Frontend</h1>
      <p>统一的商城前台与管理后台入口。</p>
      <div style={{ display: "flex", gap: 16 }}>
        <Link href="/shop">进入前台</Link>
        <Link href="/admin">进入后台</Link>
      </div>
    </main>
  );
}
