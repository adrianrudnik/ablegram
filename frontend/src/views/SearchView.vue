<template>
  <!--  <MetaBar class="mb-3" />-->
  <SearchQueryInput />
  <!--  <QueryInput class="mb-3" v-if="dataPoints" />-->
  <!--  <SearchResultList />-->

  <SearchResultList />
</template>

<script setup lang="ts">
import SearchResultList from '@/components/search/SearchResultList.vue'
import { computed } from 'vue'
import { useStatStore } from '@/stores/stats'
import SearchQueryInput from '@/components/search/SearchQueryInput.vue'
</script>
