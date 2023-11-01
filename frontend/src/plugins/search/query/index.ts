export interface SearchQuery {
  size: number
  query: SearchQueryString
  facets?: SearchFacets
  fields?: string[]
  sort?: string[]
  search_after?: string[]
  search_before?: string[]
}

export interface SearchFacets {
  [key: string]: SearchFacet
}

export interface SearchFacet {
  field: string
  size: number
}

export interface SearchQueryString {
  query: string
}
