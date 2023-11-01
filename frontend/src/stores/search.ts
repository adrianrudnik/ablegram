import {defineStore} from "pinia";
import {computed, ref, watch} from "vue";
import type {SearchQuery} from "@/plugins/search/query";
import type {SearchResult} from "@/plugins/search/result";
import {useStatStore} from "@/stores/stats";
import {useSearchResultStore} from "@/stores/results";

export const useSearchStore = defineStore('search', () => {
  const currentQuery = ref<SearchQuery|undefined>()
  const currentOffset = ref(0)
  const lastSortKey = ref<string[]|undefined>([])

  const executeQuerySearch = async (query: SearchQuery): Promise<SearchResult>=>  {
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

    // Extract global result information
    useStatStore().searchResultCount = result.total_hits

    // Remember the query, in case we need to load more
    currentQuery.value = query

    // Get the last result to retrieve the sort offset
    const last = result.hits[result.hits.length - 1]
    if (last) {
      lastSortKey.value = [last.score.toString(), last.id]
    }

    return result
  }

  const clear = () => {
    useSearchResultStore().clear()
    useStatStore().searchResultCount = 0
    currentQuery.value = undefined
    currentOffset.value = 0
    lastSortKey.value = []
  }

  return {
    clear,
    currentQuery,
    currentOffset,
    lastSortKey,
    executeQuerySearch
  }
})

