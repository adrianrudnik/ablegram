export interface SearchQuery {
  size: number
  query: SearchQueryString
  facets?: SearchFacets
  fields?: string[]
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
