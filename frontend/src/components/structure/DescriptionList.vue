<template>
  <div class="surface-section">
    <div class="font-medium text-xl text-900 mb-3">{{ props.title }}</div>
    <div v-if="!!$slots.description" class="text-500 mb-5">
      <slot name="description" />
    </div>
    <ul class="list-none p-0 m-0">
      <slot />
    </ul>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  title: string
}>()
</script>
