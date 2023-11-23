import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { fetchApi } from '@/plugins/api'

const DefaultUsername = 'Unknown user'
const DefaultIp = 'Unknown IP'

export const useSessionStore = defineStore('session', () => {
  const username = ref(DefaultUsername)
  const ip = ref(DefaultIp)
  const isAdmin = ref(false)
  const isGuest = computed(() => !isAdmin.value)

  const hello = async () => {
    try {
      const r = await fetchApi<{
        username: string
        ip: string
        role: 'admin' | 'guest'
      }>('/api/auth/hello')
      username.value = r.username
      ip.value = r.ip
      isAdmin.value = r.role === 'admin'
    } catch (e) {
      username.value = DefaultUsername
      ip.value = DefaultIp
      isAdmin.value = false

      console.error('Server hello failed', e)
    }
  }

  const goodbye = async () => {
    try {
      await fetchApi('/api/auth/goodbye', {
        method: 'POST'
      })
    } catch (e) {
      console.error('Server goodbye failed', e)
    }

    // Refresh the session
    await hello()
  }

  return {
    hello,
    goodbye,
    username,
    ip,
    isAdmin,
    isGuest
  }
})
