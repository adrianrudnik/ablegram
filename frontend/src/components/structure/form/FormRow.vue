<template>
  <div class="field mb-4">
    <span class="font-medium text-xl" v-if="title">{{ props.title }}</span>
    <slot />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  title?: string
}>()
</script>
