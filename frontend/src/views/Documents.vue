<template>
  <div class="w-full">
    <div class="space-y-6">
      <div class="flex justify-between items-center">
        <h1 class="text-3xl font-bold text-dark-100 text-left">Documents</h1>
        <button class="btn btn-primary" @click="showEditor = true">
          New Document
        </button>
      </div>

      <DocumentEditor
        v-if="showEditor"
        :initial-document="editingDocument"
        :is-editing="!!editingDocument"
        @save="handleSave"
        @cancel="handleCancel"
      />

      <DocumentViewer
        v-else-if="viewingDocument"
        :document="viewingDocument"
        @close="handleCloseViewer"
        @edit="handleEdit(viewingDocument)"
      />

      <div v-else class="space-y-6">
        <div
          class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4"
        >
          <div class="flex-1">
            <input
              type="text"
              placeholder="Search documents..."
              class="input w-full text-left"
              v-model="searchQuery"
            />
          </div>
          <select class="input w-full sm:w-48 text-left" v-model="selectedTag">
            <option value="">All Tags</option>
            <option v-for="tag in availableTags" :key="tag" :value="tag">
              {{ tag }}
            </option>
          </select>
        </div>

        <div class="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          <div
            v-for="doc in filteredDocuments"
            :key="doc.id"
            class="card hover:bg-dark-700 transition-colors duration-200 flex flex-col"
          >
            <div class="flex-1">
              <div class="flex justify-between items-start">
                <h3 class="text-lg font-medium text-dark-100 text-left">
                  {{ doc.title }}
                </h3>
              </div>
              <p class="mt-2 text-dark-400 text-left text-line-clamp-2">
                {{ doc.description }}
              </p>

              <div class="mt-4 flex flex-wrap gap-2">
                <span
                  v-for="tag in doc.tags"
                  :key="tag"
                  class="px-2 py-1 text-xs font-medium rounded-full bg-primary-100 text-primary-800 text-left"
                >
                  {{ tag }}
                </span>
              </div>

              <div class="mt-4 flex items-center justify-between">
                <div class="flex items-center space-x-2">
                  <img
                    :src="doc.author.avatar"
                    :alt="doc.author.name"
                    class="h-6 w-6 rounded-full"
                  />
                  <span
                    class="text-sm text-dark-400 text-left text-left text-left"
                  >
                    {{ doc.author.name }}
                  </span>
                </div>
                <span
                  class="text-sm text-dark-400 text-left text-left text-left"
                >
                  {{ doc.lastUpdated }}
                </span>
              </div>
            </div>

            <div class="mt-4 flex space-x-2 pt-4 border-t border-dark-700">
              <button
                class="text-sm text-primary-500 hover:text-primary-400 text-left"
                @click="handleEdit(doc)"
              >
                Edit
              </button>
              <button
                class="text-sm text-primary-500 hover:text-primary-400 text-left"
                @click="handleView(doc)"
              >
                View
              </button>
              <button
                class="text-sm text-red-500 hover:text-red-400 text-left"
                @click="handleDelete(doc)"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import DocumentEditor from "../components/DocumentEditor.vue";
import DocumentViewer from "../components/DocumentViewer.vue";

const searchQuery = ref("");
const selectedTag = ref("");
const showEditor = ref(false);
const editingDocument = ref(null);
const viewingDocument = ref(null);

const availableTags = [
  "api",
  "database",
  "cache",
  "queue",
  "monitoring",
  "auth",
  "storage",
  "search",
  "documentation",
  "guide",
  "tutorial",
];

const documents = ref([
  {
    id: 1,
    title: "API Documentation",
    description:
      "Comprehensive guide to our REST API endpoints and authentication",
    type: "document",
    category: "Documentation",
    tags: ["api", "documentation", "auth"],
    content:
      "# API Documentation\n\nThis is a comprehensive guide to our REST API endpoints and authentication.\n\n## Authentication\n\n```javascript\nconst token = await api.authenticate({\n  username: 'user',\n  password: 'pass'\n});\n```\n\n## Endpoints\n\n### Users\n\n- `GET /api/users` - List all users\n- `POST /api/users` - Create a new user\n- `GET /api/users/:id` - Get user details",
    author: {
      name: "Sara Peksin",
      avatar: "https://ui-avatars.com/api/?name=John+Doe",
    },
    lastUpdated: "2024-03-15",
  },
  {
    id: 2,
    title: "Database Schema",
    description:
      "Detailed documentation of our database structure and relationships",
    type: "diagram",
    category: "Technical",
    tags: ["database", "documentation", "er-diagram"],
    content:
      "# Database Schema\n\n## Tables\n\n### Users\n\n```sql\nCREATE TABLE users (\n  id INT PRIMARY KEY AUTO_INCREMENT,\n  username VARCHAR(255) NOT NULL,\n  email VARCHAR(255) NOT NULL UNIQUE,\n  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP\n);\n```\n\n### Documents\n\n```sql\nCREATE TABLE documents (\n  id INT PRIMARY KEY AUTO_INCREMENT,\n  title VARCHAR(255) NOT NULL,\n  content TEXT,\n  user_id INT,\n  FOREIGN KEY (user_id) REFERENCES users(id)\n);\n```",
    author: {
      name: "Sara Peksin",
      avatar: "https://ui-avatars.com/api/?name=Jane+Smith",
    },
    lastUpdated: "2024-03-14",
  },
  {
    id: 3,
    title: "Omnichannel Service",
    description:
      "Service for handling customer interactions across multiple channels",
    type: "service",
    category: "Architecture",
    tags: ["omnichannel", "api", "service"],
    content:
      "# Omnichannel Service\n\n## Overview\n\nThe Omnichannel Service handles customer interactions across multiple channels including:\n\n- Web chat\n- Email\n- SMS\n- Social media\n\n## Architecture\n\n```mermaid\ngraph TD\n    A[Client] --> B[API Gateway]\n    B --> C[Omnichannel Service]\n    C --> D[Message Queue]\n    D --> E[Channel Handlers]\n    E --> F[External APIs]\n```\n\n## Configuration\n\n```yaml\nchannels:\n  web:\n    enabled: true\n    max_concurrent: 100\n  email:\n    enabled: true\n    smtp_server: smtp.example.com\n  sms:\n    enabled: true\n    provider: twilio\n```",
    author: {
      name: "Sara Peksin",
      avatar: "https://ui-avatars.com/api/?name=Mike+Johnson",
    },
    lastUpdated: "2024-03-13",
  },
]);

const filteredDocuments = computed(() => {
  return documents.value.filter((doc) => {
    const matchesSearch =
      doc.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      doc.description.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchesTag =
      !selectedTag.value || (doc.tags && doc.tags.includes(selectedTag.value));
    return matchesSearch && matchesTag;
  });
});

const handleSave = async (document) => {
  // Ensure tags is an array
  if (!document.tags) {
    document.tags = [];
  }

  if (editingDocument.value) {
    // Update existing document
    const index = documents.value.findIndex(
      (d) => d.id === editingDocument.value.id
    );
    if (index !== -1) {
      documents.value[index] = {
        ...document,
        id: editingDocument.value.id,
        author: editingDocument.value.author,
        lastUpdated: new Date().toISOString().split("T")[0],
      };
    }
  } else {
    // Create new document
    documents.value.push({
      ...document,
      id: documents.value.length + 1,
      author: {
        name: "Current User",
        avatar: "https://ui-avatars.com/api/?name=Current+User",
      },
      lastUpdated: new Date().toISOString().split("T")[0],
    });
  }
  showEditor.value = false;
  editingDocument.value = null;
};

const handleCancel = () => {
  showEditor.value = false;
  editingDocument.value = null;
};

const handleEdit = (doc) => {
  editingDocument.value = { ...doc };
  showEditor.value = true;
};

const handleView = (doc) => {
  viewingDocument.value = { ...doc };
};

const handleCloseViewer = () => {
  viewingDocument.value = null;
};

const handleDelete = (doc) => {
  if (confirm("Are you sure you want to delete this document?")) {
    documents.value = documents.value.filter((d) => d.id !== doc.id);
  }
};
</script>
