import type { ResultType } from '@/plugins/search/result/index'

export interface MidiTrackResult {
  id: string
  type: ResultType.MidiTrack
  tags?: string[]

  displayName?: string
  filename?: string
}
