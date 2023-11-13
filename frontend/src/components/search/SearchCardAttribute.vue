<template>
  <div class="mt-2">
    <div class="k" v-if="props.k">{{ props.k }}</div>
    <div class="v" v-if="props.v">{{ props.v }}</div>
    <div class="v" v-if="!!$slots.default">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  k?: string
  v?: string
}>()
</script>

<style lang="scss" scoped>
div {
  line-height: 1.3rem;
}
.k {
  font-weight: bold;
  font-size: 0.8rem;
}
</style>
