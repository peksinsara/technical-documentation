import axios from "axios";
import { config } from "../config";

const api = axios.create({
  baseURL: config.api.baseUrl,
  headers: {
    "Content-Type": "application/json",
  },
});

// Add a request interceptor to add the auth token to requests
api.interceptors.request.use(
  (requestConfig) => {
    const token = localStorage.getItem(config.auth.tokenKey);
    if (token) {
      requestConfig.headers.Authorization = `Bearer ${token}`;
    }
    return requestConfig;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Add a response interceptor to handle errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Handle unauthorized access
      localStorage.removeItem(config.auth.tokenKey);
      localStorage.removeItem(config.auth.userKey);
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

export default api;
