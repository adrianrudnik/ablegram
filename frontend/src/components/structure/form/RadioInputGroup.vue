<template>
  <div class="flex flex-column mb-3">
    <label v-if="title">{{ props.title }}</label>
    <small v-if="props.help" class="mt-1">{{ props.help }}</small>
  </div>
  <div class="grid grid-nogutter gap-3 mb-3">
    <slot />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  title?: string
  help?: string
}>()
</script>
