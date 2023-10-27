<template>
  <SectionHeadline>
    <template #title>
      {{ t('file-view.title') }}
    </template>

    <template #description>
      {{ t('file-view.description') }}
    </template>
  </SectionHeadline>

  <FileTable class="pb-3" />
</template>

<script setup lang="ts">
import FileTable from '@/components/parts/FileTable.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>
