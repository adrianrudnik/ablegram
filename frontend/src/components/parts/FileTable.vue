<template>
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
    :rows="15"
    :resizable-columns="true"
    column-resize-mode="fit"
    :show-gridlines="false"
    class="FileTable p-datatable-sm"
  >
    <template #header>
      <div class="flex justify-content-between align-items-center">
        <span class="">{{ t('file-table.header.hits.label', { count: fileCount }) }}</span>
        <span class="p-input-icon-left">
          <i class="pi pi-search" />
          <InputText
            type="search"
            v-model="filters['global'].value"
            :placeholder="t('file-table.header.search.placeholder')"
          />
        </span>
      </div>
    </template>

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

        <span v-else-if="slotProps.data.status === 'failed'" v-tooltip.left="slotProps.data.remark">
          <i class="pi pi-fw pi-times-circle text-red-500"></i>
          {{ t('file-table.columns.status.state.failed') }}
        </span>
      </template>
    </Column>
    <Column
      field="filename"
      :sortable="true"
      :header="t('file-table.columns.filename.label')"
    ></Column>
    <Column field="folder" :sortable="true" :header="t('file-table.columns.folder.label')"></Column>

    <template #empty>
      <span v-if="filters['global'].value == ''">
        {{ t('file-table.errors.no-files-found') }}
      </span>

      <span v-else>
        {{
          t('file-table.errors.no-files-found-for-filter', {
            filter: filters['global'].value
          })
        }}
      </span>
    </template>
  </DataTable>
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

const { t } = useI18n()

const defaultSort: DataTableSortMeta[] = [{ field: 'status', order: 1 }]

const filters = ref({
  global: { value: null, matchMode: FilterMatchMode.CONTAINS }
})

const filterFields = ['path', 'status']

const files = useFilesStore().entries
const fileCount = computed(() => useFilesStore().count)
</script>
