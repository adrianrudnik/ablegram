import type { ResultType } from '@/plugins/search/result/index'

export interface AbletonAudioTrackResult {
  id: string
  type: ResultType.AbletonAudioTrack
  tags?: string[]

  pathAbsolute?: string
  pathFolder?: string
  filename?: string

  displayName?: string
  annotation?: string

  color?: number
}
