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
