<template>
  <h3 v-if="title">{{ title }}</h3>
  <div class="flex flex-wrap gap-3">
    <slot />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  title?: string
}>()
</script>
