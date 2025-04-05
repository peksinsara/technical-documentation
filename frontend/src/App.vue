<script setup>
import { ref, computed } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useAuthStore } from "./stores/auth";
import { useDocumentStore } from "./stores/document";
import DocumentEditor from "./components/DocumentEditor.vue";

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const documentStore = useDocumentStore();
const showNewDocumentModal = ref(false);
const isProfileMenuOpen = ref(false);

const isAuthenticated = computed(() => authStore.isAuthenticated);
const user = computed(() => authStore.user);

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

const showNewDocumentButton = computed(() => {
  const currentPath = route.path;
  return ["/documents", "/services", "/diagrams"].includes(currentPath);
});

const handleSaveDocument = async (document) => {
  try {
    await documentStore.createDocument(document);
    showNewDocumentModal.value = false;
    router.push("/documents");
  } catch (error) {
    console.error("Failed to create document:", error);
  }
};
</script>

<template>
  <div class="min-h-screen bg-dark-900">
    <nav class="bg-dark-800 border-b border-dark-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <router-link to="/" class="text-xl font-bold text-dark-100">
                TechDocs
              </router-link>
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <router-link
                to="/documents"
                class="border-transparent text-dark-400 hover:border-dark-300 hover:text-dark-300 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                active-class="border-primary-500 text-dark-100"
              >
                Documents
              </router-link>
              <router-link
                to="/services"
                class="border-transparent text-dark-400 hover:border-dark-300 hover:text-dark-300 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                active-class="border-primary-500 text-dark-100"
              >
                Services
              </router-link>
              <router-link
                to="/diagrams"
                class="border-transparent text-dark-400 hover:border-dark-300 hover:text-dark-300 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                active-class="border-primary-500 text-dark-100"
              >
                Diagrams
              </router-link>
            </div>
          </div>
          <div class="flex items-center">
            <button
              v-if="isAuthenticated && showNewDocumentButton"
              @click="showNewDocumentModal = true"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
            >
              New Document
            </button>
            <div class="ml-3 relative">
              <div v-if="isAuthenticated">
                <button
                  @click="isProfileMenuOpen = !isProfileMenuOpen"
                  class="max-w-xs bg-dark-700 flex items-center text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
                >
                  <span class="sr-only">Open user menu</span>
                  <img
                    class="h-8 w-8 rounded-full"
                    :src="
                      user?.avatar ||
                      'https://ui-avatars.com/api/?name=' + user?.username
                    "
                    :alt="user?.username"
                  />
                </button>
                <!-- User menu dropdown -->
                <div
                  v-if="isProfileMenuOpen"
                  class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-dark-800 ring-1 ring-black ring-opacity-5 focus:outline-none"
                  role="menu"
                  aria-orientation="vertical"
                  aria-labelledby="user-menu"
                >
                  <div class="px-4 py-2 border-b border-dark-700">
                    <p class="text-sm text-dark-100">{{ user?.username }}</p>
                    <p class="text-xs text-dark-400">{{ user?.email }}</p>
                  </div>
                  <a
                    href="#"
                    class="block px-4 py-2 text-sm text-dark-400 hover:bg-dark-700 hover:text-dark-300"
                    role="menuitem"
                    @click="logout"
                  >
                    Logout
                  </a>
                </div>
              </div>
              <div v-else>
                <router-link
                  to="/login"
                  class="text-dark-400 hover:text-dark-300 px-3 py-2 rounded-md text-sm font-medium"
                >
                  Login
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <router-view></router-view>
    </main>

    <!-- New Document Modal -->
    <div
      v-if="showNewDocumentModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4"
    >
      <div
        class="bg-dark-800 rounded-lg shadow-xl max-w-5xl w-full max-h-[90vh] overflow-y-auto"
      >
        <div class="px-6 py-4 border-b border-dark-700">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-medium text-dark-100">
              Create New Document
            </h3>
            <button
              @click="showNewDocumentModal = false"
              class="text-dark-400 hover:text-dark-300"
            >
              <span class="sr-only">Close</span>
              <svg
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </div>
        </div>
        <div class="px-6 py-4">
          <DocumentEditor
            @save="handleSaveDocument"
            @cancel="showNewDocumentModal = false"
          />
        </div>
      </div>
    </div>
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
  @apply px-3 py-2 rounded-md text-sm font-medium text-dark-400 hover:text-dark-300 hover:bg-dark-700;
}

.nav-link-active {
  @apply bg-dark-900 text-dark-100;
}

.btn {
  @apply px-4 py-2 rounded-md text-sm font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-dark-800 focus:ring-white;
}

.btn-primary {
  @apply bg-primary-600 text-white hover:bg-primary-700;
}

.btn-secondary {
  @apply bg-dark-700 text-white hover:bg-dark-600;
}

.input {
  @apply appearance-none relative block w-full px-3 py-2 border border-dark-700 placeholder-dark-500 text-dark-100 rounded-md focus:outline-none focus:ring-primary-500 focus:border-primary-500 focus:z-10 sm:text-sm bg-dark-700;
}
</style>
