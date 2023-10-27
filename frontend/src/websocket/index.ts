import { useWebSocket } from '@vueuse/core'
import { useFilesStore } from '@/stores/files'
import { useStatStore } from '@/stores/stats'
import type { PushMessage } from '@/websocket/messages/global'
import { PushMessageType } from '@/websocket/messages/global'
import { hydrateTags, useTagStore } from '@/stores/tags'

export const websocket = useWebSocket(import.meta.env.VITE_WEBSOCKET_URL, {
  autoReconnect: true,
  async onMessage(ws, event) {
    const payload = JSON.parse(event.data) as PushMessage

    switch (payload.type) {
      case PushMessageType.FileStatus:
        useFilesStore().update(payload)
        break

      case PushMessageType.MetricUpdate:
        useStatStore().update(payload.k, payload.v)
        break

      case PushMessageType.ProcessingStatus:
        // Hydrate the new tags
        await hydrateTags()
        useStatStore().isProcessing = payload.status
        break
    }
  }
})
