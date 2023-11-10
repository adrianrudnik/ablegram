import {defineStore} from "pinia";
import {setupStore} from "@/stores/base";
import {useSearchStore} from "@/stores/search";
import i18n from "@/plugins/i18n";
import {resolveAbletonColorByIndex} from "@/plugins/colors";

const { t, te, tm, locale } = i18n.global

export interface Tag {
  id: string
  base: string
  parts: string[]
  value?: string | number | boolean
  count: number
  color?: string
  trans: TagTranslation // String is used for example tags
  type: TagValueType
  search: string
}

interface TagTranslation {
  parts: string[]
  value?: string
}

export const enum TagValueType {
  BooleanValue = 'boolean',
  StringValue = 'string',
  NumericValue = 'numeric',
  SemverValue = 'semver',
  AbletonColorValue = 'ableton-color'
}

const numericValueTags = [
  'file:mtime-year',
  'file:mtime-weekday',
  'file:mtime-month',
  'file:mtime-quarter',
  'file:mtime-weekno',
  'file:btime-year',
  'file:btime-weekday',
  'file:btime-month',
  'file:btime-quarter',
  'file:btime-weekno',
  'time-signature:numerator',
  'time-signature:denominator',
]

export const useTagStore = defineStore('tags', setupStore<Tag>())

export const hydrateTags = async () => {
  const r = await useSearchStore().executeQuerySearch({
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

export function createTagFromString(raw: string, count: number): Tag | null {
  if (!raw || raw.trim() === "") return null

  const {base, value} = splitValueFromTag(raw)
  const parts = splitPartsFromTag(base)

  const tag: Tag = {
    id: raw,
    base: base,
    type: determineValueType(base, value),
    count: count,
    value: value,
    parts: parts,
    trans: { parts: []},
    search: '',
  }

  // Parse the type, if not string and value is present
  if (tag.type !== TagValueType.StringValue && value !== undefined) {
    switch(tag.type) {
      case TagValueType.BooleanValue:
        tag.value = value === 'true'
        break

      // Semver will become exponential number, to be correctly sortable.
      // We end up with a numeric value, but a string translation.
      case TagValueType.SemverValue:
        tag.value = parseVersionNumber(value) ?? undefined
        if (tag.value === undefined) tag.type = TagValueType.StringValue
        tag.trans.value = value
        break

      case TagValueType.NumericValue:
        if (!Number.isNaN(Number(value))) {
          tag.value = parseInt(value)
        } else {
          tag.type = TagValueType.StringValue
        }
        break

      case TagValueType.AbletonColorValue:
        if (!Number.isNaN(Number(value))) {
          const v = parseInt(value)
          tag.color = resolveAbletonColorByIndex(v) ?? undefined
          if (tag.color) {
            tag.value = v
          }
        } else {
          tag.type = TagValueType.StringValue
        }

        tag.trans.value = t('color.ableton.' + value)
        break
    }
  }

  compileTranslations(tag)

  compileSearchString(tag)

  return tag
}

function splitValueFromTag(tag: string): { base: string, value?: string } {
  const v = tag.split('=', 2)
  if (v.length === 2) {
    return {base: v[0], value: v[1] }
  }

  return {base: v[0]}
}

function splitPartsFromTag(tag: string): string[] {
  return tag.split(':')
}

function determineValueType(baseTag: string, value?: string): TagValueType {
  if (value !== undefined && (value === 'true' || value === 'false')) {
    return TagValueType.BooleanValue
  }

  if (baseTag === 'ableton:version') {
    return TagValueType.SemverValue
  }

  if (numericValueTags.includes(baseTag)) {
    return TagValueType.NumericValue
  }


  if (baseTag === 'color:ableton') {
    return TagValueType.AbletonColorValue
  }

  return TagValueType.StringValue;
}

function parseVersionNumber(v: string | undefined): number | null {
  if (v === undefined) return null

  if (v.match(/^v?[0-9]*$/)) {
    return (parseInt(v) + 1) * 10e5
  }

  if (v.match(/^v?\d/) && v.includes('.')) {
    // Extract a common version number, skipping on typical alpha characters at the end.
    const e1 = v.match(/^(\d+\.)?(\d+\.)?(\d+)/)
    if (!e1) return null

    // Split and reverse the version number
    // Then multiply each part with a exponential number to get a sortable number
    const e2 = e1[0].split('.')
    const exps = [10e5, 10e3, 1]

    let e3 = 0
    for (let i = 0; i < exps.length; i++) {
      if (e2[i] === undefined) continue
      e3 += (parseInt(e2[i]) + 1) * exps[i]
    }

    return e3
  }

  return null
}

function compileTranslations(tag: Tag) {
  // Simple boolean translation
  if (tag.type === TagValueType.BooleanValue) {
    tag.trans.value = t(tag.value ? 'common.label.true' : 'common.label.false')
  }

  // Special handling for mtime/btime fields with shared logic
  if (tag.id.startsWith('file:mtime') || tag.id.startsWith('file:btime')) {
    const marker = tag.parts[1].substring(1)

    switch (marker) {
      case 'time-quarter':
        tag.trans.value = t('datetime.quarter.' + tag.value)
        break;
      case 'time-weekday':
        tag.trans.value = t('datetime.weekday.' + tag.value)
        break;
      case 'time-month':
        tag.trans.value = t('datetime.month.' + tag.value)
        break;
    }
  }


  // Special handling for zodiac signs
  switch (tag.parts[1]) {
    case 'zodiac-western':
      tag.trans.value = t('datetime.zodiac-western.' + tag.value)
      break;
    case 'zodiac-chinese':
      tag.trans.value = t('datetime.zodiac-chinese.' + tag.value)
      break;
  }

  // Translate the parts of the tag by index
  for(let i = 0; i < tag.parts.length; i++) {
    tag.trans.parts[i] = t('tag.' + tag.parts.slice(0, i + 1).join(':'))
  }

  // If we have a string value, try to translate the full path value against the base
  // EN message catalog. If not found, we fall back to the literal string value.
  // Also respect already translated ones.
  if (tag.type === TagValueType.StringValue && tag.trans.value === undefined) {
    if (te('tag.' + tag.id, 'en')) {
      tag.trans.value = t('tag.' + tag.id)
    } else {
      tag.trans.value = tag.value?.toString() ?? ''
    }
  }
}

function compileSearchString(tag: Tag) {
  tag.search = [
      tag.parts.join(' '),
      tag.trans.parts.join(' '),
      (tag.trans.value ?? '')
  ].join(' ').toLowerCase()
}