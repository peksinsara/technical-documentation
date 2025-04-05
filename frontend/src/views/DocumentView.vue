<template>
  <div class="container mx-auto px-4 py-8">
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div
        class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary-500"
      ></div>
    </div>

    <div v-else-if="error" class="text-center text-red-500">
      {{ error }}
    </div>

    <div v-else-if="!documentId" class="text-center text-red-500">
      No document ID provided. Please go back to the documents page and try
      again.
    </div>

    <div v-else class="max-w-4xl mx-auto">
      <div class="mb-8">
        <h1 class="text-4xl font-bold mb-4">{{ document.title }}</h1>
        <div class="flex items-center space-x-4 text-dark-400">
          <span>By {{ document.author?.username }}</span>
          <span>•</span>
          <span>{{ formatDate(document.createdAt) }}</span>
        </div>
      </div>

      <div class="prose prose-invert max-w-none">
        <div v-html="document.content"></div>
      </div>

      <div class="mt-8 pt-8 border-t border-dark-700">
        <div class="flex flex-wrap gap-2">
          <span
            v-for="tag in document.tags"
            :key="tag.name"
            class="px-3 py-1 text-sm bg-dark-700 text-dark-300 rounded-full"
          >
            {{ tag.name }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute } from "vue-router";
import { useDocumentStore } from "../stores/document";

const route = useRoute();
const documentStore = useDocumentStore();
const document = ref(null);
const loading = ref(true);
const error = ref(null);

const documentId = computed(() => {
  const id = route.params.id;
  return id;
});

const formatDate = (date) => {
  return new Date(date).toLocaleDateString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
  });
};

onMounted(async () => {
  try {


    if (!documentId.value) {
      error.value = "No document ID provided";
      loading.value = false;
      return;
    }

    const doc = await documentStore.getDocumentById(documentId.value);
    document.value = doc;
  } catch (err) {
    error.value = "Failed to load document";
    console.error("Error loading document:", err);
  } finally {
    loading.value = false;
  }
});
</script>
