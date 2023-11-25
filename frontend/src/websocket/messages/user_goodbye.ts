import { PushMessageType } from '@/websocket/messages/global'

export interface UserGoodbyePushMessage {
  type: PushMessageType.UserGoodbye
  id: string
}
