import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { fetchApi } from '@/plugins/api'

const DefaultUsername = 'Unknown user'

export const useSessionStore = defineStore('session', () => {
  const username = ref(DefaultUsername)
  const isAdmin = ref(false)
  const isGuest = computed(() => !isAdmin.value)

  const hello = async () => {
    try {
      const r = await fetchApi<{
        username: string,
        role: 'admin' | 'guest'
      }>('/api/auth/hello')

      isAdmin.value = r.role === 'admin'
    } catch (e) {
      username.value = DefaultUsername
      isAdmin.value = false

      console.error('Server hello failed', e)
    }
  }

  return {
    hello,
    username,
    isAdmin,
    isGuest,
  }
})
