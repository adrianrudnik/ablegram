<template>
  <div>
    <SearchTag :tag="tag" v-for="tag in tags" :key="tag.id" show-count />
  </div>
</template>

<script setup lang="ts">
import { useTagStore } from '@/stores/tags'
import { computed } from 'vue'
import SearchTag from '@/components/structure/SearchTag.vue'

const tags = computed(() => useTagStore().entries)
</script>
