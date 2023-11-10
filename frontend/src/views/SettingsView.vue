<template>
  <SectionHeadline :title="t('settings-view.title')">
    <template #description>
      <p>{{ t('settings-view.description') }}</p>
    </template>
  </SectionHeadline>

  <LogSettings />
</template>

<script setup lang="ts">
import LogSettings from '@/components/settings/LogSettings.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>
