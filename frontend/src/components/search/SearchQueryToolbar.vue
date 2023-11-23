<template>
  <div class="SearchQueryToolbar">
    <MegaMenu :model="items" class="w-full" breakpoint="320px" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useSearchStore } from '@/stores/search'
import MegaMenu from 'primevue/megamenu'
import { addRandomTagFilter, useActiveFiltersStore } from '@/stores/search-filters'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const { reset: resetSearch } = useSearchStore()
const { clear: resetFilters } = useActiveFiltersStore()

const items = ref([
  {
    label: t('search-query-toolbar.tools.header'),
    icon: 'pi pi-sliders-v',
    items: [
      [
        {
          label: t('search-query-toolbar.tools.tags.header'),
          items: [
            {
              label: t('search-query-toolbar.tools.tags.add-random'),
              icon: 'pi pi-plus',
              command: () => addRandomTagFilter()
            }
          ]
        },
        {
          label: t('search-query-toolbar.tools.filters.header'),
          items: [
            {
              label: t('search-query-toolbar.tools.filters.clear-all'),
              icon: 'pi pi-filter-slash',
              command: () => resetFilters()
            }
          ]
        },
        {
          label: t('search-query-toolbar.tools.search.header'),
          items: [
            {
              label: t('search-query-toolbar.tools.search.clear'),
              icon: 'pi pi-trash',
              command: () => resetSearch()
            }
          ]
        }
      ]
    ]
  }
])
</script>

<style lang="scss">
.SearchQueryToolbar {
  zoom: 0.9;
  margin-left: 8px;
  margin-right: 8px;

  .p-megamenu {
    border: 0;
    background-color: var(--gray-200);
    border-top-left-radius: 0;
    border-top-right-radius: 0;

    .p-megamenu-root-list > .p-menuitem > .p-menuitem-content > .p-menuitem-link {
      padding: 0.4rem !important;
      margin: 0;
    }
  }
}
</style>
