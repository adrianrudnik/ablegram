import type { Ref } from 'vue'
import { computed, ref } from 'vue'

export interface StoreResource {
  id: number | string
}

export function setupStore<T extends StoreResource>() {
  return () => {
    const entries = ref<T[]>([]) as Ref<T[]>

    const all = computed(() => entries.value)

    function overwrite(resource: T[]) {
      entries.value = resource
    }

    function update(resource: T) {
      const idx = entries.value.findIndex((entry) => resource.id === entry.id)

      if (idx === -1) {
        entries.value.push(resource)
      } else {
        entries.value[idx] = resource
      }
    }

    function updateBatch(resources: T[]) {
      resources.forEach(update)
    }

    function remove(resource: T) {
      const idx = entries.value.findIndex((entry) => resource.id === entry.id)

      if (idx !== -1) {
        entries.value.splice(idx, 1)
      }
    }

    function removeById(id: string) {
      const idx = entries.value.findIndex((entry) => id === entry.id)

      if (idx !== -1) {
        entries.value.splice(idx, 1)
      }
    }

    function get(id: string): T | undefined {
      return entries.value.find((entry) => entry.id === id)
    }

    function getRandomElement(): T | undefined {
      if (entries.value.length === 0) return undefined

      const rand = Math.floor(Math.random() * entries.value.length)
      return entries.value[rand]
    }

    const count = computed(() => entries.value.length)

    const clear = () => entries.value.splice(0, entries.value.length)

    return {
      all,
      entries,
      overwrite,
      update,
      updateBatch,
      remove,
      removeById,
      get,
      getRandomElement,
      count,
      clear
    }
  }
}
