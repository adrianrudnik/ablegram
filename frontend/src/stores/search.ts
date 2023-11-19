import { defineStore, storeToRefs } from 'pinia'
import { computed, ref, watch } from 'vue'
import type { SearchQuery } from '@/plugins/search/query'
import type { SearchResult } from '@/plugins/search/result'
import { useStatStore } from '@/stores/stats'
import { useSearchResultStore } from '@/stores/results'
import { compileFilter, useActiveFiltersStore } from '@/stores/search-filters'
import { createQueryInstanceWithDefaults } from '@/plugins/search/query'

export const useSearchStore = defineStore('search', () => {
  // Holds the user given query string
  const currentQueryString = ref('')

  // Holds the last query that was executed.
  // Is used to load more results, with offset.
  const currentQueryInstance = ref<SearchQuery | undefined>()

  // Holds the last sort key that was used.
  // Is used to load more results based on the search_after mechanic, by the InfiniteScrollTrigger component.
  // https://www.elastic.co/guide/en/elasticsearch/reference/current/paginate-search-results.html#search-after
  const lastSortKey = ref<string[] | undefined>([])

  // Holds the total result count of the current query
  const totalHits = ref(0)

  // Used to determine if the search was used before. If not, we can show intro stuff and hide the hit count.
  const isClean = ref(true)

  // Used to determine if there are results to show, or even available
  const hasResults = computed(() => currentQueryInstance.value !== undefined && totalHits.value > 0)

  const resultViewMode = ref<'elements' | 'files'>('elements')

  const { isSearching } = storeToRefs(useStatStore())

  const activeFilters = computed(() => useActiveFiltersStore().all)

  // Install a watcher that will retrigger when a tag is added to the active filters
  watch(activeFilters.value, async () => {
    // Also ensure we reset the lastSortKey to begin a fresh result
    lastSortKey.value = []
    await search(currentQueryInstance.value ?? createQueryInstanceWithDefaults(), true)
  })

  const search = async (
    query: SearchQuery,
    storeResults: boolean = true
  ): Promise<SearchResult | null> => {
    // Catch a reset by the user
    if (currentQueryString.value.trim() === '' && activeFilters.value.length == 0) {
      reset()
      return null
    }

    isSearching.value = true
    isClean.value = false

    // If the view mode is files, we need to add a filter to the query.
    let actualQueryString = ''

    if (resultViewMode.value === 'files') {
      actualQueryString = query.query.query + ' +tags:"type:file"'
    } else {
      actualQueryString = query.query.query + ' -tags:"type:file"'
    }

    // If we have active filters, we need to add them to the query.
    if (activeFilters.value.length > 0) {
      actualQueryString += ' ' + activeFilters.value.map((f) => compileFilter(f)).join(' ')
    }

    try {
      const response = await fetch(import.meta.env.VITE_API_URL + '/search/query', {
        credentials: 'include',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          ...query,
          query: {
            ...query.query,
            query: actualQueryString
          }
        }),
        method: 'POST'
      })

      if (!response.ok) {
        throw new Error(response.statusText)
      }

      const result: SearchResult = await response.json()

      // Parse the results into the store if requested.
      // loadMore for example has its own logic, so this is optional.
      if (storeResults) {
        useSearchResultStore().overwrite(
          result.hits.map((h) => {
            h.fields.id = h.id
            return h.fields
          })
        )
      }

      // Remember the query, in case we need to load more
      currentQueryInstance.value = query

      // Total hits are valuable for the user interface
      totalHits.value = result.total_hits

      // Get the last result to retrieve the sort offset
      // This helps to later load more results based on search_after feature.
      const last = result.hits[result.hits.length - 1]
      if (last) {
        lastSortKey.value = [last.score.toString(), last.id]
      }

      return result
    } finally {
      isSearching.value = false
    }
  }

  const resetLoadMore = () => {
    lastSortKey.value = []
  }

  const loadMore = async () => {
    if (currentQueryInstance.value === undefined) {
      return
    }

    currentQueryInstance.value.search_after = lastSortKey.value

    const result = await search(currentQueryInstance.value, false)

    // Append the results to the store
    if (result !== null) {
      useSearchResultStore().updateBatch(
        result.hits.map((h) => {
          h.fields.id = h.id
          return h.fields
        })
      )
    }
  }

  const reset = () => {
    useSearchResultStore().clear()
    useActiveFiltersStore().clear()
    currentQueryString.value = ''
    currentQueryInstance.value = undefined
    totalHits.value = 0
    isClean.value = true
    lastSortKey.value = []
  }

  return {
    currentQueryString,
    currentQueryInstance,
    lastSortKey,
    hasResults,
    isClean,
    totalHits,
    search,
    loadMore,
    reset,
    resetLoadMore,
    resultViewMode
  }
})
