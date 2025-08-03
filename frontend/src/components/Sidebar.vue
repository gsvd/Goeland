<script lang="ts" setup>
import { computed } from 'vue'
import { useStore } from '../store'

const store = useStore()
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
  <aside class="w-[80px] flex flex-col justify-between items-center py-3 text-base0 bg-base03">
    <div class="flex flex-col gap-3 items-center">
      <button v-for="account in store.getAccounts" :key="account.ID" class="w-12 h-12 rounded-full bg-base04 flex items-center justify-center cursor-pointer font-bold hover:outline-2 hover:outline-base0" :title="account.JID"
        :class="{ 'outline-2 outline-base0': store.getActiveAccountJID === account.JID && !store.uiState.showAuth }"
        @click="store.setActiveAccount(account)">
        {{ account.JID.charAt(0).toUpperCase() }}
      </button>
    </div>
    <div class="flex flex-col gap-3 items-center">
      <button v-for="item in additionalItems" :key="item.id" class="w-12 h-12 rounded-full bg-base04 flex items-center justify-center cursor-pointer font-bold hover:outline-2 hover:outline-base0" :title="item.title"
      :class="{ 'outline-2 outline-base0': store.uiState.showAuth }"
        @click="item.action">
        <span class="text-base0 text-xl">{{ item.icon }}</span>
      </button>
    </div>
  </aside>
</template>
\