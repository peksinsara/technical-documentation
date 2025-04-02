<template>
  <div class="space-y-6">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-dark-100">Diagrams</h1>
      <button class="btn btn-primary">New Diagram</button>
    </div>

    <div class="flex space-x-4 overflow-x-auto pb-4">
      <button
        v-for="type in diagramTypes"
        :key="type.id"
        class="flex-shrink-0 px-4 py-2 rounded-lg text-sm font-medium"
        :class="
          selectedType === type.id
            ? 'bg-primary-600 text-white'
            : 'bg-dark-800 text-dark-300 hover:bg-dark-700'
        "
        @click="selectedType = type.id"
      >
        {{ type.name }}
      </button>
    </div>

    <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      <div
        v-for="diagram in filteredDiagrams"
        :key="diagram.id"
        class="card group"
      >
        <div
          class="relative aspect-video bg-dark-800 rounded-lg overflow-hidden"
        >
          <img
            :src="diagram.thumbnail"
            :alt="diagram.title"
            class="w-full h-full object-cover"
          />
          <div
            class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-50 transition-opacity duration-200 flex items-center justify-center"
          >
            <button class="opacity-0 group-hover:opacity-100 btn btn-primary">
              Edit Diagram
            </button>
          </div>
        </div>
        <div class="mt-4">
          <h3 class="text-lg font-medium text-dark-100">{{ diagram.title }}</h3>
          <p class="mt-1 text-sm text-dark-400 text-left text-left text-left">
            {{ diagram.description }}
          </p>
        </div>
        <div class="mt-4 flex items-center justify-between">
          <div class="flex items-center space-x-2">
            <img
              :src="diagram.author.avatar"
              :alt="diagram.author.name"
              class="h-6 w-6 rounded-full"
            />
            <span class="text-sm text-dark-400 text-left text-left text-left">{{
              diagram.author.name
            }}</span>
          </div>
          <div class="flex items-center space-x-2">
            <button class="text-sm text-primary-500 hover:text-primary-400">
              <svg
                class="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                />
              </svg>
            </button>
            <button class="text-sm text-primary-500 hover:text-primary-400">
              <svg
                class="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";

const selectedType = ref("all");

const diagramTypes = [
  { id: "all", name: "All Diagrams" },
  { id: "architecture", name: "Architecture" },
  { id: "sequence", name: "Sequence" },
  { id: "flowchart", name: "Flowchart" },
  { id: "er", name: "ER Diagram" },
];

const diagrams = ref([
  {
    id: 1,
    title: "System Architecture",
    description:
      "High-level overview of the system components and their interactions",
    type: "architecture",
    thumbnail:
      "https://via.placeholder.com/400x225/1a1a1a/ffffff?text=System+Architecture",
    author: {
      name: "Sara Peksin",
      avatar: "https://ui-avatars.com/api/?name=John+Doe",
    },
  },
  {
    id: 2,
    title: "User Authentication Flow",
    description: "Sequence diagram showing the user authentication process",
    type: "sequence",
    thumbnail:
      "https://via.placeholder.com/400x225/1a1a1a/ffffff?text=Auth+Flow",
    author: {
      name: "Sara Peksin",
      avatar: "https://ui-avatars.com/api/?name=Jane+Smith",
    },
  },
  {
    id: 3,
    title: "Database Schema",
    description: "ER diagram showing the database structure and relationships",
    type: "er",
    thumbnail:
      "https://via.placeholder.com/400x225/1a1a1a/ffffff?text=DB+Schema",
    author: {
      name: "Sara Peksin",
      avatar: "https://ui-avatars.com/api/?name=Mike+Johnson",
    },
  },
]);

const filteredDiagrams = computed(() => {
  if (selectedType.value === "all") {
    return diagrams.value;
  }
  return diagrams.value.filter(
    (diagram) => diagram.type === selectedType.value
  );
});
</script>
