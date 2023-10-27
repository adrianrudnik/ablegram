<template>
  <SectionHeadline :title="t('goodbye-view.title')">
    <template #description>
      {{ t('goodbye-view.description') }}
    </template>
  </SectionHeadline>
</template>

<script setup lang="ts">
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>
