<script lang="ts" setup>
import { computed } from 'vue'
import { useStore } from '../store'

const store = useStore()
const activeId = computed(() => store.activeAccountId)
const additionalItems = [
  {
    id: 'add-account',
    title: 'Add Account',
    icon: '+',
    action: () => store.setShowAuth(true),
  },
]
</script>

<template v-if="store.getAccounts.length === 0 || store.uiState.showAuth">
  <aside class="sidebar">
    <div class="account-list">
      <button v-for="account in store.getAccounts" :key="account.ID" class="account-item" :title="account.JID"
        :class="{ active: activeId === account.ID && !store.uiState.showAuth }"
        @click="store.setActiveAccount(account.ID)">
        {{ account.JID.charAt(0).toUpperCase() }}
      </button>
      <button v-for="item in additionalItems" :key="item.id" class="account-item" :title="item.title"
      :class="{ active: store.uiState.showAuth }"
        @click="item.action">
        <span class="text-white">{{ item.icon }}</span>
      </button>
    </div>
  </aside>
</template>

<style scoped>
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
</style>