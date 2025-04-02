<script setup>
import { computed } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "./stores/auth";

const router = useRouter();
const authStore = useAuthStore();

const isAuthenticated = computed(() => authStore.isAuthenticated);

const navigation = [
  { name: "Home", href: "/", label: "Home" },
  { name: "Documents", href: "/documents", label: "Documents" },
  { name: "Services", href: "/services", label: "Services" },
  { name: "Diagrams", href: "/diagrams", label: "Diagrams" },
];

const logout = async () => {
  await authStore.logout();
  router.push("/login");
};
</script>

<template>
  <div class="min-h-screen bg-dark-900">
    <nav class="bg-dark-800 border-b border-dark-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <h1 class="text-xl font-bold text-dark-100">Technical Docs</h1>
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <router-link
                v-for="item in navigation"
                :key="item.name"
                :to="item.href"
                class="inline-flex items-center px-1 pt-1 text-sm font-medium"
                :class="[
                  $route.name === item.name
                    ? 'text-primary-500 border-b-2 border-primary-500'
                    : 'text-dark-400 text-left text-left hover:text-dark-300 hover:border-dark-300',
                  'border-b-2 border-transparent',
                ]"
              >
                {{ item.label }}
              </router-link>
            </div>
          </div>
          <div class="flex items-center">
            <template v-if="isAuthenticated">
              <button
                @click="logout"
                class="text-dark-400 text-left text-left hover:text-dark-300 text-sm font-medium"
              >
                Logout
              </button>
            </template>
            <template v-else>
              <router-link
                to="/login"
                class="text-dark-400 text-left text-left hover:text-dark-300 text-sm font-medium mr-4"
              >
                Login
              </router-link>
              <router-link to="/register" class="btn btn-primary">
                Register
              </router-link>
            </template>
          </div>
        </div>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <router-view></router-view>
    </main>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}

.nav-link {
  @apply px-3 py-2 rounded-md text-sm font-medium text-gray-300 hover:text-white hover:bg-gray-700;
}

.nav-link-active {
  @apply bg-gray-900 text-white;
}

.btn {
  @apply px-4 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white;
}

.btn-primary {
  @apply bg-blue-600 text-white hover:bg-blue-700;
}

.btn-secondary {
  @apply bg-gray-700 text-white hover:bg-gray-600;
}

.input {
  @apply appearance-none relative block w-full px-3 py-2 border border-gray-700 placeholder-gray-500 text-gray-100 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm bg-gray-700;
}
</style>
