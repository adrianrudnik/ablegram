<template>
  <div class="ActiveFilterBar" v-if="activeFilters.entries.length > 0">
    <div class="flex flex-row flex-wrap">
      <div
        class="Filter"
        v-for="filter in activeFilters.entries"
        :key="filter.id"
        :id="'af_' + createIdFrom(filter.id)"
      >
        <Dropdown
          v-model="filter.mode"
          @change="updateActiveFilterMode(filter, $event.value)"
          :options="searchModes"
          option-label="name"
          option-value="mode"
          :append-to="'#af_' + createIdFrom(filter.id)"
        />
        <div class="Content">
          <component :is="resolveComponent(filter.type)" :value="filter.content" />
        </div>
        <Button @click="removeActiveFilter(filter)" size="small" label="X" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ActiveFilter } from '@/stores/search-filters'
import { ActiveFilterMode, ActiveFilterType, useActiveFiltersStore } from '@/stores/search-filters'
import Button from 'primevue/button'
import ActiveFilterTagEntry from '@/components/search/ActiveFilterTagEntry.vue'
import Dropdown from 'primevue/dropdown'
import { createIdFrom } from '@/plugins/id'

const searchModes = [
  { name: 'prefer', mode: ActiveFilterMode.SHOULD },
  { name: 'must', mode: ActiveFilterMode.MUST },
  { name: 'not', mode: ActiveFilterMode.NOT }
]

const activeFilters = useActiveFiltersStore()

const updateActiveFilterMode = (v: ActiveFilter, mode: ActiveFilterMode) => {
  activeFilters.update({ ...v, mode: mode })
}

const removeActiveFilter = (v: ActiveFilter) => activeFilters.remove(v)

function resolveComponent(type: ActiveFilterType): any {
  switch (type) {
    case ActiveFilterType.TAG:
      return ActiveFilterTagEntry
    default:
      return ActiveFilterTagEntry
  }
}
</script>

<style lang="scss">
.Filter {
  display: flex;
  flex-wrap: nowrap;

  zoom: 0.8;

  margin-right: 0.4rem;
  margin-bottom: 0.4rem;

  .p-dropdown {
    border-width: 2px;
    border-right: 0;
    border-color: black;
    border-radius: var(--border-radius) 0 0 var(--border-radius);
  }

  .p-button {
    border-radius: 0 var(--border-radius) var(--border-radius) 0;
  }

  .Content {
    display: flex;
    align-items: center;

    border-right: 0;
    border-left: 0;
    border-top: 2px solid black;
    border-bottom: 2px solid black;

    padding-right: 0.5rem;
  }
}
</style>
