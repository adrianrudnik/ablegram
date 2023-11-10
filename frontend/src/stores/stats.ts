import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useStatStore = defineStore('stats', () => {
  const isProcessing = ref(false)
  const isSearching = ref(false)

  // Processing related
  const processRoutineCount = ref(-1)

  // Search related
  const searchResultCount = ref(0)

  type CounterEntry = {
    [key: string]: number
  }

  // Meta status bar
  const counters = ref<CounterEntry>({})

  const inProgress = computed(() => isProcessing.value || isSearching.value)

  const updateMetrics = (values: { [key: string]: number }) => {
    for (const [k, v] of Object.entries(values)) {
      counters.value[k] = v
    }
  }

  return {
    updateMetrics,
    isProcessing,
    isSearching,
    inProgress,
    processRoutineCount,
    searchResultCount,
    counters
  }
})
