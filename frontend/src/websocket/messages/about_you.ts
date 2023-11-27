import { PushMessageType } from '@/websocket/messages/global'

export interface ClientIdPushMessage {
  type: PushMessageType.ClientId
  id: string
}
