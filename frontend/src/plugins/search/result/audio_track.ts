import type { ResultType } from '@/plugins/search/result/index'

export interface AudioTrackResult {
  id: string
  type: ResultType.AudioTrack
  tags?: string[]

  displayName?: string
  filename?: string
}
