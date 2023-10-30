<template>
  <div class="SearchTag mb-1 mr-1 inline-block">
    <div class="Parts flex align-items-center gap-0">
      <span class="p-1 px-2 bg-black-alpha-90 text-white border-900">
        {{ props.tag.trans.topic }}
      </span>
      <span class="p-1 px-2 bg-black-alpha-60 text-white border-900">
        {{ props.tag.trans.detail }}
      </span>
      <span v-if="color" :style="'background-color: ' + color" class="p-1 px-2"> &nbsp; </span>
      <span v-if="props.tag.trans.extra" class="p-1 px-2 border-900">
        {{ props.tag.trans.extra }}
      </span>
      <span v-if="showCount" class="p-1 px-2 bg-gray-200 border-900">
        {{ props.tag.count }}
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Tag } from '@/stores/tags'
import { TagType } from '@/stores/tags'
import { ref } from 'vue'
import { resolveColorByIndex } from '@/plugins/colors'

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

const color = ref<string | null>()

if (props.tag.type === TagType.ColorValue) {
  if (typeof props.tag.value === 'number') {
    color.value = resolveColorByIndex(props.tag.value) ?? null
  }
}
</script>

<style lang="scss">
.SearchTag {
  cursor: default;

  & > div > span {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  & > div > span {
    border-top-width: 2px !important;
    border-top-style: solid;
    border-bottom-width: 2px !important;
    border-bottom-style: solid;
  }

  & > div > span:first-child {
    border-left-width: 2px !important;
    border-left-style: solid;
    border-top-left-radius: var(--border-radius) !important;
    border-bottom-left-radius: var(--border-radius) !important;
  }

  & > div > span:last-child {
    border-right-width: 2px !important;
    border-right-style: solid;
    border-top-right-radius: var(--border-radius) !important;
    border-bottom-right-radius: var(--border-radius) !important;
  }
}
</style>
