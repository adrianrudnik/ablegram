import type { Ref } from 'vue'
import { computed, ref } from 'vue'

export interface StoreResource {
  id: string
}

export function setupStore<T extends StoreResource>() {
  return () => {
    const entries = ref<T[]>([]) as Ref<T[]>

    function update(resource: T) {
      const idx = entries.value.findIndex((entry) => resource.id === entry.id)

      if (idx === -1) {
        entries.value.push(resource)
      } else {
        entries.value[idx] = resource
      }
    }

    function remove(resource: T) {
      const idx = entries.value.findIndex((entry) => resource.id === entry.id)

      if (idx !== -1) {
        entries.value.splice(idx, 1)
      }
    }

    function get(id: string): T | undefined {
      return entries.value.find((entry) => entry.id === id)
    }

    const count = computed(() => entries.value.length)

    function clear(): void {
      entries.value = []
    }

    return { entries, update, remove, get, count, clear }
  }
}
