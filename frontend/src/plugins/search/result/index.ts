import type { LiveSetResult } from '@/plugins/search/result/result_live_set'
import type { MidiTrackResult } from '@/plugins/search/result/result_midi_track'
import type { AudioTrackResult } from '@/plugins/search/result/result_audio_track'

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

export type HitFieldset = LiveSetResult | MidiTrackResult | AudioTrackResult
