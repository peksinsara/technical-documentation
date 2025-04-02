<template>
  <div class="space-y-6">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-dark-100">Use Cases</h1>
      <button class="btn btn-primary">New Use Case</button>
    </div>

    <div class="grid gap-6 md:grid-cols-2">
      <div v-for="useCase in useCases" :key="useCase.id" class="card">
        <div class="flex items-start justify-between">
          <div>
            <h3 class="text-xl font-semibold text-dark-100">
              {{ useCase.title }}
            </h3>
            <p class="text-sm text-dark-400 text-left text-left mt-1">
              {{ useCase.description }}
            </p>
          </div>
          <span
            :class="getPriorityClass(useCase.priority)"
            class="px-2 py-1 text-xs font-medium rounded-full"
          >
            {{ useCase.priority }}
          </span>
        </div>

        <div class="mt-4 space-y-4">
          <div>
            <h4 class="text-sm font-medium text-dark-300">Steps</h4>
            <div class="mt-2 space-y-2">
              <div
                v-for="(step, index) in useCase.steps"
                :key="index"
                class="flex items-start space-x-2"
              >
                <span
                  class="flex-shrink-0 w-6 h-6 flex items-center justify-center rounded-full bg-primary-100 text-primary-800 text-xs"
                >
                  {{ index + 1 }}
                </span>
                <span class="text-sm text-dark-400 text-left text-left">{{
                  step
                }}</span>
              </div>
            </div>
          </div>

          <div>
            <h4 class="text-sm font-medium text-dark-300">Expected Results</h4>
            <ul class="mt-2 space-y-1">
              <li
                v-for="(result, index) in useCase.expectedResults"
                :key="index"
                class="text-sm text-dark-400 text-left text-left flex items-start"
              >
                <span class="text-primary-500 mr-2">•</span>
                {{ result }}
              </li>
            </ul>
          </div>

          <div v-if="useCase.edgeCases.length">
            <h4 class="text-sm font-medium text-dark-300">Edge Cases</h4>
            <div class="mt-2 space-y-2">
              <div
                v-for="(edgeCase, index) in useCase.edgeCases"
                :key="index"
                class="text-sm text-dark-400 text-left text-left"
              >
                <span class="font-medium text-primary-500"
                  >Case {{ index + 1 }}:</span
                >
                {{ edgeCase }}
              </div>
            </div>
          </div>
        </div>

        <div class="mt-4 flex space-x-2">
          <button class="text-sm text-primary-500 hover:text-primary-400">
            Edit
          </button>
          <button class="text-sm text-primary-500 hover:text-primary-400">
            Test
          </button>
          <button class="text-sm text-primary-500 hover:text-primary-400">
            Documentation
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const useCases = ref([
  {
    id: 1,
    title: "User Authentication",
    description:
      "User successfully logs in to the system using valid credentials",
    priority: "High",
    steps: [
      "Navigate to the login page",
      "Enter valid email address",
      "Enter valid password",
      'Click the "Login" button',
    ],
    expectedResults: [
      "User is redirected to the dashboard",
      "Authentication token is stored securely",
      "Last login timestamp is updated",
    ],
    edgeCases: [
      "Invalid email format",
      "Incorrect password",
      "Account locked after multiple failed attempts",
    ],
  },
  {
    id: 2,
    title: "Document Creation",
    description: "User creates a new technical document with markdown support",
    priority: "Medium",
    steps: [
      'Click "New Document" button',
      "Enter document title",
      "Select document category",
      "Write content using markdown editor",
      "Add relevant tags",
      'Click "Save" button',
    ],
    expectedResults: [
      "Document is saved successfully",
      "Document appears in the documents list",
      "Version history is created",
      "Tags are properly associated",
    ],
    edgeCases: [
      "Empty document title",
      "Maximum file size exceeded",
      "Invalid markdown syntax",
    ],
  },
]);

const getPriorityClass = (priority) => {
  switch (priority.toLowerCase()) {
    case "high":
      return "bg-red-100 text-red-800";
    case "medium":
      return "bg-yellow-100 text-yellow-800";
    case "low":
      return "bg-green-100 text-green-800";
    default:
      return "bg-gray-100 text-gray-800";
  }
};
</script>
