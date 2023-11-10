<template>
  <SectionHeadline :title="props.title">
    <template #description>
      <slot />
    </template>
  </SectionHeadline>
</template>

<script setup lang="ts">
import SectionHeadline from '@/components/structure/SectionHeadline.vue'

const props = defineProps<{
  title: string
}>()
</script>
