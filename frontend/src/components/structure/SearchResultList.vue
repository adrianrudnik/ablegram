<template>
  <component
    v-for="result in results"
    :key="result.id"
    class="mb-3"
    :is="resolveComponent(result.type)"
    :result="result"
  />

  <InfiniteScrollTrigger @trigger="loadMore" class="mb-6" />
</template>

<script lang="ts" setup>
import LiveSetResultItem from '@/components/parts/search/LiveSetResultItem.vue'
import MidiTrackResultItem from '@/components/parts/search/MidiTrackResultItem.vue'
import AudioTrackResultItem from '@/components/parts/search/AudioTrackResultItem.vue'
import UnknownResultItem from '@/components/parts/search/UnknownResultItem.vue'
import { useSearchResultStore } from '@/stores/results'
import { computed } from 'vue'
import type { ResultType } from '@/plugins/search/result'
import InfiniteScrollTrigger from '@/components/structure/InfiniteScrollTrigger.vue'
import { useSearchStore } from '@/stores/search'

const currentQuery = computed(() => useSearchStore().currentQuery)
const lastSortKey = computed(() => useSearchStore().lastSortKey)

const { executeQuerySearch } = useSearchStore()

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

const loadMore = async () => {
  if (!currentQuery.value) {
    return
  }

  // Get the last result to retrieve the sort offset
  if (!lastSortKey.value) return

  const q = currentQuery.value
  if (lastSortKey.value) {
    q.search_after = lastSortKey.value
  }

  const result = await executeQuerySearch(q)

  useSearchResultStore().updateBatch(
    result.hits.map((h) => {
      h.fields.id = h.id
      return h.fields
    })
  )
}

const results = computed(() => useSearchResultStore().entries)
</script>
