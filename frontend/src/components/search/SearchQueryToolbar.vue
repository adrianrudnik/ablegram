<template>
  <Menubar :model="items" class="w-full" />

  <Message severity="info" v-if="resultViewMode === 'files'" :closable="false">
    {{ t('search-query-toolbar.file-search-notice') }}
  </Message>
</template>

<script setup lang="ts">
import Menubar from 'primevue/menubar'
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useSearchStore } from '@/stores/search'
import { useI18n } from 'vue-i18n'
import Message from 'primevue/message'

const { t } = useI18n()
const { resultViewMode, currentQueryInstance } = storeToRefs(useSearchStore())
const { search, resetLoadMore } = useSearchStore()

const switchViewMode = (mode: 'elements' | 'files') => {
  if (mode !== resultViewMode.value) {
    // Switch mode
    resultViewMode.value = mode

    // Redo the search
    if (currentQueryInstance.value !== undefined) {
      resetLoadMore()
      search(currentQueryInstance.value)
    }
  }
}

const items = computed(() => [
  {
    label: 'View',
    icon: 'pi pi-search',
    items: [
      {
        label: 'Results',
        icon: 'pi pi-list',
        items: [
          {
            label: 'Single elements',
            icon: resultViewMode.value === 'elements' ? 'pi pi-circle-on' : 'pi pi-circle-off',
            command: () => switchViewMode('elements')
          },
          {
            label: 'Grouped by file',
            icon: resultViewMode.value === 'files' ? 'pi pi-circle-on' : 'pi pi-circle-off',
            command: () => switchViewMode('files')
          }
        ]
      }
    ]
  }
])
</script>

<style scoped lang="scss">
.p-menubar {
  background-color: white;
  border: none;
  padding: 0;
  margin-top: 0.4rem;
}
</style>
