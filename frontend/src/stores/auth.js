import { defineStore } from "pinia";
import { config } from "../config";
import api from "../services/api";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: JSON.parse(localStorage.getItem(config.auth.userKey)) || null,
    token: localStorage.getItem(config.auth.tokenKey) || null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
  },

  actions: {
    async login(credentials) {
      try {
        const response = await api.post("/api/auth/login", credentials);
        const { token, user } = response.data;

        this.token = token;
        this.user = user;

        localStorage.setItem(config.auth.tokenKey, token);
        localStorage.setItem(config.auth.userKey, JSON.stringify(user));

        return true;
      } catch (error) {
        console.error("Login error:", error);
        throw error;
      }
    },

    async register(userData) {
      try {
        const response = await api.post("/api/auth/register", userData);
        const { token, user } = response.data;

        this.token = token;
        this.user = user;

        localStorage.setItem(config.auth.tokenKey, token);
        localStorage.setItem(config.auth.userKey, JSON.stringify(user));

        return true;
      } catch (error) {
        console.error("Registration error:", error);
        throw error;
      }
    },

    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem(config.auth.tokenKey);
      localStorage.removeItem(config.auth.userKey);
    },
  },
});
