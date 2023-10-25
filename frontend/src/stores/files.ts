import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'

export type FileStatus = 'pending' | 'processed' | 'failed'

export interface File {
  id: string
  path: string
  folder: string
  filename: string
  status: FileStatus
  remark?: string
}

export const useFilesStore = defineStore('files', setupStore<File>())
