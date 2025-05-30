import { defineStore } from 'pinia'

export const useStore = defineStore('main', {
  state: () => ({
    isLoading: false
  }),

  getters: {
    loading: (state) => state.isLoading
  }
}) 