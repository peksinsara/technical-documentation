export const config = {
  api: {
    baseUrl: import.meta.env.VITE_API_URL || "http://localhost:8080",
  },
  app: {
    name: import.meta.env.VITE_APP_NAME || "Technical Documentation",
    description:
      import.meta.env.VITE_APP_DESCRIPTION || "Technical Documentation System",
  },
  auth: {
    tokenKey: "auth_token",
    userKey: "auth_user",
  },
  database: {
    host: import.meta.env.VITE_DB_HOST || "localhost",
    port: parseInt(import.meta.env.VITE_DB_PORT) || 3306,
    name: import.meta.env.VITE_DB_NAME || "technical_docs",
    user: import.meta.env.VITE_DB_USER || "root",
    password: import.meta.env.VITE_DB_PASSWORD || "",
  },
};
