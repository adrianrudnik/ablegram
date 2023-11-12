<template>
  <Menubar :model="items" class="w-full" />

  <SearchExamples v-if="showExamples" class="mt-3" />

  <Message severity="info" v-if="resultViewMode === 'files'" :closable="false">
    {{ t('search-query-toolbar.file-search-notice') }}
  </Message>
</template>

<script setup lang="ts">
import Menubar from 'primevue/menubar'
import { computed, ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useSearchStore } from '@/stores/search'
import { useI18n } from 'vue-i18n'
import Message from 'primevue/message'
import SearchExamples from '@/components/search/SearchExamples.vue'

const { t } = useI18n()
const { resultViewMode, currentQueryInstance } = storeToRefs(useSearchStore())
const { search, resetLoadMore } = useSearchStore()

const showExamples = ref(false)

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
  },
  {
    label: showExamples.value ? 'Hide examples' : 'Show examples',
    icon: 'pi pi-question',
    command: () => (showExamples.value = !showExamples.value)
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
