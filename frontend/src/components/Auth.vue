<script lang="ts" setup>
import { onDeactivated, onMounted } from 'vue'
import { useStore } from '../store'
import AuthLogin from './AuthLogin.vue'
import AuthJoin from './AuthJoin.vue'

const store = useStore()

onMounted(() => {
  store.setShowAuth(true)
})

onDeactivated(() => {
  store.setShowAuth(false)
})
</script>

<template>
  <div v-if="store.uiState.showAuth" class="flex min-h-screen flex-row justify-center items-center divide-x-2 divide-base04 divide-dashed text-base0">
    <AuthLogin />
    <AuthJoin />
  </div>
</template>
