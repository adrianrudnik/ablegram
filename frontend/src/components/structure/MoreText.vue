<template>
  <div v-if="isExpanded">
    {{ props.text }}
  </div>
  <div v-else>
    {{ shortText }}
    <button @click="isExpanded = true" class="p-button-text">more</button>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { truncate } from 'lodash'

const props = withDefaults(
  defineProps<{
    expanded?: boolean
    length?: number
    text: string
  }>(),
  {
    expanded: false,
    length: 90
  }
)

const isExpanded = computed(() => props.expanded || props.text.length <= props.length)
const shortText = computed(() => truncate(props.text, { length: props.length }))
</script>

<style scoped lang="scss">
button {
  border: unset;

  &:hover {
    background-color: black;
    color: white;
  }
}
</style>
