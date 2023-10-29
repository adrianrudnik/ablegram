import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'
import { executeQuerySearch } from '@/plugins/search'

export const enum TagType {
  StringValue = 'string',
  NumValue = 'numeric',
  VersionValue = 'version'
}

export const enum TagCategory {
  Unknown,
  Software,
  Location,
  Tracks,
  Tempo,
  Year,
  Month,
  Quarter,
  Weekday,
  WeekNumber,
  WesternZodiac,
  ChineseZodiac
}

interface TagTranslation {
  topic: string
  detail: string
  extra: string | undefined
}

export interface Tag {
  id: string
  type: TagType
  category: TagCategory
  realm: string
  topic: string
  detail: string
  extra: string | undefined // This will be transformed to a proper value into ->value.
  value?: number | string
  count: number | string
  trans: TagTranslation // String is used for example tags
}

export const useTagStore = defineStore('tags', setupStore<Tag>())

export const hydrateTags = async () => {
  const tags = useTagStore()

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

  // Parse the resulting tags
  for (const term of r.facets?.tags.terms ?? []) {
    const t = createTagFromString(term.term, term.count)
    if (t) useTagStore().update(t)
  }
}

export function createTagFromString(term: string, count: number | string): Tag | null {
  const parts = term.split(':')

  // Tags are only valid if they contain at least 3 parts
  if (parts.length < 3) return null

  const t: Tag = {
    id: term,
    type: TagType.StringValue,
    category: TagCategory.Unknown,
    realm: parts[0],
    topic: parts[1],
    detail: parts[2],
    extra: parts[3] ?? undefined,
    count: count,
    trans: {
      topic: 'tags.' + parts.slice(0, 2).join(':'),
      detail: 'tags.' + parts.slice(0, 3).join(':'),
      extra: parts[3] ? 'tags.' + parts.slice(0, 4).join(':') : undefined
    }
  }

  t.category = categorizeTag(t)

  // If we have a 4th part, there is a value to be attached
  if (parts.length === 4) {
    t.extra = parts[3]
    const type = classifyTag(t)

    switch (type) {
      case TagType.NumValue:
        if (!Number.isNaN(Number(parts[3]))) {
          t.value = parseInt(parts[3])
          t.trans.extra = t.value.toString()
          t.type = type
        }
        break
      case TagType.VersionValue:
        t.value = parseVersionNumber(parts[3]) ?? undefined
        if (t.value) {
          t.type = type
          t.trans.extra = t.extra
        }
        break
    }
  }

  overrideTranslations(t)

  return t
}

function categorizeTag(tag: Tag): TagCategory {
  if (tag.topic === 'ableton') return TagCategory.Software
  if (tag.topic === 'file' && tag.detail === 'location') return TagCategory.Location
  if (tag.topic === 'live-set' && tag.detail === 'tracks') return TagCategory.Tracks
  if (tag.topic === 'live-set' && tag.detail === 'tempo') return TagCategory.Tempo
  if (tag.topic === 'file' && tag.detail.substring(1) === 'time-year') return TagCategory.Year
  if (tag.topic === 'file' && tag.detail.substring(1) === 'time-month') return TagCategory.Month
  if (tag.topic === 'file' && tag.detail.substring(1) === 'time-quarter') return TagCategory.Quarter
  if (tag.topic === 'file' && tag.detail.substring(1) === 'time-weekday') return TagCategory.Weekday
  if (tag.topic === 'file' && tag.detail.substring(1) === 'time-weekno')
    return TagCategory.WeekNumber

  if (tag.topic === 'file' && tag.detail === 'zodiac-western') return TagCategory.WesternZodiac
  if (tag.topic === 'file' && tag.detail === 'zodiac-chinese') return TagCategory.ChineseZodiac

  return TagCategory.Unknown
}

function classifyTag(tag: Tag): TagType {
  if (tag.topic === 'ableton' && tag.detail === 'version') {
    return TagType.VersionValue
  }

  const numValTags = [
    'live-set:tempo',
    'file:mtime-year',
    'file:mtime-weekday',
    'file:mtime-month',
    'file:mtime-quarter',
    'file:mtime-weekno',
    'file:btime-year',
    'file:btime-weekday',
    'file:btime-month',
    'file:btime-quarter',
    'file:btime-weekno'
  ]

  if (numValTags.includes(tag.topic + ':' + tag.detail)) {
    return TagType.NumValue
  }

  return TagType.StringValue
}

function overrideTranslations(tag: Tag) {
  if (tag.topic === 'file') {
    const marker = tag.detail.substring(1)

    switch (marker) {
      case 'time-quarter':
        tag.trans.extra = 'datetime.quarter.' + tag.value
        break
      case 'time-weekday':
        tag.trans.extra = 'datetime.weekday.' + tag.value
        break
      case 'time-month':
        tag.trans.extra = 'datetime.month.' + tag.value
        break
    }

    switch (tag.detail) {
      case 'zodiac-western':
        tag.trans.extra = 'datetime.zodiac-western.' + tag.extra
        break
      case 'zodiac-chinese':
        tag.trans.extra = 'datetime.zodiac-chinese.' + tag.extra
        break
    }
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
