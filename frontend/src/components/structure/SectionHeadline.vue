<template>
  <div class="mb-4">
    <div
      class="SectionHeadline flex md:align-items-center md:justify-content-between flex-column md:flex-row pb-2"
    >
      <div>
        <div class="title text-2xl font-medium mb-3" v-if="props.title">{{ props.title }}</div>
        <div class="title text-2xl font-medium mb-3" v-if="!!$slots.title">
          <slot name="title" />
        </div>
        <div class="desc mr-0 md:mr-3" v-if="!!$slots.description">
          <slot name="description" />
        </div>
      </div>
      <span class="right p-input-icon-left w-full md:w-auto" v-if="$slots.right">
        <slot name="right" />
      </span>
    </div>

    <slot />
  </div>
</template>

<script lang="ts" setup>
const props = defineProps<{
  title?: string
}>()
</script>

<style lang="scss">
.SectionHeadline {
  .title {
    .p-tag {
      position: relative;
      top: -0.2rem;
    }
  }
}
</style>
