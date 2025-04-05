<template>
  <div class="container mx-auto px-4 py-8 max-w-7xl space-y-6">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-dark-100">Services</h1>
    </div>
    <div class="flex space-x-4 overflow-x-auto pb-4">
      <button
        v-for="category in categories"
        :key="category.id"
        @click="selectedCategory = category.id"
        class="px-4 py-2 rounded-lg whitespace-nowrap"
        :class="
          selectedCategory === category.id
            ? 'bg-primary-600 text-white'
            : 'bg-gray-800 text-gray-300 hover:bg-gray-700'
        "
      >
        {{ category.name }}
      </button>
    </div>

    <div class="flex space-x-4">
      <div class="flex-1">
        <input
          type="text"
          placeholder="Search services..."
          class="input"
          v-model="searchQuery"
        />
      </div>
      <select class="input w-48" v-model="selectedTag">
        <option value="">All Tags</option>
        <option v-for="tag in availableTags" :key="tag" :value="tag">
          {{ tag }}
        </option>
      </select>
    </div>

    <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      <div
        v-for="service in filteredServices"
        :key="service.id"
        class="card hover:bg-dark-700 transition-colors duration-200"
      >
        <div class="flex justify-between items-start">
          <h3 class="text-lg font-medium text-dark-100">{{ service.name }}</h3>
          <span class="text-sm text-dark-400 text-left text-left">{{
            service.category
          }}</span>
        </div>
        <p class="mt-2 text-dark-400 text-left text-left line-clamp-2">
          {{ service.description }}
        </p>

        <div class="mt-4 flex flex-wrap gap-2">
          <span
            v-for="tag in service.tags"
            :key="tag"
            class="px-2 py-1 text-xs font-medium rounded-full bg-primary-100 text-primary-800"
          >
            {{ tag }}
          </span>
        </div>

        <div class="mt-4 flex items-center">
          <div class="flex items-center space-x-2">
            <div
              class="w-2 h-2 rounded-full"
              :class="getHealthClass(service.health.status)"
            ></div>
            <span class="text-sm text-dark-400 text-left text-left">{{
              service.health.message
            }}</span>
          </div>
        </div>

        <div class="mt-4 flex space-x-2">
          <button class="text-sm text-primary-500 hover:text-primary-400">
            View
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useServiceStore } from "../stores/service";
import { useAuthStore } from "../stores/auth";

const serviceStore = useServiceStore();
const authStore = useAuthStore();

const searchQuery = ref("");
const selectedCategory = ref("all");
const selectedTag = ref("");

const categories = [
  { id: "all", name: "All Services" },
  { id: "crmconnector", name: "Crm Connector" },
  { id: "asterisk", name: "Asterisk" },
  { id: "omnichannel", name: "Omnichannel" },
  { id: "clickhouse", name: "Clickhouse" },
];
const availableTags = [
  "api",
  "database",
  "cache",
  "queue",
  "monitoring",
  "auth",
  "storage",
  "search",
];

const services = ref([]);

// Fetch services when component is mounted
onMounted(async () => {
  if (!authStore.isAuthenticated) {
    return;
  }

  try {
    await serviceStore.fetchServices();
    services.value = serviceStore.services;
  } catch (error) {
    console.error("Failed to fetch services:", error);
  }
});

const filteredServices = computed(() => {
  return services.value.filter((service) => {
    const matchesSearch =
      service.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      service.description
        .toLowerCase()
        .includes(searchQuery.value.toLowerCase());
    const matchesCategory =
      selectedCategory.value === "all" ||
      service.category === selectedCategory.value;
    const matchesTag =
      !selectedTag.value || service.tags.includes(selectedTag.value);
    return matchesSearch && matchesCategory && matchesTag;
  });
});

const getHealthClass = (status) => {
  switch (status) {
    case "healthy":
      return "bg-green-500";
    case "degraded":
      return "bg-yellow-500";
    case "unhealthy":
      return "bg-red-500";
    default:
      return "bg-gray-500";
  }
};
</script>
