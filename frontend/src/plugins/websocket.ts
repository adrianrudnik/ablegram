import { useWebSocket } from '@vueuse/core'
import type { File } from '@/stores/files'
import { useFilesStore } from '@/stores/files'
import { useStatStore } from '@/stores/stats'

export const websocket = useWebSocket(import.meta.env.VITE_WEBSOCKET_URL, {
  autoReconnect: true,
  onMessage(ws, event) {
    const payload = JSON.parse(event.data) as PushMessage

    switch (payload.type) {
      case PushMessageType.FileStatus:
        useFilesStore().update(payload)
        break

      case PushMessageType.MetricUpdate:
        useStatStore().update(payload.k, payload.v)
        break
    }
  }
})

enum PushMessageType {
  FileStatus = 'file_status',
  MetricUpdate = 'metric_update'
}

interface FileStatusPushMessage extends File {
  type: PushMessageType.FileStatus
}

interface MetricUpdatePushMessage {
  type: PushMessageType.MetricUpdate
  k: string
  v: number
}

type PushMessage = FileStatusPushMessage | MetricUpdatePushMessage
