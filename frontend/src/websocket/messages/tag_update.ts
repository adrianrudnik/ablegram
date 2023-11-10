import { PushMessageType } from '@/websocket/messages/global'

export interface TagUpdatePushMessage {
  type: PushMessageType.TagUpdate
  tags: { [key: string]: number }
}
