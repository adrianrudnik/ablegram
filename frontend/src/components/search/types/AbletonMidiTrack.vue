<template>
  <SearchResultCard :result="result" />
</template>

<script setup lang="ts">
import SearchResultCard from '@/components/search/SearchResultCard.vue'
import type { MidiTrackResult } from '@/plugins/search/result/result_midi_track'

const props = defineProps<{
  result: MidiTrackResult
}>()
</script>
