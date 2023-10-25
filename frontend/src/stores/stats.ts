import i18n from '@/plugins/i18n'
import { defineStore } from 'pinia'
import { computed, ref, watch } from 'vue'

export const useStatStore = defineStore('stats', () => {
  const files = ref(0)
  const midiTracks = ref(0)
  const audioTracks = ref(0)

  return {
    files,
    midiTracks,
    audioTracks
  }
})
