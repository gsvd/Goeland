import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    isLoading: false,
    error: null as string | null
  }),

  actions: {
    setLoading(loading: boolean) {
      this.isLoading = loading
    },
    setError(error: string | null) {
      this.error = error
    }
  },

  getters: {
    getLoadingState: (state) => state.isLoading,
    getError: (state) => state.error
  }
}) 