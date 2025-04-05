<template>
  <div class="space-y-6">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-dark-100">
        {{ isEditing ? "Edit Document" : "New Document" }}
      </h1>
      <div class="flex space-x-4">
        <button class="btn btn-secondary" @click="$emit('cancel')">
          Cancel
        </button>
        <button class="btn btn-primary" @click="handleSave" :disabled="saving">
          {{ saving ? "Saving..." : "Save Document" }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-1 space-y-6">
        <div class="card">
          <h2 class="text-lg font-medium text-dark-100 mb-4">
            Document Details
          </h2>

          <div class="space-y-2">
            <label class="block text-sm font-medium text-dark-300">Title</label>
            <input
              v-model="document.title"
              type="text"
              class="input"
              placeholder="Enter document title"
            />
          </div>

          <div class="space-y-2 mt-4">
            <label class="block text-sm font-medium text-dark-300">Type</label>
            <select v-model="document.type" class="input">
              <option value="">Select a type</option>
              <option value="service">Service</option>
              <option value="document">Document</option>
              <option value="diagram">Diagram</option>
            </select>
          </div>

          <div v-if="document.type === 'service'" class="space-y-2 mt-4">
            <label class="block text-sm font-medium text-dark-300"
              >Service</label
            >
            <select v-model="document.serviceId" class="input">
              <option value="">Select a service</option>
              <option
                v-for="service in services"
                :key="service.ID"
                :value="service.ID"
              >
                {{ service.name }}
              </option>
            </select>
            <p v-if="services.length === 0" class="text-xs text-red-500 mt-1">
              No services available. Please contact an administrator.
            </p>
          </div>

          <div class="space-y-2 mt-4">
            <label class="block text-sm font-medium text-dark-300"
              >Category</label
            >
            <select v-model="document.category" class="input">
              <option value="">Select a category</option>
              <option value="Documentation">Documentation</option>
              <option value="Technical">Technical</option>
              <option value="Guide">Guide</option>
              <option value="API">API</option>
              <option value="Architecture">Architecture</option>
            </select>
          </div>

          <div class="space-y-2 mt-4">
            <label class="block text-sm font-medium text-dark-300">Tags</label>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="tag in document.tags"
                :key="tag"
                class="px-2 py-1 text-xs font-medium rounded-full bg-primary-100 text-primary-800 flex items-center"
              >
                {{ tag }}
                <button
                  @click="removeTag(tag)"
                  class="ml-1 text-primary-800 hover:text-primary-900"
                >
                  ×
                </button>
              </span>
            </div>
            <div class="grid grid-cols-2 gap-2">
              <button
                v-for="tag in availableTags"
                :key="tag"
                @click="toggleTag(tag)"
                :class="[
                  'px-2 py-1 text-xs font-medium rounded-full transition-colors duration-200',
                  document.tags.includes(tag)
                    ? 'bg-primary-100 text-primary-800'
                    : 'bg-gray-700 text-gray-300 hover:bg-gray-600',
                ]"
              >
                {{ tag }}
              </button>
            </div>
          </div>

          <div class="space-y-2 mt-4">
            <label class="block text-sm font-medium text-dark-300"
              >Description</label
            >
            <textarea
              v-model="document.description"
              class="input h-24"
              placeholder="Enter document description"
            ></textarea>
          </div>
        </div>
      </div>

      <div class="lg:col-span-2">
        <div class="card">
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-lg font-medium text-dark-100">Content</h2>
            <div class="flex space-x-2">
              <button
                @click="insertCodeBlock"
                class="btn btn-secondary text-sm"
              >
                Insert Code Block
              </button>
              <button
                @click="previewMode = !previewMode"
                class="btn btn-secondary text-sm"
              >
                {{ previewMode ? "Edit" : "Preview" }}
              </button>
            </div>
          </div>

          <div v-if="!previewMode" class="space-y-4">
            <textarea
              v-model="document.content"
              class="input font-mono h-[400px]"
              placeholder="Write your content here..."
            ></textarea>
          </div>
          <div v-else class="prose prose-invert max-w-none">
            <div v-html="renderedContent"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { marked } from "marked";
import hljs from "highlight.js";
import "highlight.js/styles/github-dark.css";
import { useServiceStore } from "../stores/service";
import { useAuthStore } from "../stores/auth";

const props = defineProps({
  initialDocument: {
    type: Object,
    default: () => ({
      title: "",
      description: "",
      type: "",
      category: "",
      tags: [],
      content: "",
      serviceId: null,
    }),
  },
  isEditing: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["save", "cancel"]);
const serviceStore = useServiceStore();
const authStore = useAuthStore();

const document = ref({ ...props.initialDocument });
const previewMode = ref(false);
const saving = ref(false);
const services = ref([]);

// Ensure document.tags is always an array
if (!document.value.tags) {
  document.value.tags = [];
}

// Fetch services when component is mounted
onMounted(async () => {
  if (!authStore.isAuthenticated) {
    return;
  }

  try {
    await serviceStore.fetchServices();
    services.value = serviceStore.services;

    // If we're creating a new document and a service ID was passed in the URL
    const urlParams = new URLSearchParams(window.location.search);
    const serviceId = urlParams.get("serviceId");
    if (serviceId && !props.isEditing) {
      // Ensure serviceId is a number
      const numericServiceId = parseInt(serviceId, 10);
      if (!isNaN(numericServiceId)) {
        document.value.serviceId = numericServiceId;
        document.value.type = "service";
      }
    }
  } catch (error) {
    console.error("Failed to load services:", error);
  }

  // Initialize with empty document if not editing
  if (!props.isEditing) {
    document.value = {
      title: "",
      description: "",
      type: "",
      category: "",
      tags: [],
      content: "",
      serviceId: document.value.serviceId || null,
    };
  } else {
    // Ensure tags is an array when editing
    if (!document.value.tags) {
      document.value.tags = [];
    }
  }
});

const availableTags = [
  // Service-related tags
  "omnichannel",
  "crmconnector",
  "clickhouse",
  "api",
  "database",
  "cache",
  "queue",
  "monitoring",
  "auth",
  "storage",
  "search",

  // Document-related tags
  "documentation",
  "guide",
  "tutorial",
  "specification",
  "architecture",
  "design",
  "implementation",
];

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
  return marked(document.value.content);
});

const toggleTag = (tag) => {
  if (document.value.tags.includes(tag)) {
    removeTag(tag);
  } else {
    document.value.tags.push(tag);
  }
};

const removeTag = (tag) => {
  document.value.tags = document.value.tags.filter((t) => t !== tag);
};

const insertCodeBlock = () => {
  const codeBlock = "\n```\n// Your code here\n```\n";
  document.value.content += codeBlock;
};

const handleSave = async () => {
  try {
    // Process tags to ensure they are objects with a name property
    const processedTags = document.value.tags.map((tag) => {
      if (typeof tag === "object" && tag.name) {
        return tag;
      } else if (typeof tag === "string") {
        return { name: tag };
      } else {
        return { name: String(tag) };
      }
    });

    // Ensure serviceId is a number
    const serviceId = document.value.serviceId
      ? parseInt(document.value.serviceId, 10)
      : null;

    const documentToSave = {
      ...document.value,
      tags: processedTags,
      serviceId,
    };

    emit("save", documentToSave);
  } catch (error) {
    console.error("Error preparing document for save:", error);
    throw error;
  }
};
</script>

<style>
.prose {
  @apply text-dark-100;
}

.prose pre {
  @apply bg-gray-800 p-4 rounded-lg overflow-x-auto;
}

.prose code {
  @apply bg-gray-800 px-1 py-0.5 rounded text-sm;
}

.prose pre code {
  @apply bg-transparent p-0;
}

.prose h1,
.prose h2,
.prose h3,
.prose h4 {
  @apply text-dark-100 mt-6 mb-4;
}

.prose p {
  @apply text-dark-300 mb-4;
}

.prose ul,
.prose ol {
  @apply text-dark-300 mb-4 pl-4;
}

.prose li {
  @apply mb-2;
}

.prose a {
  @apply text-primary-500 hover:text-primary-400;
}

.prose blockquote {
  @apply border-l-4 border-primary-500 pl-4 italic text-dark-300 my-4;
}
</style>
