import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { websocket } from '@/websocket'
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

  return next()
}
