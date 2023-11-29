export class ApiError extends Error {
  public statusCode: number

  constructor(message: string, statusCode: number) {
    super(message)
    this.statusCode = statusCode
    this.name = 'ApiError'
  }
}

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
    throw new ApiError(r.statusText, r.status)
  }

  return r.json()
}

export async function openLocalPath(path: string) {
  return await fetchApi('/api/open', {
    method: 'POST',
    body: JSON.stringify({ path })
  })
}

export async function suggestSearchResult(id: string) {
  return await fetchApi('/api/suggestions', {
    method: 'POST',
    body: JSON.stringify({ id })
  })
}
