<template>
  <div ref="trigger"></div>
</template>

<script setup lang="ts">
import { nextTick, onBeforeUnmount, ref, watch } from 'vue'
import { useIntersectionObserver, watchImmediate } from '@vueuse/core'

const emit = defineEmits(['trigger'])

const trigger = ref()

const { stop } = useIntersectionObserver(trigger, ([{ isIntersecting }], observerElement) => {
  if (isIntersecting) emit('trigger')
})

onBeforeUnmount(() => stop())
</script>
