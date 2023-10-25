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

  const preload = async () => {
    try {
      const response = await fetch('/about.json')
      const about = await response.json()

      version.value = about.version ?? '0.0.0-unknown'
      versionCommitHash.value = about['commit-hash'] ?? 'unknown'
    } catch (error: any) {
      console.debug('Failed to load about.json: ' + error.toString())
    }
  }

  return {
    version,
    versionCommitHash,
    currentLocale,
    preload
  }
})
