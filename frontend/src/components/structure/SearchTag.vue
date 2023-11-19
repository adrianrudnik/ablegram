<template>
  <div class="ResultTag inline-block">
    <div class="flex align-items-center gap-0">
      <span class="p-1 px-2" v-for="(part, idx) in tag.parts" :key="idx">
        {{ tag.trans.parts[idx] ?? part }}
      </span>
      <span
        v-if="tag.color !== undefined"
        :style="'background-color: ' + tag.color"
        class="p-1 px-2"
        >&nbsp;</span
      >
      <span class="TagValue p-1 px-2" v-if="tag.value !== undefined">
        {{ tag.trans.value ?? tag.value }}
      </span>
      <div v-if="showCount" class="TagCount p-1 px-2">
        {{ n(tag.count) }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Tag } from '@/stores/tags'
import { useI18n } from 'vue-i18n'

const { n } = useI18n()

const props = withDefaults(
  defineProps<{
    tag: Tag
    showCount?: boolean
    disableTranslation?: boolean
  }>(),
  {
    showRealm: false,
    showCount: false,
    disableTranslation: false
  }
)
</script>

<style lang="scss">
.ResultTag {
  // Ensure the tags do not break up
  & > div > span {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  // Ensure a common border setup for all parts
  & > div > span {
    border-color: var(--surface-900) !important;
    border-top-width: 2px !important;
    border-top-style: solid;
    border-bottom-width: 2px !important;
    border-bottom-style: solid;
  }

  // Leading part should have a left border radius and darkest color
  & > div > span:first-child {
    border-left-width: 2px !important;
    border-left-style: solid;
    border-top-left-radius: var(--border-radius) !important;
    border-bottom-left-radius: var(--border-radius) !important;
    background-color: var(--gray-900);
    color: white;
  }

  & > div > span:nth-child(2) {
    background-color: var(--gray-600);
    color: white;
  }

  & > div > span:nth-child(3) {
    background-color: var(--gray-300);
    color: black;
  }

  // Ensure the last part (mostly the value one) is pure white with black font
  & > div > span:last-of-type {
    border-right-width: 2px !important;
    border-right-style: solid;
    border-top-right-radius: var(--border-radius) !important;
    border-bottom-right-radius: var(--border-radius) !important;
  }

  & > div > span:last-of-type {
    background-color: white;
    color: black;
  }

  & > div > div.TagCount {
    font-size: 0.8em;
    color: black;
    border: 1px dashed black;
    border-left: none;
    border-top-right-radius: var(--border-radius) !important;
    border-bottom-right-radius: var(--border-radius) !important;
  }
}
</style>
