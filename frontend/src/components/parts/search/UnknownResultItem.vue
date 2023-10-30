<template>
  <div>Unknown item: {{ result.type }}</div>
</template>

<script setup lang="ts">
import type { HitFieldset } from '@/plugins/search/result'

const props = defineProps<{
  result: HitFieldset
}>()
</script>

<style scoped lang="scss"></style>

<i18n lang="yaml"></i18n>
