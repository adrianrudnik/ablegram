import { defineStore, storeToRefs } from 'pinia'
import { computed, ref } from 'vue'
import type { SearchQuery } from '@/plugins/search/query'
import type { SearchResult } from '@/plugins/search/result'
import { useStatStore } from '@/stores/stats'
import { useSearchResultStore } from '@/stores/results'

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

  const isClean = ref(true)
  // Used to determine if there are results to show, or even available
  const hasResults = computed(() => currentQueryInstance.value !== undefined && totalHits.value > 0)

  const { isSearching } = storeToRefs(useStatStore())

  const search = async (
    query: SearchQuery,
    storeResults: boolean = true
  ): Promise<SearchResult | null> => {
    // Catch a reset by the user
    if (currentQueryString.value.trim() === '') {
      reset()
      return null
    }

    isSearching.value = true
    isClean.value = false

    try {
      const response = await fetch(import.meta.env.VITE_API_URL + '/search/query', {
        credentials: 'include',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(query),
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
    reset
  }
})
