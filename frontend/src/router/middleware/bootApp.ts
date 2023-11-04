import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { websocket } from '@/websocket'
import { useUiStore } from '@/stores/ui'
import { startWebsocketWatchers } from '@/websocket/watchers'

export async function bootApp(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext
) {
  console.debug('Booting app')
  console.debug('Booting index: ' + websocket.status.value)

  // We start the watchers first, so that we can react to history changes
  // that we get force-pushed once we connect to the websocket, i.e.:
  // processRoutineCount will go from -1 to 0, which will trigger the
  // watcher to hydrate the tags.
  startWebsocketWatchers()

  await useUiStore().preload()

  return next()
}
