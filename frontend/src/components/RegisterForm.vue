<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-900">
    <div class="max-w-md w-full space-y-8 p-8 bg-gray-800 rounded-lg shadow-lg">
      <div>
        <h2 class="text-center text-3xl font-extrabold text-gray-100">
          Create your account
        </h2>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="handleSubmit">
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="username" class="sr-only">Username</label>
            <input
              id="username"
              v-model="username"
              name="username"
              type="text"
              required
              class="input rounded-t-md"
              placeholder="Username"
            />
          </div>
          <div>
            <label for="email" class="sr-only">Email</label>
            <input
              id="email"
              v-model="email"
              name="email"
              type="email"
              required
              class="input"
              placeholder="Email"
            />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input
              id="password"
              v-model="password"
              name="password"
              type="password"
              required
              class="input"
              placeholder="Password (minimum 6 characters)"
            />
          </div>
          <div>
            <label for="confirmPassword" class="sr-only"
              >Confirm Password</label
            >
            <input
              id="confirmPassword"
              v-model="confirmPassword"
              name="confirmPassword"
              type="password"
              required
              class="input rounded-b-md"
              placeholder="Confirm Password"
            />
          </div>
        </div>

        <div v-if="passwordError" class="text-red-500 text-sm text-center">
          {{ passwordError }}
        </div>

        <div>
          <button
            type="submit"
            class="btn btn-primary w-full"
            :disabled="loading"
          >
            {{ loading ? "Creating account..." : "Create account" }}
          </button>
        </div>

        <div v-if="error" class="text-red-500 text-sm text-center">
          {{ error }}
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/auth";

const router = useRouter();
const authStore = useAuthStore();

const username = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");
const loading = ref(false);
const error = ref("");

const passwordError = computed(() => {
  if (password.value.length < 6) {
    return "Password must be at least 6 characters long";
  }
  if (password.value !== confirmPassword.value) {
    return "Passwords do not match";
  }
  return "";
});

const handleSubmit = async () => {
  if (passwordError.value) {
    error.value = passwordError.value;
    return;
  }

  try {
    loading.value = true;
    error.value = "";
    const userData = {
      username: username.value,
      email: email.value,
      password: password.value,
    };
    await authStore.register(userData);
    router.push("/documents");
  } catch (err) {
    error.value = err.response?.data?.error || "Failed to create account";
  } finally {
    loading.value = false;
  }
};
</script>
