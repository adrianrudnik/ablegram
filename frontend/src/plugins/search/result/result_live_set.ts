import type { ResultType } from '@/plugins/search/result/index'

export interface LiveSetResult {
  id: string
  type: ResultType.LiveSet
  tags?: string[]

  displayName?: string
  filename?: string
  path?: string
  majorVersion?: string
  minorVersion?: string
  creator?: string
  revision?: string

  scaleRootNote?: string
  scaleName?: string
  scale?: string

  inKey?: boolean
  tempo?: number
}
