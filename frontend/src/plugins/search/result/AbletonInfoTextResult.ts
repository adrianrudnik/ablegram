import type { ResultType } from '@/plugins/search/result/index'

export interface AbletonInfoTextResult {
  id: string
  type: ResultType.AbletonSampleReference
  tags?: string[]
  pathAbsolute?: string
  pathFolder?: string
  filename?: string

  displayName?: string

  parent?: string
}
