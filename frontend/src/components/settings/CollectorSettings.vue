<template>
  <SectionHeadline :title="t('collector-settings.title')">
    <template #description>
      <p>{{ t('collector-settings.description') }}</p>
    </template>

    <CollectorTargetCard
      :target="target"
      v-for="target in current.collector.targets"
      :key="target.id"
      class="surface-200 mb-3"
    />
  </SectionHeadline>
</template>

<script setup lang="ts">
import { useConfigStore } from '@/stores/config'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import CollectorTargetCard from '@/components/settings/CollectorTargetCard.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'

const { t } = useI18n()

const { current } = storeToRefs(useConfigStore())
</script>
