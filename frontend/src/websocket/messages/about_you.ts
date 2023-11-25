import { PushMessageType } from '@/websocket/messages/global'

export interface AboutYouPushMessage {
  type: PushMessageType.AboutYou
  id: string
}
