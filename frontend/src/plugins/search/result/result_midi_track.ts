import type { ResultType } from '@/plugins/search/result/index'

export interface MidiTrackResult {
  id: string
  type: ResultType.MidiTrack
  tags?: string[]

  pathAbsolute?: string
  pathFolder?: string
  filename?: string

  displayName?: string
  annotation?: string

  color?: number
}
