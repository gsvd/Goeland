<script lang="ts" setup>
import { reactive } from 'vue'
import { useStore } from '../store'
import { useErrorHandler } from '../composables/useErrorHandler'

const form = reactive({
  address: '',
  password: '',
})

const store = useStore()
const { errorCode, fieldErrors, handleError, clearErrors } = useErrorHandler()

async function login() {
  store.loading = true
  clearErrors()
  try {
    await store.login(form.address, form.password)
  } catch (err: any) {
    handleError(err)
  } finally {
    store.loading = false
  }
}
</script>

<template>
  <div class="flex flex-1 flex-col items-center justify-center py-16">
    <h1 class="text-2xl font-medium mb-12">Login</h1>

    <form @submit.prevent="login" class="w-full max-w-sm space-y-6">
      <div>
        <label class="block text-sm font-medium mb-2">XMPP address</label>
        <input
          v-model="form.address"
          type="text"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-ardoise"
          placeholder="gsvd@goeland.im"
        />
        <p v-if="fieldErrors.address" class="text-red text-sm mt-1">
          {{ $t(`errors.${fieldErrors.address}`) }}
        </p>
      </div>

      <div>
        <label class="block text-sm font-medium mb-2">Password</label>
        <input
          v-model="form.password"
          type="password"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-ardoise"
          placeholder="********"
        />
        <p v-if="fieldErrors.password" class="text-red text-sm mt-1">
          {{ $t(`errors.${fieldErrors.password}`) }}
        </p>
      </div>

      <input
        type="submit"
        class="w-full py-2 rounded-md bg-base03 hover:bg-base04 transition cursor-pointer"
        :disabled="store.loading"
        value="Login"
      />
    </form>

    <p v-if="errorCode && !Object.keys(fieldErrors).length" class="text-red text-sm mt-4">
      {{ $t(`errors.${errorCode}`) }}
    </p>
  </div>
</template>
