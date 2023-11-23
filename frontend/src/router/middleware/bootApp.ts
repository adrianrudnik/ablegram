import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { websocket } from '@/websocket'
import { fetchApi } from '@/plugins/api'
import { useConfigStore } from '@/stores/config'
import { useSessionStore } from '@/stores/session'

export async function bootApp(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext
) {
  console.debug('Booting app')
  console.debug('Booting index: ' + websocket.status.value)

  // Ensure we say hello to the server to identify us
  await useSessionStore().hello()

  // Ensure the config is loaded
  await useConfigStore().load()

  return next()
}
