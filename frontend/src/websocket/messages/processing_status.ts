import { PushMessageType } from '@/websocket/messages/global'

export interface ProcessingStatusPushMessage extends File {
  type: PushMessageType.ProcessingStatus
  routines: number
}
