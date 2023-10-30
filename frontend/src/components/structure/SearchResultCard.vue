<template>
  <div class="surface-card shadow-2 border-round border-top border-1">
    <div v-if="props.header" class="bg-black-alpha-90 p-2 text-white font-semibold">
      {{ header }}
    </div>

    <div class="p-2">
      <div class="text-xl font-medium text-900 mb-3" v-if="title">{{ title }}</div>

      <div class="mb-2">
        <slot />
      </div>

      <div v-if="!!$slots.actions" class="flex gap-2">
        <slot name="actions" />

        <Button
          outlined
          icon="pi pi-tags"
          @click="showTags = true"
          v-if="props.tags && !showTags"
          size="small"
          :label="t('result-item-component.button.show-tags')"
        />
        <Button
          icon="pi pi-tags"
          @click="showTags = false"
          v-if="props.tags && showTags"
          size="small"
          :label="t('result-item-component.button.hide-tags')"
        />
      </div>

      <TagRow v-if="showTags && props.tags" :tags="props.tags" class="text-sm mt-3" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import TagRow from '@/components/parts/search/TagRow.vue'
import Button from 'primevue/button'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = defineProps<{
  header?: string
  title?: string
  tags?: string[]
}>()

const showTags = ref(false)
</script>
