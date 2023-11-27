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
      v-if="isAdmin && dialogRef.data.result.pathFolder"
      @click="openLocalPath(dialogRef.data.result.pathFolder)"
      :label="t('common.button.open-folder')"
    />

    <Button
      outlined
      icon="pi pi-file-o"
      v-if="isAdmin && dialogRef.data.result.pathAbsolute"
      @click="openLocalPath(dialogRef.data.result.pathAbsolute)"
      :label="t('common.button.open-file')"
    />

    <Button
      outlined
      icon="pi pi-heart"
      v-if="isGuest && dialogRef.data.result.id"
      @click="suggestSearchResult(dialogRef.data.result.id)"
      :label="t('search-result-focus-dialog.actions.suggest')"
    />
  </div>
</template>

<script setup lang="ts">
import type { Component } from 'vue'
import { inject } from 'vue'
import type { HitFieldset } from '@/plugins/search/result'
import Button from 'primevue/button'
import { openLocalPath, suggestSearchResult } from '@/plugins/api'
import { useI18n } from 'vue-i18n'
import { useSessionStore } from '@/stores/session'

const { t } = useI18n()

const { isAdmin, isGuest } = useSessionStore()

const dialogRef = inject('dialogRef') as {
  data: {
    component: Component
    result: HitFieldset
    expanded: boolean
  }
}
</script>
