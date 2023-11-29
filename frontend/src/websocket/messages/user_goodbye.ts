import { PushMessageType } from '@/websocket/messages/global'

export interface ClientGoodbyePushMessage {
  type: PushMessageType.ClientGoodbye
  id: string
}
