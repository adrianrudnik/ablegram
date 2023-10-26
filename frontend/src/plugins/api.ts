export async function fetchApi<T>(
  url: string,
  config: RequestInit = {},
  throwOnError: boolean = true
): Promise<T> {
  const r = await fetch(import.meta.env.VITE_API_URL + url, {
    credentials: 'include',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    },
    ...config
  })

  if (throwOnError && !r.ok) {
    throw new Error(r.statusText)
  }

  return r.json()
}

export interface SearchQuery {
  size: number
  query: SearchQueryString
  facets?: SearchFacets
  fields?: string[]
}

interface SearchFacets {
  [key: string]: SearchFacet
}
interface SearchFacet {
  field: string
  size: number
}

interface SearchQueryString {
  query: string
}

/**
 * {
 *   "size": 0,
 *   "from": 0,
 *   "query": {
 *     "query": "*"
 *   },
 *   "facets": {
 *     "tags": {
 *       "field": "tags",
 *       "size": 1000
 *     }
 *   },
 *   "fields": [
 *     "*"
 *   ]
 * }
 */
export interface SearchResult {
  status: StatusNode
  facets?: FacetsNode
}

interface StatusNode {
  total: number
  failed: number
  success: number
}

interface FacetsNode {
  [key: string]: FacetNode
}

interface FacetNode {
  field: string
  total: number
  missing: number
  other: number
  terms: FacetTerm[]
}

interface FacetTerm {
  term: string
  count: number
}

export async function fetchSearch(query: SearchQuery): Promise<SearchResult> {
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
