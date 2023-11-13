import type { ResultType } from '@/plugins/search/result/index'

export interface AbletonMdiClipResult {
  id: string
  type: ResultType.AbletonSampleReference
  tags?: string[]
  pathAbsolute?: string
  pathFolder?: string
  filename?: string

  displayName?: string

  scaleRootNote?: string
  scaleName?: string

  timeSignature?: string
}
