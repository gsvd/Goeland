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
  <div class="layout">
    <Sidebar />

    <main class="main-content">
      <Loading v-if="store.isLoading" />

      <Auth />
      <div v-if="!store.getActiveAccount && !store.uiState.showAuth">
        <p class="text-center text-gray-400">Select an account to manage</p>
      </div>
      <div v-else-if="store.getActiveAccount && !store.uiState.showAuth">
        <p class="text-white text-xl text-center">
          Logged in as {{ store.getActiveAccount?.JID }}
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

.main-content {
  flex: 1;
  overflow-y: auto;
  background-color: #36393f;
}
</style>
