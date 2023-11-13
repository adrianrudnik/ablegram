<template>
  <SearchResultCard :result="result" :expanded="expanded" />
</template>

<script setup lang="ts">
import SearchResultCard from '@/components/search/SearchResultCard.vue'
import type { AbletonMidiTrackResult } from '@/plugins/search/result/AbletonMidiTrackResult'

const props = defineProps<{
  result: AbletonMidiTrackResult
  expanded: boolean
}>()
</script>
