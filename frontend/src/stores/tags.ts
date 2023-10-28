import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'
import type { FileStatus } from '@/stores/files'
import { fetchApi } from '@/plugins/api'
import { executeQuerySearch } from '@/plugins/search'

export interface Tag {
  type: 'sys'
  id: string
  value?: number
  count: number | string // string for example tag
}

export const useTagStore = defineStore('tags', setupStore<Tag>())

export const hydrateTags = async () => {
  const tags = useTagStore()
  tags.clear()

  const r = await executeQuerySearch({
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
    const t: Tag = {
      id: term.term,
      count: term.count,
      type: 'sys'
    }

    // The value part might be a number that we need to sort by later on,
    // so we optimize that here.
    const s = term.term.split(':')

    if (typeof s[3] !== 'undefined') {
      // Plain number is easily to parse
      if (!Number.isNaN(Number(s[3]))) {
        t.value = parseInt(s[3])
      }

      // Version number like, we strip everything until the first letter
      if (s[3].match(/^v?\d/) && s[3].includes('.')) {
        const vn = parseVersionNumber(s[3])
        if (vn) t.value = vn
      }
    }

    tags.update(t)
  }
}

function parseVersionNumber(v: string): number | null {
  if (v.match(/^v?\d/) && v.includes('.')) {
    // Extract a common version number, skipping on typical alpha characters at the end.
    const e1 = v.match(/^(\d+\.)?(\d+\.)?(\d+)/)
    if (!e1) return null

    // Split and reverse the version number
    // Then multiply each part with a exponential number to get a sortable number
    const e2 = e1[0].split('.').reverse()
    const exps = [10e5, 10e2, 1]

    let e3 = 0
    for (let i = 0; i < exps.length; i++) {
      if (e2[i] === undefined) continue
      e3 += parseInt(e2[i]) * exps[i]
    }

    return e3
  }

  return null
}
