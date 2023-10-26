import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useStatStore = defineStore('stats', () => {
  const validFiles = ref(0)
  const invalidFiles = ref(0)
  const liveSets = ref(0)
  const indexDocuments = ref(0)
  const midiTracks = ref(0)
  const audioTracks = ref(0)

  const update = (k: string, v: number) => {
    switch (k) {
      case 'files_valid':
        validFiles.value = v
        break
      case 'files_invalid':
        invalidFiles.value = v
        break
      case 'live_sets':
        liveSets.value = v
        break
      case 'index_docs':
        indexDocuments.value = v
        break
      case 'midi_tracks':
        midiTracks.value = v
        break
      case 'audio_tracks':
        audioTracks.value = v
        break
    }
  }

  return {
    update,
    validFiles,
    invalidFiles,
    liveSets,
    indexDocuments,
    midiTracks,
    audioTracks
  }
})
