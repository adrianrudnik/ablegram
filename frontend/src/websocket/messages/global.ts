import type { FileStatusPushMessage } from '@/websocket/messages/file_status'
import type { MetricUpdatePushMessage } from '@/websocket/messages/metric_update'
import type { ProcessingStatusPushMessage } from '@/websocket/messages/processing_status'

export enum PushMessageType {
  FileStatus = 'file_status',
  MetricUpdate = 'metric_update',
  ProcessingStatus = 'processing_status'
}

export type PushMessage =
  | FileStatusPushMessage
  | MetricUpdatePushMessage
  | ProcessingStatusPushMessage
