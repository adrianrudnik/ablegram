import { useWebSocket } from '@vueuse/core'
import { useFilesStore } from '@/stores/files'
import { nextTick } from 'vue'
import {useStatStore} from "@/stores/stats";

export const websocket = useWebSocket(import.meta.env.VITE_WEBSOCKET_URL, {
  autoReconnect: true,
  onMessage(ws, event) {
    const payload = JSON.parse(event.data)

    switch (payload.type) {
      case 'file_status':
        useFilesStore().update({
          id: payload.path,
          ...payload
        })
        break

      case 'index_status':
        useStatStore().documents = payload.document_count;
        break
    }
  }
})
