<template>
  <div class="SearchTag mb-1 mr-1 inline-block">
    <div class="flex align-items-center gap-0">
      <div class="p-1 px-2 text-base bg-black-alpha-10 border-900">
        {{ parts.topic }}
      </div>
      <div v-if="parts.detail" class="p-1 px-2 bg-black-alpha-90 text-white border-900">
        {{ parts.detail }}
      </div>
      <div v-if="parts.extra" class="p-1 px-2 bg-black-alpha-50 text-white border-900">
        {{ parts.extra }}
      </div>
      <div v-if="showCount" class="p-1 px-2 bg-black-alpha-30 border-900">
        {{ props.tag.count }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Tag as TagInterface } from '@/stores/tags'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = withDefaults(
  defineProps<{
    tag: TagInterface
    showRealm?: boolean
    showCount?: boolean
    disableTranslation?: boolean
  }>(),
  {
    showRealm: false,
    showCount: false,
    disableTranslation: false
  }
)

const translate = (prefix: string, value: string | number | null): string | null => {
  if (props.disableTranslation) return String(value)

  if (value === null) {
    return null
  }

  if (prefix.endsWith('ableton:version:')) {
    return value.toString()
  }

  if (prefix.endsWith('time-weekday:')) {
    return t('datetime.weekday.' + value.toString())
  }

  if (prefix.endsWith('time-month:')) {
    return t('datetime.month.' + value.toString())
  }

  if (prefix.endsWith('time-quarter:')) {
    return t('datetime.quarter.' + value.toString())
  }

  if (prefix.endsWith('zodiac-western:')) {
    return t('datetime.zodiac-western.' + value.toString())
  }

  if (prefix.endsWith('zodiac-chinese:')) {
    return t('datetime.zodiac-chinese.' + value.toString())
  }

  if (typeof value === 'number') {
    return value.toString()
  }

  if (!Number.isNaN(Number(value))) {
    return value.toString()
  }

  return t(prefix + value)
}

const parts = computed(() => {
  const p = props.tag.id.split(':')
  return {
    realm: translate('tags.', p[0] ?? null),
    topic: translate('tags.' + p[0] + ':', p[1] ?? null),
    detail: translate('tags.' + p[0] + ':' + p[1] + ':', p[2] ?? null),
    extra: translate('tags.' + p[0] + ':' + p[1] + ':' + p[2] + ':', p[3] ?? null)
  }
})
</script>

<style lang="scss">
.SearchTag {
  cursor: default;

  & > div > div {
    border-top-width: 2px !important;
    border-top-style: solid;
    border-bottom-width: 2px !important;
    border-bottom-style: solid;
  }

  & > div > div:first-child {
    border-left-width: 2px !important;
    border-left-style: solid;
    border-top-left-radius: var(--border-radius) !important;
    border-bottom-left-radius: var(--border-radius) !important;
  }

  & > div > div:last-child {
    border-right-width: 2px !important;
    border-right-style: solid;
    border-top-right-radius: var(--border-radius) !important;
    border-bottom-right-radius: var(--border-radius) !important;
  }
}
</style>
