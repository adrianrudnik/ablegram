<template>
  <SearchResultCard :result="result" :expanded="expanded">
    <div v-if="expanded">
      <SearchCardAttribute
        v-if="parent"
        :v="parent"
        :k="t('ableton-info-text-result-card.found-in.title')"
      />
    </div>
  </SearchResultCard>
</template>

<script setup lang="ts">
import SearchResultCard from '@/components/search/SearchResultCard.vue'
import SearchCardAttribute from '@/components/search/SearchCardAttribute.vue'
import { useI18n } from 'vue-i18n'
import type { AbletonInfoTextResult } from '@/plugins/search/result/AbletonInfoTextResult'

const { t } = useI18n()

const props = withDefaults(
  defineProps<{
    result: AbletonInfoTextResult
    expanded?: boolean
  }>(),
  {
    expanded: false
  }
)

const parent = props.result.parent?.replace(/\//g, ' > ')
</script>
