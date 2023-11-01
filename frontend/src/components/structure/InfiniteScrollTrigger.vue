<template>
  <div class="infinite-scroll-trigger" ref="trigger"></div>
</template>

<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref } from 'vue'

const emit = defineEmits(['trigger'])

const observer = new IntersectionObserver(
  (entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        emit('trigger')
      }
    })
  },
  { threshold: 0.5 }
)

const trigger = ref<Element>()

onMounted(() => {
  if (trigger.value) {
    console.log('OBSERVE')
    observer.observe(trigger.value)
  }
})

onBeforeUnmount(() => {
  if (trigger.value) {
    console.log('UNOBSERVE')
    observer.unobserve(trigger.value)
  }
})
</script>
