import { PushMessageType } from '@/websocket/messages/global'
import type { UserRoles } from '@/stores/users'

export interface UserCurrentPushMessage {
  type: PushMessageType.UserCurrent
  id: string
  display_name: string
  role: UserRoles
  ip?: string
}
