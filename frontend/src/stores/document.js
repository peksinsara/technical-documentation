import { defineStore } from "pinia";
import axios from "axios";
import { config } from "../config";

export const useDocumentStore = defineStore("document", {
  state: () => ({
    documents: [],
    loading: false,
    error: null,
  }),

  actions: {
    async fetchDocuments() {
      this.loading = true;
      try {
        const response = await axios.get(
          `${config.api.baseUrl}/api/documents`,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem(
                config.auth.tokenKey
              )}`,
            },
          }
        );
        if (response.data.length > 0) {
        }
        this.documents = response.data;
      } catch (error) {
        this.error = error.response?.data?.error || "Failed to fetch documents";
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async getDocumentById(id) {
      this.loading = true;
      try {
        if (!id) {
          throw new Error("No document ID provided");
        }

        // Ensure id is a number
        const numericId = parseInt(id, 10);

        if (isNaN(numericId)) {
          throw new Error("Invalid document ID");
        }

        const url = `${config.api.baseUrl}/api/documents/${numericId}`;

        const response = await axios.get(url, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem(
              config.auth.tokenKey
            )}`,
          },
        });
        return response.data;
      } catch (error) {
        this.error = error.response?.data?.error || "Failed to fetch document";
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async createDocument(document) {
      this.loading = true;
      try {
        const formattedDocument = {
          ...document,
          tags: document.tags.map((tag) => {
            if (typeof tag === "object" && tag.name) {
              return tag;
            } else if (typeof tag === "string") {
              return { name: tag };
            } else {
              console.warn("Unexpected tag format:", tag);
              return { name: String(tag) };
            }
          }),
          // Ensure serviceId is a number or null
          serviceId: document.serviceId
            ? parseInt(document.serviceId, 10)
            : null,
        };

        if (
          formattedDocument.serviceId !== null &&
          isNaN(formattedDocument.serviceId)
        ) {
          console.error(
            "Invalid service ID format:",
            formattedDocument.serviceId
          );
          throw new Error("Invalid service ID format");
        }

        const response = await axios.post(
          `${config.api.baseUrl}/api/documents`,
          formattedDocument,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem(
                config.auth.tokenKey
              )}`,
            },
          }
        );
        this.documents.push(response.data);
        return response.data;
      } catch (error) {
        this.error = error.response?.data?.error || "Failed to create document";
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async updateDocument(document) {
      this.loading = true;
      try {
        // Format tags properly for the backend - ensure they are objects with a name property
        const formattedDocument = {
          ...document,
          tags: document.tags.map((tag) => {
            if (typeof tag === "object" && tag.name) {
              return tag;
            } else if (typeof tag === "string") {
              return { name: tag };
            } else {
              return { name: String(tag) };
            }
          }),
          serviceId: document.serviceId
            ? parseInt(document.serviceId, 10)
            : null,
        };

        if (
          formattedDocument.serviceId !== null &&
          isNaN(formattedDocument.serviceId)
        ) {
          console.error(
            "Invalid service ID format for update:",
            formattedDocument.serviceId
          );
          throw new Error("Invalid service ID format");
        }

        const response = await axios.put(
          `${config.api.baseUrl}/api/documents/${document.id}`,
          formattedDocument,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem(
                config.auth.tokenKey
              )}`,
            },
          }
        );
        const index = this.documents.findIndex((d) => d.id === document.id);
        if (index !== -1) {
          this.documents[index] = response.data;
        }
        return response.data;
      } catch (error) {
        console.error("Error updating document:", error);
        this.error = error.response?.data?.error || "Failed to update document";
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async deleteDocument(id) {
      this.loading = true;
      try {
        await axios.delete(`${config.api.baseUrl}/api/documents/${id}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem(
              config.auth.tokenKey
            )}`,
          },
        });
        this.documents = this.documents.filter((d) => d.id !== id);
      } catch (error) {
        this.error = error.response?.data?.error || "Failed to delete document";
        throw error;
      } finally {
        this.loading = false;
      }
    },
  },
});
