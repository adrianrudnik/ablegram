import { computed } from 'vue'
import { createI18n } from 'vue-i18n'
import messages from '@intlify/unplugin-vue-i18n/messages'
import { useUiStore } from '@/stores/ui'
export const i18n = createI18n<false>({
  fallbackLocale: 'en',
  locale: navigator.language ?? 'en',
  messages,
  legacy: false,
  fallbackWarn: false,
  missingWarn: false,
  silentFallbackWarn: true,
  silentTranslationWarn: true
})

export const t = i18n.global.t

export default i18n
