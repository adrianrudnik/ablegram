<template>
  <div>
    <SearchResultCard :title="result.displayName">
      <div class="tags" v-if="tags">
        <SearchTag :tag="tag" v-for="tag in tags" :key="tag.id" />
      </div>
    </SearchResultCard>
  </div>
</template>

<script setup lang="ts">
import type { LiveSetResult } from '@/plugins/search/result/result_live_set'
import SearchResultCard from '@/components/parts/search/SearchResultCard.vue'
import { useTagStore } from '@/stores/tags'
import { computed } from 'vue'
import SearchTag from '@/components/structure/SearchTag.vue'
import type {Tag} from "@/stores/tags";

const props = defineProps<{
  result: LiveSetResult
}>()

const tags = computed(() => {
  const rawTags = props.result.tags?.filter(t => useTagStore().entries.find((tt) => tt.id === t)) ?? [];
  return rawTags.map(t => useTagStore().entries.find((tt) => tt.id === t)) as Tag[]
})
</script>

<style scoped lang="scss"></style>

<i18n lang="yaml"></i18n>
