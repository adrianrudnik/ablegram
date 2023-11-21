<template>
  <label v-if="title">{{ title }}</label>
  <div class="flex flex-wrap gap-3 mb-3">
    <slot />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  title?: string
}>()
</script>
