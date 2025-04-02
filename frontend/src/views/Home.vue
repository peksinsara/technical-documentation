<template>
  <div>
    <div class="relative bg-dark-800 overflow-hidden">
      <div class="max-w-7xl mx-auto">
        <div
          class="relative z-10 pb-8 sm:pb-16 md:pb-20 lg:max-w-2xl lg:w-full lg:pb-28 xl:pb-32"
        >
          <main
            class="mt-10 mx-auto max-w-7xl px-4 sm:mt-12 sm:px-6 md:mt-16 lg:mt-20 lg:px-8 xl:mt-28"
          >
            <div class="sm:text-center lg:text-left">
              <h1
                class="text-4xl tracking-tight font-extrabold text-dark-100 sm:text-5xl md:text-6xl"
              >
                <span class="block">Technical Documentation</span>
                <span class="block text-primary-500">Bicom Systems</span>
              </h1>
              <p
                class="mt-3 text-base text-dark-400 text-left text-left sm:mt-5 sm:text-lg sm:max-w-xl sm:mx-auto md:mt-5 md:text-xl lg:mx-0"
              >
                All technical documentation for services and products to make
                developers day better
              </p>
              <div
                class="mt-5 sm:mt-8 sm:flex sm:justify-center lg:justify-start"
              >
                <div class="rounded-md shadow">
                  <router-link
                    to="/documents"
                    class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-primary-600 hover:bg-primary-700 md:py-4 md:text-lg md:px-10"
                  >
                    Get Started
                  </router-link>
                </div>
                <div class="mt-3 sm:mt-0 sm:ml-3">
                  <router-link
                    to="/services"
                    class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-primary-700 bg-primary-100 hover:bg-primary-200 md:py-4 md:text-lg md:px-10"
                  >
                    Learn More
                  </router-link>
                </div>
              </div>
            </div>
          </main>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import {
  DocumentTextIcon,
  CodeBracketIcon,
  ChartBarIcon,
  UserGroupIcon,
  CloudArrowUpIcon,
} from "@heroicons/vue/24/outline";

const features = [
  {
    name: "Markdown Support",
    description:
      "Write documentation using Markdown with syntax highlighting and preview.",
    icon: DocumentTextIcon,
  },
  {
    name: "Code Snippets",
    description:
      "Include code snippets with syntax highlighting for multiple languages.",
    icon: CodeBracketIcon,
  },
  {
    name: "Diagrams",
    description: "Create and embed diagrams using Mermaid and PlantUML.",
    icon: ChartBarIcon,
  },
  {
    name: "Collaboration",
    description:
      "Work together with your team in real-time with comments and reviews.",
    icon: UserGroupIcon,
  },
  {
    name: "Version Control",
    description:
      "Track changes and maintain document history with version control.",
    icon: CloudArrowUpIcon,
  },
];
</script>
