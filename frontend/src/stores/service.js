import { defineStore } from "pinia";
import { config } from "../config";

export const useServiceStore = defineStore("service", {
  state: () => ({
    services: [
      { ID: "crmconnector", name: "Crm Connector" },
      { ID: "asterisk", name: "Asterisk" },
      { ID: "omnichannel", name: "Omnichannel" },
      { ID: "clickhouse", name: "Clickhouse" },
    ],
    loading: false,
    error: null,
  }),

  actions: {
    // This function now just returns the hard-coded services
    async fetchServices() {
      this.loading = true;
      try {
        return this.services;
      } catch (error) {
        this.error = error.message;
        throw error;
      } finally {
        this.loading = false;
      }
    },
  },
});
