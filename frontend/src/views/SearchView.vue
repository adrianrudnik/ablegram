<template>
  <MetaBar class="mb-3" />
  <QueryInput class="mb-3" v-if="dataPoints" />
  <SearchResultList />
</template>

<script setup lang="ts">
import MetaBar from '@/components/parts/MetaBar.vue'
import QueryInput from '@/components/parts/QueryInput.vue'
import SearchResultList from '@/components/structure/SearchResultList.vue'
import { computed } from 'vue'
import { useStatStore } from '@/stores/stats'

const dataPoints = computed(() => useStatStore().indexDocumentCount)
</script>
