<script lang="ts" setup>
import { onMounted } from 'vue'
import { useStore } from './store'
import Auth from './components/Auth.vue'
import Loading from './components/Loading.vue'
import Sidebar from './components/Sidebar.vue'

const store = useStore()

onMounted(() => {
  store.getAllAccounts()
})
</script>

<template>
  <div class="flex min-h-screen">
    <Sidebar />

    <main class="flex-1 overflow-y-auto bg-base02">
      <Loading v-if="store.isLoading" />

      <Auth />
      <div v-if="store.getActiveAccount && !store.uiState.showAuth">
        <p class="text-xl text-center text-base0">
          Logged in as {{ store.getActiveAccount?.JID }}
        </p>
      </div>
    </main>
  </div>
</template>

