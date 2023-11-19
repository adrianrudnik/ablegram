export function createIdFrom(v: string): string {
  return v.replace(/\W/g, '_')
}
