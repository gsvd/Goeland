<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { useStore } from '../store'

const error = ref<string | null>(null)

const form = reactive({
  jid: '',
  password: '',
})

const store = useStore()

async function addAccount() {
  try {
    await store.login(form.jid, form.password)
    store.setShowAddAccount(false)
  } catch (err: any) {
    error.value = err || 'ERR_UNKNOWN'
  }
}

</script>

<template>
  <div class="flex flex-col items-center justify-center min-h-screen">
    <form @submit.prevent="addAccount" class="w-full max-w-sm space-y-4">
      <div>
        <input v-model="form.jid" type="text"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-gray-400" placeholder="gsvd@goeland.im" />
      </div>

      <div>
        <input v-model="form.password" type="password"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-gray-400" placeholder="********" />
      </div>

      <input type="submit" class="w-full py-2 bg-[#5865F2] rounded-md hover:bg-[#4752C4] transition cursor-pointer" value="Connect" />
    </form>
    <p v-if="error" class="text-red-500 text-sm mt-4">
      {{ $t(`errors.${error}`) }}
    </p>
  </div>
</template>
