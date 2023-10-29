<template>
  <div class="SearchTag mb-1 mr-1 inline-block">
    <div class="Parts flex align-items-center gap-0">
      <span class="p-1 px-2 bg-black-alpha-90 text-white border-900">
        {{ t(props.tag.trans.topic) }}
      </span>
      <span class="p-1 px-2 bg-black-alpha-60 text-white border-900">
        {{ t(props.tag.trans.detail) }}
      </span>
      <span v-if="props.tag.trans.extra" class="p-1 px-2 border-900">
        {{ t(props.tag.trans.extra) }}
      </span>
      <span v-if="showCount" class="p-1 px-2 bg-black-alpha-30 border-900">
        {{ props.tag.count }}
      </span>

      <!--      <div class="p-1 px-2 text-base bg-black-alpha-10 border-900">-->
      <!--        {{ t(props.tag.trans.detail) }}-->
      <!--      </div>-->
      <!--      <div v-if="parts.detail" class="p-1 px-2 bg-black-alpha-90 text-white border-900">-->
      <!--        {{ t(props.tag.trans.extra) }}-->
      <!--      </div>-->
      <!--      <div v-if="parts.extra" class="p-1 px-2 bg-black-alpha-50 text-white border-900">-->
      <!--        {{ props.tag.extra }}-->
      <!--      </div>-->
      <!--      <div v-if="showCount" class="p-1 px-2 bg-black-alpha-30 border-900">-->
      <!--        {{ props.tag.count }}-->
      <!--      </div>-->
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Tag } from '@/stores/tags'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = withDefaults(
  defineProps<{
    tag: Tag
    showCount?: boolean
    disableTranslation?: boolean
  }>(),
  {
    showRealm: false,
    showCount: false,
    disableTranslation: false
  }
)

//
// const pt = (part: number, value: string | null): string => {
//   switch(part) {
//
//   }
// }
//
// const translate = (prefix: string, value: string | number | null): string | null => {
//   if (props.disableTranslation) return String(value)
//
//   if (value === null) {
//     return null
//   }
//
//   if (prefix.endsWith('ableton:version:')) {
//     return value.toString()
//   }
//
//   if (prefix.endsWith('time-weekday:')) {
//     return t('datetime.weekday.' + value.toString())
//   }
//
//   if (prefix.endsWith('time-month:')) {
//     return t('datetime.month.' + value.toString())
//   }
//
//   if (prefix.endsWith('time-quarter:')) {
//     return t('datetime.quarter.' + value.toString())
//   }
//
//   if (prefix.endsWith('zodiac-western:')) {
//     return t('datetime.zodiac-western.' + value.toString())
//   }
//
//   if (prefix.endsWith('zodiac-chinese:')) {
//     return t('datetime.zodiac-chinese.' + value.toString())
//   }
//
//   if (typeof value === 'number') {
//     return value.toString()
//   }
//
//   if (!Number.isNaN(Number(value))) {
//     return value.toString()
//   }
//
//   return t(prefix + value)
// }
</script>

<style lang="scss">
.SearchTag {
  cursor: default;

  & > div > span {
    border-top-width: 2px !important;
    border-top-style: solid;
    border-bottom-width: 2px !important;
    border-bottom-style: solid;
  }

  & > div > span:first-child {
    border-left-width: 2px !important;
    border-left-style: solid;
    border-top-left-radius: var(--border-radius) !important;
    border-bottom-left-radius: var(--border-radius) !important;
  }

  & > div > span:last-child {
    border-right-width: 2px !important;
    border-right-style: solid;
    border-top-right-radius: var(--border-radius) !important;
    border-bottom-right-radius: var(--border-radius) !important;
  }
}
</style>
