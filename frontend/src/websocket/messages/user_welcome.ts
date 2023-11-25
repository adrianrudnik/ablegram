import { PushMessageType } from '@/websocket/messages/global'
import type { UserRoles } from '@/stores/users'

export interface UserWelcomePushMessage {
  type: PushMessageType.UserWelcome
  id: string
  display_name: string
  role: UserRoles
}
