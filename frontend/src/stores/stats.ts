import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useStatStore = defineStore('stats', () => {
  const isProcessing = ref(false)
  const isSearching = ref(false)

  // Processing related
  const processRoutineCount = ref(-1)

  // Search related
  const searchResultCount = ref(0)

  // Meta status bar
  const validFileCount = ref(0)
  const invalidFileCount = ref(0)
  const liveSetCount = ref(0)
  const indexDocumentCount = ref(0)
  const midiTrackCount = ref(0)
  const audioTrackCount = ref(0)

  const inProgress = computed(() => isProcessing.value || isSearching.value)

  const updateMetrics = (values: { [key: string]: number }) => {
    for (const [k, v] of Object.entries(values)) {
      switch (k) {
        case 'files_valid':
          validFileCount.value = v
          break
        case 'files_invalid':
          invalidFileCount.value = v
          break
        case 'live_sets':
          liveSetCount.value = v
          break
        case 'index_docs':
          indexDocumentCount.value = v
          break
        case 'midi_tracks':
          midiTrackCount.value = v
          break
        case 'audio_tracks':
          audioTrackCount.value = v
          break
      }
    }
  }

  return {
    updateMetrics,
    isProcessing,
    isSearching,
    inProgress,
    processRoutineCount,
    searchResultCount,
    validFileCount,
    invalidFileCount,
    liveSetCount,
    indexDocumentCount,
    midiTrackCount,
    audioTrackCount
  }
})
