import type { LiveSetResult } from '@/plugins/search/result/result_live_set'

export interface SearchResult {
  status: StatusNode
  facets?: FacetsNode
  hits: Hit[]
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
  LiveSet = 'LiveSet'
}

export type HitFieldset = LiveSetResult
