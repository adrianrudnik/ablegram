import { useWebSocket } from '@vueuse/core'
import { useFilesStore } from '@/stores/files'
import { useStatStore } from '@/stores/stats'
import type { PushMessage } from '@/websocket/messages/global'
import { PushMessageType } from '@/websocket/messages/global'
import { hydrateTags } from '@/stores/tags'
import router from '@/router'

export const websocket = useWebSocket(import.meta.env.VITE_WEBSOCKET_URL, {
  autoReconnect: true,
  async onMessage(ws, event) {
    const payload = JSON.parse(event.data) as PushMessage

    switch (payload.type) {
      case PushMessageType.FileStatus:
        useFilesStore().update(payload)
        break

      case PushMessageType.MetricUpdate:
        useStatStore().updateMetrics(payload.values)
        break

      case PushMessageType.ProcessingStatus:
        // Hydrate the new tags on completion and update them in periods while
        // eslint-disable-next-line no-case-declarations
        const stop = setInterval(async () => {
          await hydrateTags()
        }, 1500)

        // Once finished, stop that interval and reload one last time
        if (payload.routines === 0) {
          clearInterval(stop)
          await hydrateTags()
        }

        useStatStore().isProcessing = payload.routines !== 0
        break

      case PushMessageType.ForceNavigate:
        await router.push({ name: payload.target })
        break
    }
  }
})
