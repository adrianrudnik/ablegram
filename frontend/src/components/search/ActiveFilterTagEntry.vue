<template>
  <SearchTag :tag="props.value" />
</template>

<script setup lang="ts">
import type { Tag } from '@/stores/tags'
import SearchTag from '@/components/structure/SearchTag.vue'

const props = defineProps<{
  value: Tag
}>()
</script>
