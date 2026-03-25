import Link from "next/link";

export function AppHeader() {
  return (
    <header className="app-header">
      <div>
        <p className="eyebrow">GolangStore</p>
        <h1 className="brand-title">商城前后台体验台</h1>
      </div>
      <nav className="header-nav">
        <Link href="/shop">Shop</Link>
        <Link href="/admin">Admin</Link>
      </nav>
    </header>
  );
}
