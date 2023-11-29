import { PushMessageType } from '@/websocket/messages/global'
import type { UserRoles } from '@/stores/users'

export interface UserClientPushMessage {
  type: PushMessageType.UserClient
  id: string
  ip?: string

  user_id: string
  user_display_name: string
  user_role: UserRoles
}
