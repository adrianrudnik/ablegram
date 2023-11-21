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

    <Button label="Filesystem target" icon="pi pi-plus" @click="addFilesystemTarget" />
  </SectionHeadline>
</template>

<script setup lang="ts">
import { useConfigStore } from '@/stores/config'
import { useI18n } from 'vue-i18n'
import { storeToRefs } from 'pinia'
import CollectorTargetCard from '@/components/settings/CollectorTargetCard.vue'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import Button from 'primevue/button'
import CollectorTargetForm from '@/components/settings/CollectorTargetForm.vue'
import { useDialog } from 'primevue/usedialog'

const { t } = useI18n()

const { current } = storeToRefs(useConfigStore())

const dialog = useDialog()

const addFilesystemTarget = () => {
  dialog.open(CollectorTargetForm, {
    props: {
      header: t('collector-target-form.title')
    }
  })
}
</script>
