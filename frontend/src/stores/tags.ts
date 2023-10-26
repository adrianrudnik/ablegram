import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'
import type { FileStatus } from '@/stores/files'
import { fetchApi, fetchSearch } from '@/plugins/api'

export type TagType = 'system'

export interface Tag {
  type: TagType
  id: string
  value?: string
  count: number
}

export const useTagStore = defineStore('tags', setupStore<Tag>())

export const hydrateTags = async () => {
  const tags = useTagStore()
  tags.clear()

  const r = await fetchSearch({
    size: 4,
    query: {
      query: '*'
    },
    facets: {
      tags: {
        field: 'tags',
        size: 1000
      }
    }
  })

  if (!r.facets?.tags.terms) {
    return
  }

  for (const term of r.facets?.tags.terms ?? []) {
    tags.update({
      id: term.term,
      count: term.count,
      type: 'system'
    })
  }
}
