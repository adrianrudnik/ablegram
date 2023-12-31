<template>
  <Card class="surface-200">
    <template #content>
      <div class="flex flex-row align-items-center">
        <i class="pr-3 pi pi-database"></i>
        <div class="flex flex-column flex-grow-1 pr-2">
          <div class="mb-2">{{ props.target.id }}</div>
          <div class="font-mono text-xs text-break-all">{{ props.target.uri }}</div>
        </div>
        <div class="zoom-80 flex-nowrap flex gap-2">
          <Button icon="pi pi-pencil" aria-label="Change" @click="editFilesystemTarget" />
          <Button
            icon="pi pi-trash"
            severity="danger"
            aria-label="Remove"
            @click="confirmDelete($event)"
          />
        </div>
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import Card from 'primevue/card'
import Button from 'primevue/button'
import type { CollectorTargetConfig } from '@/stores/config'
import { useConfirm } from 'primevue/useconfirm'
import { useI18n } from 'vue-i18n'
import { fetchApi } from '@/plugins/api'
import { useConfigStore } from '@/stores/config'
import { useDialog } from 'primevue/usedialog'
import CollectorTargetForm from '@/components/settings/CollectorTargetForm.vue'

const { t } = useI18n()

const configStore = useConfigStore()

const props = defineProps<{
  target: CollectorTargetConfig
}>()

const dialog = useDialog()
const confirm = useConfirm()

const editFilesystemTarget = () => {
  dialog.open(CollectorTargetForm, {
    props: {
      header: t('collector-target-form.title')
    },
    data: {
      target: props.target
    }
  })
}

const confirmDelete = (event: Event) => {
  confirm.require({
    target: event.currentTarget as HTMLElement,
    header: t('collector-target-card.confirm-delete.header'),
    message: t('collector-target-card.confirm-delete.message', { id: props.target.id }),
    icon: 'pi pi-exclamation-triangle',
    acceptClass: 'p-button-danger',
    accept: async () => {
      await fetchApi(`/api/config/collector/targets/${props.target.id}`, {
        method: 'DELETE'
      })

      await configStore.load()
    }
  })
}
</script>
