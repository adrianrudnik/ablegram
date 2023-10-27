<template>
  <div class="SearchTag mb-1 mr-1 inline-block">
    <div class="flex align-items-center gap-0">
      <div v-if="showRealm" class="p-1 px-2 bg-white border-900">
        {{ parts.realm }}
      </div>
      <div class="p-1 px-2 text-base bg-black-alpha-10 border-900">
        {{ parts.topic }}
      </div>
      <div v-if="parts.detail" class="p-1 px-2 bg-black-alpha-90 text-white border-900">
        {{ parts.detail }}
      </div>
      <div v-if="parts.extra" class="p-1 px-2 bg-black-alpha-50 text-white border-900">
        {{ parts.extra }}
      </div>
      <div
        v-if="showCount"
        class="p-1 px-2 bg-black-alpha-30 border-900"
        v-tooltip.top="t('search-tag-component.count.label', { count: props.tag.count })"
      >
        {{ props.tag.count }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Tag as TagInterface } from '@/stores/tags'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import messages from '@intlify/unplugin-vue-i18n/messages'

const { t } = useI18n()

const props = withDefaults(
  defineProps<{
    tag: TagInterface
    showRealm?: boolean
    showCount?: boolean
  }>(),
  {
    showRealm: false,
    showCount: false
  }
)

const translate = (prefix: string, value: string | number | null): string | null => {
  if (value === null) {
    return null
  }
  if (typeof value === 'number') {
    return value.toString()
  }

  if (!Number.isNaN(Number(value))) {
    return value.toString()
  }

  return t(prefix + value)
}

const parts = computed(() => {
  const p = props.tag.id.split(':')
  return {
    realm: translate('tags.', p[0] ?? null),
    topic: translate('tags.' + p[0] + ':', p[1] ?? null),
    detail: translate('tags.' + p[0] + ':' + p[1] + ':', p[2] ?? null),
    extra: p[3]
  }
})
</script>

<style lang="scss">
.SearchTag {
  cursor: default;

  & > div > div {
    border-top-width: 2px !important;
    border-top-style: solid;
    border-bottom-width: 2px !important;
    border-bottom-style: solid;
  }

  & > div > div:first-child {
    border-left-width: 2px !important;
    border-left-style: solid;
    border-top-left-radius: var(--border-radius) !important;
    border-bottom-left-radius: var(--border-radius) !important;
  }

  & > div > div:last-child {
    border-right-width: 2px !important;
    border-right-style: solid;
    border-top-right-radius: var(--border-radius) !important;
    border-bottom-right-radius: var(--border-radius) !important;
  }
}
</style>
