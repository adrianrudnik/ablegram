import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { fetchApi } from '@/plugins/api'
import { GuestRole, useUserClientStore } from '@/stores/users'
import { useFilesStore } from '@/stores/files'
import { websocket } from '@/websocket'
import type { UserRoles } from '@/stores/users'

const DefaultUsername = 'Unknown user'
const DefaultIp = 'Unknown IP'

export const useSessionStore = defineStore('session', () => {
  const clientId = ref<string | undefined>(undefined)

  const userId = ref<string | undefined>(undefined)
  const displayName = ref(DefaultUsername)
  const role = ref<'admin' | 'guest'>('guest')
  const ip = ref(DefaultIp)

  const isAdmin = computed(() => role.value === 'admin')
  const isGuest = computed(() => role.value === 'guest')

  const { clear: clearUsers } = useUserClientStore()
  const { clear: clearFiles } = useFilesStore()

  const hello = async () => {
    try {
      const r = await fetchApi<{
        id: string
        display_name: string
        role: UserRoles
      }>('/api/auth', {
        method: 'POST'
      })
      userId.value = r.id
      displayName.value = r.display_name
      role.value = r.role
    } catch (e) {
      displayName.value = DefaultUsername
      role.value = GuestRole

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
    clientId,
    userId,
    displayName,
    role,
    ip,
    isAdmin,
    isGuest
  }
})
