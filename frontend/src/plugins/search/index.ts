import type { SearchQuery } from './query'
import type { SearchResult } from './result'
import { useStatStore } from '@/stores/stats'

export async function executeQuerySearch(query: SearchQuery): Promise<SearchResult> {
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

  return result
}
