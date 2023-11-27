import type { FileStatusPushMessage } from '@/websocket/messages/file_status'
import type { MetricUpdatePushMessage } from '@/websocket/messages/metric_update'
import type { ProcessingStatusPushMessage } from '@/websocket/messages/processing_status'
import type { ForceNavigatePushMessage } from '@/websocket/messages/force_navigate'
import type { TagUpdatePushMessage } from '@/websocket/messages/tag_update'
import type { ClientWelcomePushMessage } from '@/websocket/messages/user_welcome'
import type { ClientGoodbyePushMessage } from '@/websocket/messages/user_goodbye'
import type { ClientIdPushMessage } from '@/websocket/messages/about_you'
import type { UserClientPushMessage } from '@/websocket/messages/user_current'

// TypeScript discriminator hell
// for websocket messages, anything to satisfy the god of types.

export enum PushMessageType {
  ClientId = 'client_id',
  FileStatus = 'file_status',
  MetricUpdate = 'metric_update',
  TagUpdate = 'tag_update',
  ProcessingStatus = 'processing_status',
  ForceNavigate = 'force_navigate',
  ClientWelcome = 'client_welcome',
  ClientGoodbye = 'client_goodbye',
  UserClient = 'user_client'
}

export type PushMessage =
  | FileStatusPushMessage
  | MetricUpdatePushMessage
  | TagUpdatePushMessage
  | ProcessingStatusPushMessage
  | ForceNavigatePushMessage
  | ClientWelcomePushMessage
  | ClientGoodbyePushMessage
  | UserClientPushMessage
  | ClientIdPushMessage
