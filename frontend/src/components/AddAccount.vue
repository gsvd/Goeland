<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { AddAccount } from '../../wailsjs/go/main/App.js'

const error = ref<string | null>(null)
const success = ref<boolean>(false)

const form = reactive({
  jid: '',
  password: '',
})

async function login() {
  try {
    await AddAccount(form.jid, form.password)
    error.value = null
    success.value = true
  } catch (err: any) {
    success.value = false
    error.value = err || 'ERR_UNKNOWN'
  }
}
</script>

<template>
  <div class="flex flex-col items-center justify-center min-h-screen">
    <form @submit.prevent="login" class="w-full max-w-sm space-y-4">
      <div>
        <input v-model="form.jid" type="text"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-gray-400" placeholder="gsvd@goeland.im" />
      </div>

      <div>
        <input v-model="form.password" type="password"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-gray-400" placeholder="********" />
      </div>

      <input type="submit" class="w-full py-2 bg-[#5865F2] rounded-md hover:bg-[#4752C4] transition cursor-pointer" />
    </form>
    <p v-if="error && !success" class="text-red-500 text-sm mt-4">
      {{ $t(`errors.${error}`) }}
    </p>
    <p v-if="success" class="text-green-500 text-sm mt-4">
      {{ $t('account.added_success') }}
    </p>
  </div>
</template>
