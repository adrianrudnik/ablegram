<template>
  <div class="IndexTag">
    <div class="Label">
      Structure:
        <div class="BaseName">
          {{ base }}
        </div>

        <div class="ValueType" v-if="value">
          ={{'{'}}{{ value }}{{'}'}}
        </div>
    </div>

    <slot />
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  base: string,
  value: string,
}>()
</script>

<style lang="scss">
.IndexTag {
  .Label {
    display: flex;
    align-items: baseline;
    margin-top: 1rem;
  }

  .BaseName {
    margin-left: .5rem;
    font-family: var(--vp-font-family-mono);
  }

  .ValueType {
    font-family: var(--vp-font-family-mono);
    color: var(--vp-c-text-2);
  }
}
</style>
