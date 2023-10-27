import { PushMessageType } from '@/websocket/messages/global'

export interface MetricUpdatePushMessage {
  type: PushMessageType.MetricUpdate
  k: string
  v: number
}
