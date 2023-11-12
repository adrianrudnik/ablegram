<template>
  <AppStatus class="mb-3" />

  <SectionHeadline :title="t('settings-view.title')">
    <template #description>
      <p>{{ t('settings-view.description') }}</p>
    </template>

    <LogSettings />

    <BehaviourSettings />
  </SectionHeadline>
</template>

<script setup lang="ts">
import LogSettings from '@/components/settings/LogSettings.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import { useI18n } from 'vue-i18n'
import BehaviourSettings from '@/components/settings/BehaviourSettings.vue'
import AppStatus from '@/components/parts/AppStatus.vue'

const { t } = useI18n()
</script>

<style scoped>
.p-card {
  margin-bottom: 1rem;
}
</style>
