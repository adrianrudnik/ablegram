<template>
  <li class="p-0 mb-3">
    <component
      :is="resultTypeComponentMap[props.result.type] ?? UnknownResultItem"
      :result="props.result"
    />
  </li>
</template>

<script setup lang="ts">
import type { HitFieldset } from '@/plugins/search/result'
import LiveSetResultItem from '@/components/parts/search/LiveSetResultItem.vue'
import UnknownResultItem from '@/components/parts/search/UnknownResultItem.vue'
import MidiTrackResultItem from '@/components/parts/search/MidiTrackResultItem.vue'

const resultTypeComponentMap = {
  LiveSet: LiveSetResultItem,
  MidiTrack: MidiTrackResultItem,
  AudioTrack: UnknownResultItem
}

const props = defineProps<{
  result: HitFieldset
}>()
</script>
