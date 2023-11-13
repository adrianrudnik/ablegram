import type { AbletonLiveSetResult } from '@/plugins/search/result/AbletonLiveSetResult'
import type { AbletonMidiTrackResult } from '@/plugins/search/result/AbletonMidiTrackResult'
import type { AbletonAudioTrackResult } from '@/plugins/search/result/AbletonAudioTrackResult'
import type { AbletonSampleReferenceResult } from '@/plugins/search/result/AbletonSampleReferenceResult'
import type { AbletonMdiClipResult } from '@/plugins/search/result/AbletonMdiClipResult'

export interface SearchResult {
  status: StatusNode
  facets?: FacetsNode
  hits: Hit[]
  total_hits: number
}

export interface StatusNode {
  total: number
  failed: number
  success: number
}

export interface FacetsNode {
  [key: string]: FacetNode
}

export interface FacetNode {
  field: string
  total: number
  missing: number
  other: number
  terms: FacetTerm[]
}

export interface FacetTerm {
  term: string
  count: number
}

export interface Hit {
  id: string
  score: number
  fields: HitFieldset
}

export enum ResultType {
  AbletonLiveSet = 'AbletonLiveSet',
  AbletonMidiTrack = 'AbletonMidiTrack',
  AbletonAudioTrack = 'AbletonAudioTrack',
  AbletonReturnTrack = 'AbletonReturnTrack',
  AbletonGroupTrack = 'AbletonGroupTrack',
  AbletonPreHearTrack = 'AbletonPreHearTrack',
  AbletonMidiClip = 'AbletonMidiClip',
  AbletonAudioClip = 'AbletonAudioClip',
  AbletonMixer = 'AbletonMixer',
  AbletonDeviceChain = 'AbletonDeviceChain',
  AbletonScene = 'AbletonScene',
  AbletonSampleReference = 'AbletonSampleReference'
}

export type HitFieldset =
  | AbletonLiveSetResult
  | AbletonMidiTrackResult
  | AbletonAudioTrackResult
  | AbletonSampleReferenceResult
  | AbletonMdiClipResult
