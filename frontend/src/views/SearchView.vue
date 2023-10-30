<template>
  <MetaBar class="mb-3" />
  <QueryInput class="mb-3" v-if="dataPoints" />
  <SearchResultList>
    <SearchResultListItem v-for="result of results" :key="result.id" :result="result" />
  </SearchResultList>
</template>

<script setup lang="ts">
import MetaBar from '@/components/parts/MetaBar.vue'
import QueryInput from '@/components/parts/QueryInput.vue'
import SearchResultList from '@/components/structure/SearchResultList.vue'
import SearchResultListItem from '@/components/structure/SearchResultListItem.vue'
import { computed } from 'vue'
import { useSearchResultStore } from '@/stores/results'
import { useStatStore } from '@/stores/stats'

const dataPoints = computed(() => useStatStore().indexDocumentCount)

const results = computed(() => useSearchResultStore().entries)
</script>
