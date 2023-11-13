<template>
  <component
    :is="dialogRef.data.component"
    :result="dialogRef.data.result"
    :expanded="dialogRef.data.expanded"
    class="mb-3"
  />
  <div class="flex mb-3 gap-2">
    <Button
      outlined
      icon="pi pi pi-folder-open"
      v-if="dialogRef.data.result.pathFolder"
      @click="openLocalPath(dialogRef.data.result.pathFolder)"
      :label="t('common.button.open-folder')"
    />

    <Button
      outlined
      icon="pi pi-file-o"
      v-if="dialogRef.data.result.pathAbsolute"
      @click="openLocalPath(dialogRef.data.result.pathAbsolute)"
      :label="t('common.button.open-file')"
    />
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue'
import type { Component } from 'vue'
import type { HitFieldset } from '@/plugins/search/result'
import Button from 'primevue/button'
import { openLocalPath } from '@/plugins/api'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const dialogRef = inject('dialogRef') as {
  data: {
    component: Component
    result: HitFieldset
    expanded: boolean
  }
}
</script>
