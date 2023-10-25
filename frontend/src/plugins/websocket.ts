import { useWebSocket } from '@vueuse/core'

export const websocket = useWebSocket(import.meta.env.VITE_WEBSOCKET_URL, {
  autoReconnect: true
})
