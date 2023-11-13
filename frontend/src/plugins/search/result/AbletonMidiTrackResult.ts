import type { ResultType } from '@/plugins/search/result/index'

export interface AbletonMidiTrackResult {
  id: string
  type: ResultType.AbletonMidiTrack
  tags?: string[]

  pathAbsolute?: string
  pathFolder?: string
  filename?: string

  displayName?: string
  annotation?: string

  color?: number
}
