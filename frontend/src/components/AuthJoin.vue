<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { useStore } from '../store'

const form = reactive({
  jid: '',
  password: '',
  confirmPassword: '',
})

const store = useStore()

async function join() {}
</script>

<template>
  <div class="flex flex-1 flex-col items-center justify-center py-16">
    <h1 class="text-2xl font-medium mb-12">Join</h1>
    <form @submit.prevent="join" class="w-full max-w-sm space-y-6">
      <div>
        <div class="flex items-center gap-x-4 mb-3">
          <label class="block text-sm font-medium">JID</label>
        </div>
        <input v-model="form.jid" type="text"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-ardoise" placeholder="gsvd@goeland.im"
          required />
      </div>

      <div>
        <label class="block text-sm font-medium mb-3">Password</label>
        <input v-model="form.password" type="password"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-ardoise" placeholder="********"
          required />
      </div>

      <div>
        <label class="block text-sm font-medium mb-3">Confirm Password</label>
        <input v-model="form.confirmPassword" type="password"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-ardoise" placeholder="********"
          required />
      </div>

      <input type="submit" class="w-full py-2 rounded-md bg-base03 hover:bg-base04 transition cursor-pointer"
        value="Join" />
    </form>
    <p v-if="store.getError" class="text-red-500 text-sm mt-4">
      {{ $t(`errors.${store.getError}`) }}
    </p>
  </div>
</template>