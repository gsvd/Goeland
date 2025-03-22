<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import AddAccount from './components/AddAccount.vue'
import { HasAnyAccount } from '../wailsjs/go/main/App.js'

const hasAnyAccount = ref(false)

onMounted(async () => {
  hasAnyAccount.value = await HasAnyAccount()
})
</script>

<template>
  <div class="app-wrapper">
    <div v-if="hasAnyAccount">You have some accounts.</div>
    <div v-else>
      <AddAccount/>
    </div>
  </div>
</template>

<style scoped>
.app-wrapper {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
</style>
