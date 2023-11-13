<template>
  <SearchResultCard :result="result" :expanded="expanded">
    <div v-if="expanded">
      <SearchCardAttribute
        v-if="props.result.scaleRootNote"
        :v="props.result.scaleRootNote"
        k="Scale root note"
      />
      <SearchCardAttribute
        v-if="props.result.scaleName"
        :v="props.result.scaleName"
        k="Scale name"
      />
      <SearchCardAttribute
        v-if="props.result.timeSignature"
        :v="props.result.timeSignature"
        k="Time signature"
      />
    </div>
  </SearchResultCard>
</template>

<script setup lang="ts">
import SearchResultCard from '@/components/search/SearchResultCard.vue'
import SearchCardAttribute from '@/components/search/SearchCardAttribute.vue'
import type { AbletonMdiClipResult } from '@/plugins/search/result/AbletonMdiClipResult'

const props = withDefaults(
  defineProps<{
    result: AbletonMdiClipResult
    expanded?: boolean
  }>(),
  {
    expanded: false
  }
)
</script>
