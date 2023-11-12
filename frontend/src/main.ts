import './assets/styles/main.scss'

import { createApp, nextTick } from 'vue'
import { createPinia } from 'pinia'

import PrimeVue from 'primevue/config'
import ConfirmationService from 'primevue/confirmationservice'
import BadgeDirective from 'primevue/badgedirective'
import Tooltip from 'primevue/tooltip'
import ToastService from 'primevue/toastservice'
import DialogService from 'primevue/dialogservice'

import App from './App.vue'
import router from './router'
import i18n from './plugins/i18n'

const app = createApp(App)

app
  .use(createPinia())
  .use(i18n)
  .use(router)
  .use(PrimeVue)
  .use(ConfirmationService)
  .use(ToastService)
  .use(DialogService)
  .directive('badge', BadgeDirective)
  .directive('tooltip', Tooltip)
  .directive('focus', { mounted: (el) => setTimeout(() => el.focus(), 50) })
  .mount('#app')
