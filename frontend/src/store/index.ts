import { defineStore } from "pinia";
import { store } from "../../wailsjs/go/models";
import { GetAllAccounts, AddAccount } from "../../wailsjs/go/main/App.js";

export const useStore = defineStore("app", {
  state: () => ({
    accounts: [] as store.Account[],
    isLoading: false,
    error: null as string | null,
    activeAccountId: null as number | null,
    uiState: {
      showAddAccount: false,
    },
  }),

  getters: {
    activeAccount(state): store.Account | null {
      return state.accounts.find(acc => acc.ID === state.activeAccountId) || null;
    },
  },

  actions: {
    async getAllAccounts() {
      this.isLoading = true;
      this.error = null;
      try {
        const raw = await GetAllAccounts();
        const accounts = raw.map((acc: any) => store.Account.createFrom(acc));
        this.accounts = accounts;
        if (accounts.length > 0 && this.activeAccountId === null) {
          this.activeAccountId = accounts[0].ID;
        }
      } catch (err) {
        console.error("Failed to load accounts:", err);
        this.error = "Failed to load accounts";
      } finally {
        this.isLoading = false;
      }
    },

    async login(jid: string, password: string) {
      this.isLoading = true;
      this.error = null;
      try {
        const added = await AddAccount(jid, password);
        const created = store.Account.createFrom(added);
        
        this.accounts.push(created);
        this.setActiveAccount(created.ID);
      } catch (err: any) {
        console.error("Failed to add account:", err);
        this.error = err?.message || "ERR_UNKNOWN";
      } finally {
        this.isLoading = false;
      }
    },

    setActiveAccount(id: number) {
      this.activeAccountId = id;
      this.setShowAddAccount(false);
    },

    setShowAddAccount(open: boolean) {

      this.uiState.showAddAccount = open;
    },
  },
});
