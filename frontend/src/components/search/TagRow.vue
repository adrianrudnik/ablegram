<template>
  <div class="tags" v-if="tags">
    <SearchTag :tag="tag" v-for="tag in parsedTags" :key="tag.id" />
  </div>
</template>

<script setup lang="ts">
import type { Tag } from '@/stores/tags'
import { useTagStore } from '@/stores/tags'
import { computed } from 'vue'
import SearchTag from '@/components/structure/SearchTag.vue'
import sortBy from 'lodash/sortBy'
import orderBy from 'lodash/orderBy'

const props = defineProps<{
  tags: string[]
}>()

const parsedTags = computed(() => {
  const rawTags = props.tags?.filter((t) => useTagStore().entries.find((tt) => tt.id === t)) ?? []
  return orderBy(rawTags.map((t) => useTagStore().entries.find((tt) => tt.id === t)) as Tag[], [
    'search'
  ])
})
</script>
