import type { File } from '@/stores/files'
import { PushMessageType } from '@/websocket/messages/global'

export interface FileStatusPushMessage extends File {
  type: PushMessageType.FileStatus
}
