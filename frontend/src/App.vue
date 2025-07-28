<script lang="ts" setup>
import { computed, onMounted } from 'vue'
import { useStore } from './store'
import AddAccount from './components/AddAccount.vue'
import Loading from './components/Loading.vue'

const store = useStore()
const activeId = computed(() => store.activeAccountId)

onMounted(() => {
  store.getAllAccounts()
})
</script>

<template>
  <div class="layout">
    <aside class="sidebar">
      <div class="account-list">
        <div
          v-for="account in store.accounts"
          :key="account.ID"
          class="account-item"
          :title="account.JID"
          :class="{ active: activeId === account.ID && !store.uiState.showAddAccount }"
          @click="store.setActiveAccount(account.ID)"
        >
          {{ account.JID.charAt(0).toUpperCase() }}
        </div>
        <div
          class="account-item"
          @click="store.setShowAddAccount(true)"
          :class="{ active: store.uiState.showAddAccount }"
          title="Add Account"
        >
          <span class="text-white">+</span>
        </div>
      </div>
    </aside>

    <main class="main-content">
      <Loading v-if="store.isLoading" />

      <div v-else-if="store.accounts.length === 0 || store.uiState.showAddAccount">
        <AddAccount />
      </div>

      <div v-else-if="!store.activeAccount">
        <p class="text-center text-gray-400">Select an account to manage</p>
      </div>

      <div v-else>
        <p class="text-white text-xl text-center">
          Logged in as {{ store.activeAccount.JID }}
        </p>
      </div>
    </main>
  </div>
</template>

<style scoped>
.layout {
  display: flex;
  height: 100vh;
  background-color: #2f3136;
  color: white;
  font-family: sans-serif;
}

.sidebar {
  width: 80px;
  background-color: #202225;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 12px 0;
}

.account-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 100%;
  align-items: center;
}

.account-item {
  width: 48px;
  height: 48px;
  background-color: #36393f;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.2s;
}

.account-item:hover {
  background-color: #5865f2;
}

.account-item.active {
  border: 2px solid #5865f2;
  background-color: #444;
}

.account-item.active:hover {
  background-color: #5865f2;
}

.main-content {
  flex: 1;
  overflow-y: auto;
  background-color: #36393f;
}
</style>
