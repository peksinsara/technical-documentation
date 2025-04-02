<template>
  <div class="w-full">
    <div class="space-y-6">
      <div class="flex justify-between items-center">
        <h1 class="text-3xl font-bold text-dark-100 text-left">
          {{ document.title }}
        </h1>
        <div class="flex space-x-4">
          <button class="btn btn-secondary" @click="$emit('close')">
            Close
          </button>
          <button class="btn btn-primary" @click="$emit('edit')">Edit</button>
        </div>
      </div>

      <div class="card">
        <div
          class="flex flex-col sm:flex-row justify-between items-start space-y-4 sm:space-y-0"
        >
          <div class="space-y-2 text-left">
            <div class="flex items-center space-x-2">
              <img
                :src="document.author.avatar"
                :alt="document.author.name"
                class="h-6 w-6 rounded-full"
              />
              <span class="text-sm text-dark-400 text-left text-left">{{
                document.author.name
              }}</span>
            </div>
            <span class="text-sm text-dark-400 text-left text-left"
              >Last updated: {{ document.lastUpdated }}</span
            >
          </div>
          <div class="flex flex-wrap gap-2 justify-start">
            <span
              v-for="tag in document.tags"
              :key="tag"
              class="px-2 py-1 text-xs font-medium rounded-full bg-primary-100 text-primary-800"
            >
              {{ tag }}
            </span>
          </div>
        </div>
      </div>

      <div class="card prose prose-invert max-w-none text-left">
        <div v-html="renderedContent"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { marked } from "marked";
import hljs from "highlight.js";

const props = defineProps({
  document: {
    type: Object,
    required: true,
  },
});

defineEmits(["close", "edit"]);

// Configure marked for syntax highlighting
marked.setOptions({
  highlight: function (code, lang) {
    if (lang && hljs.getLanguage(lang)) {
      return hljs.highlight(code, { language: lang }).value;
    }
    return hljs.highlightAuto(code).value;
  },
});

const renderedContent = computed(() => {
  return marked(props.document.content);
});
</script>

<style>
.prose {
  @apply text-dark-100 text-left;
  max-width: 100%;
  overflow-x: hidden;
}

.prose pre {
  @apply bg-gray-800 p-4 rounded-lg overflow-x-auto text-left;
  max-width: 100%;
}

.prose img {
  @apply max-w-full h-auto;
}

.prose table {
  @apply w-full overflow-x-auto block text-left;
}

.prose code {
  @apply break-words text-left;
  word-break: break-word;
}

.prose h1,
.prose h2,
.prose h3,
.prose h4 {
  @apply text-dark-100 mt-6 mb-4 text-left;
}

.prose p {
  @apply text-dark-300 mb-4 text-left;
}

.prose ul,
.prose ol {
  @apply text-dark-300 mb-4 pl-4 text-left;
}

.prose li {
  @apply mb-2 text-left;
}

.prose a {
  @apply text-primary-500 hover:text-primary-400 text-left;
}

.prose blockquote {
  @apply border-l-4 border-primary-500 pl-4 italic text-dark-300 my-4 text-left;
}
</style>
