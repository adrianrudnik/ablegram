import { PushMessageType } from '@/websocket/messages/global'

export interface MetricUpdatePushMessage {
  type: PushMessageType.MetricUpdate
  values: { [key: string]: number }
}
