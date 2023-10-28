<template>
  <div class="surface-card p-4 shadow-2 border-round">
    <div class="mb-3 flex flex-column md:flex-row md:align-items-center md:justify-content-between">
      <div class="mb-3 md:mb-0">
        <div class="text-3xl font-medium text-900 mb-3" v-if="title">{{ title }}</div>
        <div class="font-medium text-500 mb-3" v-if="!!$slots.subtitle">
          <slot name="subtitle" />
        </div>
      </div>

      <div
        class="flex align-items-center justify-content-between mt-3 md:mt-0"
        v-if="!!$slots.actions"
      >
        <slot name="actions" />
      </div>
    </div>

    <slot />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  title?: string
}>()
</script>

<style scoped lang="scss"></style>

<i18n lang="yaml"></i18n>
