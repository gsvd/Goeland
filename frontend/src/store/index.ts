import { defineStore } from "pinia";
import { store } from "../../wailsjs/go/models";
import { GetAllAccounts, AddAccount } from "../../wailsjs/go/main/App.js";
import { get } from "http";
import { resolve } from "path";

export const useStore = defineStore("app", {
  state: () => ({
    accounts: [] as store.Account[],
    loading: false,
    error: null as string | null,
    activeAccountId: null as number | null,
    uiState: {
      showAuth: false,
    },
  }),

  getters: {
    getActiveAccount(state): store.Account | null {
      return state.accounts.find(acc => acc.ID === state.activeAccountId) || null;
    },
    getAccounts(state): store.Account[] {
      return state.accounts;
    },
    isLoading(state): boolean {
      return state.loading;
    },
    getError(state): string | null {
      return state.error;
    }
  },

  actions: {
    async getAllAccounts(timeout: number = 2500) {
      this.loading = true;
      try {
        const raw = await GetAllAccounts();
        if (!Array.isArray(raw)) {
          throw new Error("Invalid response from GetAllAccounts");
        }

        const accounts = raw.map((acc: any) => store.Account.createFrom(acc));
        this.accounts = accounts;
        if (accounts.length > 0 && this.activeAccountId === null) {
          this.activeAccountId = accounts[0].ID;
        }
      } catch (err) {
        console.error("Failed to load accounts:", err);
      } finally {
        setTimeout(() => {
          this.loading = false;
        }, timeout);
      }
    },

    async login(jid: string, password: string) {
      const added = await AddAccount(jid, password);
      const created = store.Account.createFrom(added);
      
      this.accounts.push(created);
      this.setActiveAccount(created.ID);
    },

    setActiveAccount(id: number) {
      this.activeAccountId = id;
      this.setShowAuth(false);
    },

    setShowAuth(open: boolean) {
      this.uiState.showAuth = open;
    },

    setError(error: string | null) {
      this.error = error;
    },
  },
});
