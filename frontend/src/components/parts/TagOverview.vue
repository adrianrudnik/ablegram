<template>
  <div>
    <SearchTag :tag="tag" v-for="tag in tags" :key="tag.id" show-count />
  </div>
</template>

<script setup lang="ts">
import { hydrateTags, useTagStore } from '@/stores/tags'
import { computed, onMounted } from 'vue'
import SearchTag from '@/components/structure/SearchTag.vue'

const tags = computed(() => useTagStore().entries)

onMounted(async () => {
  await hydrateTags()
})
</script>
