import { defineStore } from 'pinia'
import { setupStore } from '@/stores/base'
import { executeQuerySearch } from '@/plugins/search'
import i18n from '@/plugins/i18n'
import { resolveColorByIndex } from '@/plugins/colors'

const { t } = i18n.global

export const enum TagType {
  StringValue = 'string',
  NumValue = 'numeric',
  VersionValue = 'version',
  ColorValue = 'color'
}

export const enum TagCategory {
  Unknown,
  Type,
  Software,
  Info,
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
  plain?: string // contains the full tag in plain translated form for searches and copy/pastes
}

export interface Tag {
  id: string
  type: TagType
  category: TagCategory
  color?: string
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

  const tag: Tag = {
    id: term,
    type: TagType.StringValue,
    category: TagCategory.Unknown,
    realm: parts[0],
    topic: parts[1],
    detail: parts[2],
    extra: parts[3] ?? undefined,
    count: count,
    trans: {
      topic: t('tags.' + parts.slice(0, 2).join(':')),
      detail: t('tags.' + parts.slice(0, 3).join(':')),
      extra: parts[3] ? t('tags.' + parts.slice(0, 4).join(':')) : undefined
    }
  }

  tag.category = categorizeTag(tag)

  // If we have a 4th part, there is a value to be attached
  if (parts.length === 4) {
    tag.extra = parts[3]
    const type = classifyTag(tag)

    switch (type) {
      case TagType.NumValue:
        if (!Number.isNaN(Number(parts[3]))) {
          tag.value = parseInt(parts[3])
          tag.trans.extra = t(tag.value.toString())
          tag.type = type
        }
        break
      case TagType.VersionValue:
        tag.value = parseVersionNumber(parts[3]) ?? undefined
        if (tag.value) {
          tag.type = type
          tag.trans.extra = t(tag.extra)
        }
        break
      case TagType.ColorValue:
        tag.value = parseInt(parts[3])
        if (tag.value) {
          tag.type = type
          tag.trans.extra = t('color.' + tag.extra)
          tag.color = colorizeTag(tag) ?? undefined
        }
    }
  }

  overrideTranslations(tag)

  return tag
}

function categorizeTag(tag: Tag): TagCategory {
  if (tag.topic === 'type') return TagCategory.Type
  if (tag.topic === 'info') return TagCategory.Info
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

  if (tag.topic === 'color') {
    return TagType.ColorValue
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
  // Special handling for mtime/btime fields with shared logic
  if (tag.topic === 'file') {
    const marker = tag.detail.substring(1)

    switch (marker) {
      case 'time-quarter':
        tag.trans.extra = t('datetime.quarter.' + tag.value)
        break
      case 'time-weekday':
        tag.trans.extra = t('datetime.weekday.' + tag.value)
        break
      case 'time-month':
        tag.trans.extra = t('datetime.month.' + tag.value)
        break
    }

    switch (tag.detail) {
      case 'zodiac-western':
        tag.trans.extra = t('datetime.zodiac-western.' + tag.extra)
        break
      case 'zodiac-chinese':
        tag.trans.extra = t('datetime.zodiac-chinese.' + tag.extra)
        break
    }
  }

  // Finalize the plain text translated tag
  tag.trans.plain = [tag.trans.topic, tag.trans.detail, tag.trans.extra]
    .filter((v) => !!v)
    .join(':')
}

function colorizeTag(tag: Tag): string | null {
  if (tag.type !== TagType.ColorValue) return null
  if (typeof tag.value !== 'number') return null
  return resolveColorByIndex(tag.value)
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
