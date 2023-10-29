import type { File } from '@/stores/files'
import { PushMessageType } from '@/websocket/messages/global'

export interface ForceNavigatePushMessage extends File {
  type: PushMessageType.ForceNavigate
  target: string
}
