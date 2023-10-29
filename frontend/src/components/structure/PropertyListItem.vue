<template>
  <li class="PropertyListItem flex align-items-center flex-wrap">
    <div class="w-6 md:w-2 font-medium">
      <i v-if="props.icon" :class="props.icon"></i> {{ props?.title ?? '' }}
    </div>
    <div class="w-full md:w-8 md:flex-order-0 flex-order-1">
      <slot />
    </div>
    <div class="w-6 md:w-2 flex justify-content-end" v-if="!!$slots.actions">
      <slot name="actions" />
    </div>
  </li>
</template>

<script setup lang="ts">
const props = defineProps<{
  title?: string
  icon?: string
}>()
</script>
