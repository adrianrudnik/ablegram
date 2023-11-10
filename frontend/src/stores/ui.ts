import i18n from '@/plugins/i18n'
import { defineStore } from 'pinia'
import { computed, ref, watch } from 'vue'

export const useUiStore = defineStore('ui', () => {
  const version = ref('0.0.0-unknown')
  const versionCommitHash = ref('unknown')

  const currentLocale = computed(() => navigator.language ?? 'en')
  watch(currentLocale, async (n) => {
    i18n.global.locale.value = n
  })

  return {
    version,
    versionCommitHash,
    currentLocale
  }
})
