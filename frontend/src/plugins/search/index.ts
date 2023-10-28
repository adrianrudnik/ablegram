import type { SearchQuery } from './query'
import type { SearchResult } from './result'

export async function executeQuerySearch(query: SearchQuery): Promise<SearchResult> {
  const r = await fetch(import.meta.env.VITE_API_URL + '/search/query', {
    credentials: 'include',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(query),
    method: 'POST'
  })

  if (!r.ok) {
    throw new Error(r.statusText)
  }

  return r.json()
}
