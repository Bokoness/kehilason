// Composables
import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/store/auth";

const routes = [
  {
    path: "/",
    component: () => import("@/layouts/default/DefaultLayout.vue"),
    children: [
      {
        path: "",
        name: "Home",
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () =>
          import(/* webpackChunkName: "home" */ "@/views/HomeView.vue"),
      },
      {
        path: "/dashboard",
        name: "Dashboard",
        component: () => import("@/views/DashboardView.vue"),
      },
    ],
  },
  {
    path: "/dashboard",
    component: () => import("@/layouts/dashboard/DashboardLayout.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach(async (to, from, next) => {
  const store = useAuthStore();

  const isAuthenticated = await store.isAuthenticated();

  if (to.name === "Home" && isAuthenticated) next({ name: "Dashboard" });

  if (to.name !== "Home" && !isAuthenticated) next({ name: "Home" });
  else next();
});

export default router;
