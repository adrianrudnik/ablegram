<template>
  <SectionHeadline :title="t('goodbye-view.title')">
    <template #description>
      {{ t('goodbye-view.description') }}
    </template>
  </SectionHeadline>

  <SectionHeadline>
    <template #title>
      {{ t('donate.title.long') }}
    </template>
    <template #description>
      <p>{{ t('donate.description') }}</p>
      <p>{{ t('donate.usage') }}</p>

      <i class="pi pi-heart font-bold text-red-500 text-xl"></i>

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
