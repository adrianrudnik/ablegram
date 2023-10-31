<template>
  <div class="QueryInput">
    <div class="p-inputgroup flex-1">
      <AutoComplete
        v-model="currentSelection"
        :forceSelection="false"
        :panel-class="{ HidePanel: hidePanel }"
        :pt="{ token: { class: 'SearchTagChip' } }"
        :suggestions="suggestions"
        :virtualScrollerOptions="{ itemSize: 50, scrollWidth: '100vw', scrollHeight: '300px' }"
        multiple
        :placeholder="t('query-input-component.placeholder')"
        @complete="aComplete"
        @item-select="clearAfterSelect"
        @clear="clearInput"
        :class="{ 'p-invalid': !currentRequestValid }"
      >
        <template #option="slotProps">
          <div class="flex align-options-center">
            <SearchTag :tag="slotProps.option" />
          </div>
        </template>

        <template #chip="slotProps">
          <SearchTag :tag="slotProps.value" />
        </template>

        <template #removetokenicon="slotProps">
          <div
            @click="slotProps.onClick"
            class="RemoveItem inline-flex bg-red-500 align-items-center px-1 text-white cursor-pointer"
          >
            <i class="pi pi-times"></i>
          </div>
        </template>
      </AutoComplete>
    </div>

    <div class="Options flex justify-content-between gap-2 mt-2 mb-3">
      <div class="my-2 text flex gap-2">
        <span>{{ t('query-input-component.hits', { count: currentResultCount }) }}</span>
        <span class="text-red-500" v-if="!currentRequestValid">
          {{ t('query-input-component.invalid-query') }}
        </span>
      </div>

      <div class="flex justify-content-end gap-2">
        <Button
          @click="showHelp = !showHelp"
          size="small"
          :label="
            !showHelp
              ? t('query-input-component.action.show-examples')
              : t('query-input-component.action.hide-examples')
          "
          :plain="!showHelp"
          :text="!showHelp"
        />

        <Button
          @click="clearInput"
          plain
          text
          size="small"
          :label="t('query-input-component.action.reset-search')"
        />
      </div>
    </div>

    <SearchExamples v-if="showHelp" />
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import type { Tag } from '@/stores/tags'
import { useTagStore } from '@/stores/tags'
import type { AutoCompleteCompleteEvent } from 'primevue/autocomplete'
import AutoComplete from 'primevue/autocomplete'
import SearchTag from '@/components/structure/SearchTag.vue'
import { useSearchResultStore } from '@/stores/results'
import { useStatStore } from '@/stores/stats'
import { useI18n } from 'vue-i18n'
import { watchDebounced } from '@vueuse/core'
import { executeQuerySearch } from '@/plugins/search'
import Button from 'primevue/button'
import SearchExamples from '@/components/parts/search/SearchExamples.vue'

const { t } = useI18n()
const currentSelection = ref<Tag[]>([])
const suggestions = ref<Tag[]>([])
const currentPlainValue = ref('')
const hidePanel = ref(false)

const showHelp = ref(false)

const resultStore = useSearchResultStore()
const statStore = useStatStore()

const aComplete = (event: AutoCompleteCompleteEvent) => {
  currentPlainValue.value = event.query
  suggestions.value = useTagStore().entries.filter(
    (entry) =>
      entry.trans.plain?.toLowerCase().includes(event.query.toLowerCase()) ||
      entry.id.includes(event.query.toLowerCase())
  )

  // Hide the panel if no suggestions have been found
  hidePanel.value = suggestions.value.length === 0
}

const clearAfterSelect = () => {
  currentPlainValue.value = ''
}

const clearInput = () => {
  currentSelection.value = []
  currentPlainValue.value = ''
  currentResultCount.value = 0
  currentRequestValid.value = true
  resultStore.clear()
}

const currentResultCount = ref(0)
const currentRequestValid = ref(true)

const query = computed(() => {
  const plain = currentPlainValue.value
  const tags = currentSelection.value.map((tag) => 'tags:"' + tag.id + '"')

  return (plain + ' ' + tags.join('')).trim()
})

watchDebounced(
  query,
  async () => {
    if (query.value.trim() === '') return

    statStore.isSearching = true

    try {
      const result = await executeQuerySearch({
        size: 10,
        query: {
          query: query.value
        },

        fields: ['*']
      })

      resultStore.overwrite(
        result.hits.map((h) => {
          h.fields.id = h.id
          return h.fields
        })
      )

      currentResultCount.value = result.total_hits
      currentRequestValid.value = true
    } catch {
      currentRequestValid.value = false
    } finally {
      statStore.isSearching = false
    }
  },
  { debounce: 200, maxWait: 500 }
)
</script>

<style lang="scss">
.SearchTagChip {
  border: 0;
  padding: 0;
  margin: 0;

  display: flex;
  align-items: stretch;
  justify-content: center;

  border-top-width: 2px !important;
  border-top-style: solid;
  border-bottom-width: 2px !important;
  border-bottom-style: solid;

  border-left-width: 2px !important;
  border-left-style: solid;
  border-right-width: 2px !important;
  border-right-style: solid;
  border-radius: var(--border-radius);

  .RemoveItem {
    border-right-width: 0 !important;
    border-right-style: solid !important;
    border-bottom-right-radius: calc(var(--border-radius) / 1.6) !important;
    border-top-right-radius: calc(var(--border-radius) / 1.6) !important;
  }

  .SearchTag {
    margin: unset !important;

    .Parts {
      span {
        border: unset !important;
      }

      span:first-child {
        border-left-width: 2px !important;
        border-left-style: solid;
        border-top-left-radius: calc(var(--border-radius) / 1.6) !important;
        border-bottom-left-radius: calc(var(--border-radius) / 1.6) !important;
      }

      span:last-child {
        border-right-width: unset !important;
        border-right-style: unset !important;
        border-top-right-radius: unset !important;
        border-bottom-right-radius: unset !important;
      }
    }
  }
}

.p-autocomplete-panel.HidePanel {
  visibility: hidden;
}

.QueryInput {
  input {
    caret-shape: block;
  }

  .Options {
    a {
      color: var(--gray-500);
      text-decoration: underline;

      &:hover {
        color: var(--gray-700);
        text-decoration: none;
      }
    }
  }
}
</style>
