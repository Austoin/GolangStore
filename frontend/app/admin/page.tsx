import { DashboardMetrics } from "../../components/admin/DashboardMetrics";
import { AppHeader } from "../../components/shared/AppHeader";

export default function AdminDashboardPage() {
  return (
    <main className="app-shell">
      <AppHeader />
      <DashboardMetrics />
    </main>
  );
}
