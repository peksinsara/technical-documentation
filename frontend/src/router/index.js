import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "../stores/auth";

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/Home.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/documents",
    name: "Documents",
    component: () => import("../views/Documents.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/services",
    name: "Services",
    component: () => import("../views/Services.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/diagrams",
    name: "Diagrams",
    component: () => import("../views/Diagrams.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../components/LoginForm.vue"),
    meta: { requiresAuth: false },
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../components/RegisterForm.vue"),
    meta: { requiresAuth: false },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();
  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

  if (requiresAuth && !authStore.isAuthenticated) {
    next("/login");
  } else if (!requiresAuth && authStore.isAuthenticated) {
    next("/");
  } else {
    next();
  }
});

export default router;
