<template>
  <SearchQueryInput class="z-2 relative" />
  <SearchQueryToolbar class="mb-3 z-1 relative" />
  <ActiveFilterBar class="mb-3 z-0" />
  <SearchResultList />
</template>

<script setup lang="ts">
import SearchResultList from '@/components/search/SearchResultList.vue'
import SearchQueryInput from '@/components/search/SearchQueryInput.vue'
import ActiveFilterBar from '@/components/search/ActiveFilterBar.vue'
import SearchQueryToolbar from '@/components/search/SearchQueryToolbar.vue'
</script>
