<template>
  <AppStatus class="mb-3" />

  <SectionHeadline :title="t('settings-view.title')">
    <template #description>
      <p>{{ t('settings-view.description') }}</p>

      <Message severity="warn" :closable="false">
        {{ t('settings-view.restart-notice') }}
      </Message>
    </template>

    <CollectorSettings />
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
import CollectorSettings from '@/components/settings/CollectorSettings.vue'
import Message from 'primevue/message'
import { useConfigStore } from '@/stores/config'

const { t } = useI18n()

// Ensure the config is loaded
useConfigStore().load()
</script>

<style scoped>
.p-card {
  margin-bottom: 1rem;
}
</style>
