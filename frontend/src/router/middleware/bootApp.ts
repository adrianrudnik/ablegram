import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { websocket } from '@/plugins/websocket'
import { useUiStore } from '@/stores/ui'

export async function bootApp(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext
) {
  console.debug('Booting app')
  console.debug('Booting websocket: ' + websocket.status.value)

  await useUiStore().preload()

  return next()
}
