<template>
  <component
    v-for="result in results"
    :key="result.id"
    class="mb-3"
    :is="resolveComponent(result.type)"
    :result="result"
  />
</template>

<script lang="ts" setup>
import LiveSetResultItem from '@/components/parts/search/LiveSetResultItem.vue'
import MidiTrackResultItem from '@/components/parts/search/MidiTrackResultItem.vue'
import AudioTrackResultItem from '@/components/parts/search/AudioTrackResultItem.vue'
import UnknownResultItem from '@/components/parts/search/UnknownResultItem.vue'
import { useSearchResultStore } from '@/stores/results'
import { computed } from 'vue'
import type { ResultType } from '@/plugins/search/result'

function resolveComponent(type: ResultType): any {
  switch (type) {
    case 'LiveSet':
      return LiveSetResultItem
    case 'MidiTrack':
      return MidiTrackResultItem
    case 'AudioTrack':
      return AudioTrackResultItem
    default:
      return UnknownResultItem
  }
}

const results = computed(() => useSearchResultStore().entries)
</script>
