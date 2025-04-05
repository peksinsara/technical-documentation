<template>
  <div class="container mx-auto px-4 py-8 max-w-7xl text-left">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-dark-100">Documents</h1>
    </div>

    <!-- Search and Filter -->
    <div class="mb-6">
      <div class="flex items-center space-x-4">
        <div class="relative flex-1">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search documents..."
            class="input pl-10 w-full"
          />
          <span
            class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>
          </span>
        </div>
      </div>
    </div>

    <!-- Document Editor Modal -->
    <div
      v-if="showEditor"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
    >
      <div
        class="bg-dark-800 rounded-lg shadow-xl max-w-4xl w-full max-h-[90vh] overflow-y-auto"
      >
        <div class="px-6 py-4 border-b border-dark-700">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-medium text-dark-100">
              {{ editingDocument ? "Edit Document" : "New Document" }}
            </h3>
            <button
              @click="showEditor = false"
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
            :initial-document="editingDocument"
            :is-editing="!!editingDocument"
            @save="handleSave"
            @cancel="showEditor = false"
          />
        </div>
      </div>
    </div>

    <!-- Document Viewer Modal -->
    <div
      v-if="showViewer"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50"
    >
      <div
        class="bg-dark-800 rounded-lg shadow-xl max-w-4xl w-full max-h-[90vh] overflow-y-auto"
      >
        <div class="px-6 py-4 border-b border-dark-700">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-medium text-dark-100">
              {{ viewingDocument?.title }}
            </h3>
            <button
              @click="showViewer = false"
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
          <DocumentViewer
            :document="viewingDocument"
            @close="showViewer = false"
            @edit="editDocument(viewingDocument)"
          />
        </div>
      </div>
    </div>

    <!-- Document List -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="doc in filteredDocuments"
        :key="doc.ID"
        class="card hover:shadow-lg transition-shadow duration-200 cursor-pointer flex flex-col"
        @click="viewDocument(doc)"
      >
        <div class="flex-1">
          <div class="flex justify-between items-start mb-4">
            <h2
              class="text-xl font-semibold cursor-pointer hover:text-primary-500 transition-colors"
            >
              {{ doc.title }}
            </h2>
            <span class="px-2 py-1 text-sm">
              {{ doc.type }}
            </span>
          </div>

          <p class="text-gray-600 mb-4">{{ doc.description }}</p>

          <div class="mb-4">
            <span class="text-sm font-medium text-gray-500">Category:</span>
            <span class="ml-2 text-sm">{{ doc.category }}</span>
          </div>

          <div v-if="doc.serviceId" class="mb-4">
            <span class="text-sm font-medium text-gray-500">Service:</span>
            <span class="ml-2 text-sm">{{
              getServiceName(doc.serviceId)
            }}</span>
          </div>

          <div class="flex flex-wrap gap-2 mb-4">
            <span
              v-for="(tag, index) in doc.tags"
              :key="tag.name"
              class="px-3 py-1 text-xs text-white rounded-full"
              :style="{
                backgroundImage:
                  index % 4 === 0
                    ? 'linear-gradient(to right, #3B82F6, #2563EB)'
                    : index % 4 === 1
                    ? 'linear-gradient(to right, #8B5CF6, #7C3AED)'
                    : index % 4 === 2
                    ? 'linear-gradient(to right, #10B981, #059669)'
                    : 'linear-gradient(to right, #EC4899, #DB2777)',
              }"
            >
              {{ tag.name }}
            </span>
          </div>
        </div>

        <div class="flex justify-end items-center space-x-4 mt-auto pt-4">
          <button
            @click.stop="openEditModal(doc)"
            class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors"
          >
            Edit
          </button>
          <button
            @click.stop="handleDelete(doc)"
            class="px-2 py-2 bg-red-500 text-white rounded hover:bg-red-700 transition-colors"
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useDocumentStore } from "../stores/document";
import { useServiceStore } from "../stores/service";
import { useAuthStore } from "../stores/auth";
import DocumentEditor from "../components/DocumentEditor.vue";
import DocumentViewer from "../components/DocumentViewer.vue";

const documentStore = useDocumentStore();
const serviceStore = useServiceStore();
const authStore = useAuthStore();
const documents = ref([]);
const showEditor = ref(false);
const editingDocument = ref(null);
const showViewer = ref(false);
const viewingDocument = ref(null);
const searchQuery = ref("");

onMounted(async () => {
  if (!authStore.isAuthenticated) {
    return;
  }

  try {
    await Promise.all([
      documentStore.fetchDocuments(),
      serviceStore.fetchServices(),
    ]);
    documents.value = documentStore.documents;
  } catch (error) {
    console.error("Failed to fetch data:", error);
  }
});

// Filter documents by search query
const filteredDocuments = computed(() => {
  return documents.value.filter((doc) => {
    const matchesSearch =
      doc.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      doc.description.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      doc.category.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      doc.tags.some((tag) =>
        tag.toLowerCase().includes(searchQuery.value.toLowerCase())
      );

    return matchesSearch;
  });
});

const getServiceName = (serviceId) => {
  const service = serviceStore.services.find((s) => s.ID === serviceId);
  return service ? service.name : "Unknown Service";
};

const createNewDocument = () => {
  editingDocument.value = null;
  showEditor.value = true;
};

const viewDocument = (doc) => {
  if (!doc.ID) {
    return;
  }

  // Create a link element and click it to open in a new tab
  const link = document.createElement("a");
  link.href = `/document/${doc.ID}`;
  link.target = "_blank";
  link.rel = "noopener noreferrer";
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

const editDocument = (doc) => {
  editingDocument.value = { ...doc };
  showEditor.value = true;
  showViewer.value = false;
};

const deleteDocument = async (id) => {
  if (confirm("Are you sure you want to delete this document?")) {
    try {
      await documentStore.deleteDocument(id);
      documents.value = documentStore.documents;
    } catch (error) {
      console.error("Failed to delete document:", error);
    }
  }
};

const handleSave = async (document) => {
  try {
    if (editingDocument.value) {
      await documentStore.updateDocument({
        ...document,
        id: editingDocument.value.id,
      });
    } else {
      await documentStore.createDocument(document);
    }
    documents.value = documentStore.documents;
    showEditor.value = false;
    editingDocument.value = null;
  } catch (error) {
    console.error("Failed to save document:", error);
  }
};

const handleDelete = async (doc) => {
  if (confirm(`Are you sure you want to delete "${doc.title}"?`)) {
    try {
      await documentStore.deleteDocument(doc.ID);
      documents.value = documents.value.filter((d) => d.ID !== doc.ID);
    } catch (error) {
      alert("Failed to delete document. Please try again.");
    }
  }
};
</script>
