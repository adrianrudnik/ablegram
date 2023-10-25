<template>
  <div class="FileTable">
    <SectionHeadline>
      <template #title>
        {{ t('file-table.title') }}
        <Tag severity="info" :value="fileCount"></Tag>
      </template>

      <template #description>
        {{ t('file-table.description') }}
      </template>

      <template #right>
        <i class="pi pi-search"></i>
        <InputText
          v-model="filters['global'].value"
          :placeholder="t('file-table.global-search.placeholder')"
        />
      </template>
    </SectionHeadline>

    <DataTable
      v-model:filters="filters"
      :removable-sort="true"
      stateStorage="session"
      stateKey="dt-file-table"
      :value="files"
      :multi-sort-meta="defaultSort"
      :global-filter-fields="filterFields"
      sortMode="multiple"
      tableStyle="min-width: 50rem"
      paginator
      :rows="10"
      :resizable-columns="true"
      column-resize-mode="fit"
      :show-gridlines="true"
      class="p-datatable-sm"
    >
      <Column
        field="filename"
        :sortable="true"
        :header="t('file-table.columns.filename.label')"
      ></Column>
      <Column
        field="folder"
        :sortable="true"
        :header="t('file-table.columns.folder.label')"
      ></Column>
      <Column field="status" :sortable="true" :header="t('file-table.columns.status.label')">
        <template #body="slotProps">
          <span v-if="slotProps.data.status === 'pending'">
            <i class="pi pi-fw pi-spin pi-spinner"></i>
            {{ t('file-table.columns.status.state.pending') }}
          </span>

          <span v-else-if="slotProps.data.status === 'processed'">
            <i class="pi pi-fw pi-check-circle text-green-500"></i>
            {{ t('file-table.columns.status.state.processed') }}
          </span>

          <span v-else-if="slotProps.data.status === 'failed'">
            <i class="pi pi-fw pi-times-circle text-red-500"></i>
            {{ t('file-table.columns.status.state.failed') }}
          </span>
        </template>
      </Column>
    </DataTable>
  </div>
</template>

<script lang="ts" setup>
import { useFilesStore } from '@/stores/files'
import type { DataTableSortMeta } from 'primevue/datatable'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import { FilterMatchMode } from 'primevue/api'
import InputText from 'primevue/inputtext'
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import Tag from 'primevue/tag'

const { t } = useI18n()

const defaultSort: DataTableSortMeta[] = [
  { field: 'folder', order: 1 },
  { field: 'filename', order: 1 }
]

const filters = ref({
  global: { value: null, matchMode: FilterMatchMode.CONTAINS }
})

const filterFields = ['path', 'status']

const files = useFilesStore().entries
const fileCount = computed(() => useFilesStore().entries.length)
</script>
