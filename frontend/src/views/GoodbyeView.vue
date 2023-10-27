<template>
  <SectionHeadline :title="t('goodbye-view.title')">
    <template #description>
      {{ t('goodbye-view.description') }}
    </template>
  </SectionHeadline>

  <SectionHeadline :title="t('donate.title.long')">
    <template #description>
      <p>{{ t('donate.description') }}</p>
      <p>{{ t('donate.usage') }}</p>

      <ul class="pl-0 list-none">
        <li>
          <strong>{{ t('donate.processors.patreon.label') }}: </strong>
          <a :href="t('donate.processors.patreon.link')" target="_blank" rel="noopener noreferrer">
            {{ t('donate.processors.patreon.link') }}
          </a>
        </li>
      </ul>
    </template>
  </SectionHeadline>
</template>

<script setup lang="ts">
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>
