import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'
import { useTagStore } from '@/stores/tags'

export enum ActiveFilterType {
  TAG = 'tags'
}

export enum ActiveFilterMode {
  SHOULD = 'should',
  MUST = 'must',
  NOT = 'not'
}

export interface ActiveFilter {
  id: string
  type: ActiveFilterType
  mode: ActiveFilterMode
  content: any
  query: string
}

export const useActiveFiltersStore = defineStore('search-filters', setupStore<ActiveFilter>())

export function compileFilter(v: ActiveFilter): string {
  let prefix = ''
  switch (v.mode) {
    case 'must':
      prefix = '+'
      break
    case 'not':
      prefix = '-'
      break
  }

  return `${prefix}${v.query}`
}

export function addRandomTagFilter() {
  const tag = useTagStore().getRandomElement()

  if (!tag) {
    return
  }

  useActiveFiltersStore().update({
    id: tag.id,
    type: ActiveFilterType.TAG,
    mode: ActiveFilterMode.SHOULD,
    content: tag,
    query: 'tags:"' + tag.id + '"'
  })
}
