<template>
  <SearchResultCard :result="result" :expanded="expanded">
    <div v-if="expanded">
      <SearchCardAttribute
        v-if="props.result.sampleFilename"
        :v="props.result.sampleFilename"
        k="Sample filename"
      />
      <SearchCardAttribute
        v-if="props.result.sampleAbsPath"
        :v="props.result.sampleAbsPath"
        k="Sample path"
      />
      <SearchCardAttribute v-if="props.result.sampleOriginalFileSize" k="Sample size (in bytes)">
        <i18n-t keypath="units.byte.nc" tag="span" :plural="props.result.sampleOriginalFileSize">
          <template v-slot:count>
            {{ n(props.result.sampleOriginalFileSize) }}
          </template>
        </i18n-t>
      </SearchCardAttribute>
    </div>
  </SearchResultCard>
</template>

<script setup lang="ts">
import SearchResultCard from '@/components/search/SearchResultCard.vue'
import SearchCardAttribute from '@/components/search/SearchCardAttribute.vue'
import type { AbletonSampleReferenceResult } from '@/plugins/search/result/AbletonSampleReferenceResult'
import { useI18n } from 'vue-i18n'

const { n } = useI18n()

const props = withDefaults(
  defineProps<{
    result: AbletonSampleReferenceResult
    expanded?: boolean
  }>(),
  {
    expanded: false
  }
)
</script>
