import { defineStore } from 'pinia'
import { computed, ref, nextTick } from 'vue'
import { fetchApi } from '@/plugins/api'
import { useUserStore } from '@/stores/users'
import { useFilesStore } from '@/stores/files'
import { websocket } from '@/websocket'

const DefaultUsername = 'Unknown user'
const DefaultIp = 'Unknown IP'

export const useSessionStore = defineStore('session', () => {
  const id = ref<string | undefined>(undefined)
  const username = ref(DefaultUsername)
  const ip = ref(DefaultIp)
  const isAdmin = ref(false)
  const isGuest = computed(() => !isAdmin.value)

  const { clear: clearUsers } = useUserStore()
  const { clear: clearFiles } = useFilesStore()

  const hello = async () => {
    try {
      const r = await fetchApi<{
        id: string,
        ip: string
        display_name: string,
        role: 'admin' | 'guest'
      }>('/api/auth', {
        method: 'POST',
      })
      username.value = r.display_name
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
      await fetchApi('/api/auth', {
        method: 'DELETE'
      })
    } catch (e) {
      console.error('Server goodbye failed', e)
    }

    // Refresh the session
    await hello()
  }

  const reconsider = async () => {
    websocket.close()

    clearUsers()
    clearFiles()

    await hello()

    websocket.open()
  }

  return {
    hello,
    goodbye,
    reconsider,
    id,
    username,
    ip,
    isAdmin,
    isGuest
  }
})
