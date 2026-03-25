import Link from "next/link";

export function AppHeader() {
  return (
    <header className="app-header">
      <div>
        <p className="eyebrow">GolangStore Experience</p>
        <h1 className="brand-title">现代电商展厅与运营指挥台</h1>
      </div>
      <nav className="header-nav">
        <Link href="/shop">用户前台</Link>
        <Link href="/admin">管理后台</Link>
      </nav>
    </header>
  );
}
