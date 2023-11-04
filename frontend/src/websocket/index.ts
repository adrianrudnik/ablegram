import { useWebSocket } from '@vueuse/core'
import { useFilesStore } from '@/stores/files'
import { useStatStore } from '@/stores/stats'
import type { PushMessage } from '@/websocket/messages/global'
import { PushMessageType } from '@/websocket/messages/global'
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
        useStatStore().processRoutineCount = payload.routines
        useStatStore().isProcessing = payload.routines !== 0
        break

      case PushMessageType.ForceNavigate:
        await router.push({ name: payload.target })
        break
    }
  }
})
