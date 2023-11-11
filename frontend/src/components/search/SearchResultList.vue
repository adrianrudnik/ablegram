<template>
  <div class="grid">
    <div
      class="flex align-self-stretch col-12 md:col-6 lg:col-4 xl:col-3"
      v-for="result in results"
      :key="result.id"
    >
      <component :is="resolveComponent(result.type)" :result="result" />
    </div>
  </div>

  <InfiniteScrollTrigger @trigger="loadMore" class="mb-6" />

  <div class="mx-auto my-5 text-center" v-if="isSearching">
    <i class="pi pi-spin pi-spinner" style="font-size: 2rem"></i>
  </div>
</template>

<script lang="ts" setup>
import SearchResultCard from '@/components/search/SearchResultCard.vue'
import { useSearchResultStore } from '@/stores/results'
import { computed } from 'vue'
import type { ResultType } from '@/plugins/search/result'
import InfiniteScrollTrigger from '@/components/structure/InfiniteScrollTrigger.vue'
import { useSearchStore } from '@/stores/search'
import AbletonMidiTrack from '@/components/search/types/AbletonMidiTrack.vue'
import { storeToRefs } from 'pinia'
import { useStatStore } from '@/stores/stats'

const { loadMore } = useSearchStore()
const { isSearching } = storeToRefs(useStatStore())

function resolveComponent(type: ResultType): any {
  switch (type) {
    case 'AbletonMidiTrack':
      return AbletonMidiTrack
    // case 'AbletonLiveSet':
    // case 'AbletonAudioTrack':
    // case 'AbletonReturnTrack':
    // case 'AbletonGroupTrack':
    // case 'AbletonPreHearTrack':
    // case 'AbletonMidiClip':
    // case 'AbletonAudioClip':
    // case 'AbletonMixer':
    // case 'AbletonDeviceChain':
    // case 'AbletonScene':
    //   return SearchResultCard
    default:
      return SearchResultCard
  }
}

const results = computed(() => useSearchResultStore().entries)
</script>
