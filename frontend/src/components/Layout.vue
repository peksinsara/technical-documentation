<template>
  <div class="min-h-screen bg-dark-900">
    <nav class="bg-dark-800 border-b border-dark-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <router-link to="/" class="text-primary-500 text-xl font-bold">
                TechDocs
              </router-link>
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <router-link
                v-for="item in navigation"
                :key="item.name"
                :to="item.href"
                class="inline-flex items-center px-1 pt-1 text-sm font-medium text-dark-300 hover:text-primary-500"
                :class="{
                  'text-primary-500 border-b-2 border-primary-500':
                    $route.path === item.href,
                }"
              >
                {{ item.name }}
              </router-link>
            </div>
          </div>
          <div class="flex items-center">
            <button class="btn btn-secondary">Sign In</button>
          </div>
        </div>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <slot></slot>
    </main>

    <footer class="bg-dark-800 border-t border-dark-700">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <p class="text-center text-sm text-dark-400 text-left text-left">
          © {{ new Date().getFullYear() }} TechDocs. All rights reserved.
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup>
const navigation = [
  { name: "Home", href: "/" },
  { name: "Documents", href: "/documents" },
  { name: "Services", href: "/services" },
  { name: "Use Cases", href: "/use-cases" },
  { name: "Diagrams", href: "/diagrams" },
];
</script>
