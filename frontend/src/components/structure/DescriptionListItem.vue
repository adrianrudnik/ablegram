<template>
  <li class="flex align-items-center py-3 px-2 border-top-1 surface-border flex-wrap">
    <div class="text-500 w-6 md:w-2 font-medium">{{ props.title }}</div>
    <div class="text-900 w-full md:w-8 md:flex-order-0 flex-order-1">
      <slot />
    </div>
    <div class="w-6 md:w-2 flex justify-content-end" v-if="!!$slots.actions">
      <slot name="actions" />
    </div>
  </li>
</template>

<script setup lang="ts">
const props = defineProps<{
  title: string
}>()
</script>
