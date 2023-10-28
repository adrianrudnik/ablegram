<template>
  <div>
    <div class="p-inputgroup flex-1">
      <span class="p-inputgroup-addon">
        <i class="pi pi-search"></i>
      </span>
      <InputText
        @update:model-value="search"
        type="search"
        v-model="query"
        class="w-full"
        :placeholder="t('query-input-component.placeholder')"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import InputText from 'primevue/inputtext'
import { useI18n } from 'vue-i18n'
import { executeQuerySearch } from '@/plugins/search'
import { useSearchResultStore } from '@/stores/results'
import { useStatStore } from '@/stores/stats'

const { t } = useI18n()

const search = async (query: string) => {
  useStatStore().isSearching = true
  try {
    const result = await executeQuerySearch({
      size: 10,
      query: {
        query: query
      },
      fields: ['*']
    })

    useSearchResultStore().overwrite(
      result.hits.map((h) => {
        h.fields.id = h.id
        return h.fields
      })
    )
  } finally {
    useStatStore().isSearching = false
  }
}

const query = ref('')
</script>
