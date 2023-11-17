import { useWebSocket } from '@vueuse/core'
import { useFilesStore } from '@/stores/files'
import { useStatStore } from '@/stores/stats'
import type { PushMessage } from '@/websocket/messages/global'
import { PushMessageType } from '@/websocket/messages/global'
import router from '@/router'
import { createTagFromString, useTagStore } from '@/stores/tags'

export function getWebsocketUrl() {
  if (import.meta.env.VITE_WEBSOCKET_URL ?? null) {
    return import.meta.env.VITE_WEBSOCKET_URL
  }

  const loc = window.location
  const protocol = loc.protocol === 'https:' ? 'wss://' : 'ws://'
  const host = loc.host
  const path = '/ws'
  return protocol + host + path
}

export const websocket = useWebSocket(getWebsocketUrl(), {
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

      case PushMessageType.TagUpdate:
        useTagStore().updateBatch(
          Object.entries(payload.tags)
            .map((v) => createTagFromString(v[0], v[1]))
            .filter((v) => v !== null)
        )
        break

      case PushMessageType.ProcessingStatus:
        useStatStore().processRoutineCount = payload.routines
        useStatStore().isProcessing = payload.routines !== 0
        break

      case PushMessageType.ForceNavigate:
        await router.push({ name: payload.target })
        break
    }
  }
})
