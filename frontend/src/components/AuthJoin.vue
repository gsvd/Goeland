<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { useStore } from '../store'

const error = ref<string | null>(null)

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
    <h1 class="text-2xl font-bold mb-12 text-white">Join</h1>
    <form @submit.prevent="join" class="w-full max-w-sm space-y-6">
      <div>
        <div class="flex items-center gap-x-4 mb-3">
          <label class="block text-sm font-medium text-white">JID</label>
        </div>
        <input v-model="form.jid" type="text"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-gray-400" placeholder="gsvd@goeland.im"
          required />
      </div>

      <div>
        <label class="block text-sm font-bold mb-3 text-white">Password</label>
        <input v-model="form.password" type="password"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-gray-400" placeholder="********"
          required />
      </div>

      <div>
        <label class="block text-sm font-bold mb-3 text-white">Confirm Password</label>
        <input v-model="form.confirmPassword" type="password"
          class="w-full px-3 py-2 border rounded-md focus:outline-none border-gray-400" placeholder="********"
          required />
      </div>

      <input type="submit" class="w-full py-2 bg-[#5865F2] rounded-md hover:bg-[#4752C4] transition cursor-pointer"
        value="Join" />
    </form>
    <p v-if="error" class="text-red-500 text-sm mt-4">
      {{ $t(`errors.${error}`) }}
    </p>
  </div>
</template>