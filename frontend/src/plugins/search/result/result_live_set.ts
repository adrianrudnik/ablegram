import type { ResultType } from '@/plugins/search/result/index'

export interface LiveSetResult {
  id: string
  type: ResultType.AbletonLiveSet
  tags?: string[]

  pathAbsolute?: string
  pathFolder?: string
  filename?: string

  displayName?: string
  majorVersion?: string
  minorVersion?: string
  creator?: string
  revision?: string

  scaleRootNote?: string
  scaleName?: string
  scale?: string

  inKey?: boolean
  bpm?: number

  midiTrackCount?: number
  audioTrackCount?: number
}
