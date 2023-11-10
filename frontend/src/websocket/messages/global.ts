import type { FileStatusPushMessage } from '@/websocket/messages/file_status'
import type { MetricUpdatePushMessage } from '@/websocket/messages/metric_update'
import type { ProcessingStatusPushMessage } from '@/websocket/messages/processing_status'
import type { ForceNavigatePushMessage } from '@/websocket/messages/force_navigate'
import type { TagUpdatePushMessage } from '@/websocket/messages/tag_update'

// TypeScript discriminator hell
// for websocket messages, anything to satisfy the god of types.

export enum PushMessageType {
  FileStatus = 'file_status',
  MetricUpdate = 'metric_update',
  TagUpdate = 'tag_update',
  ProcessingStatus = 'processing_status',
  ForceNavigate = 'force_navigate'
}

export type PushMessage =
  | FileStatusPushMessage
  | MetricUpdatePushMessage
  | TagUpdatePushMessage
  | ProcessingStatusPushMessage
  | ForceNavigatePushMessage
