import type { ResultType } from '@/plugins/search/result/index'

export interface AbletonSampleReferenceResult {
  id: string
  type: ResultType.AbletonSampleReference
  tags?: string[]
  pathAbsolute?: string
  pathFolder?: string
  filename?: string

  displayName?: string

  sampleAbsPath?: string
  sampleFilename?: string
  sampleOriginalFileSize?: number
}
