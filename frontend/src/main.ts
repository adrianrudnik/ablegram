import './assets/styles/main.scss'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import PrimeVue from 'primevue/config'
import ConfirmationService from 'primevue/confirmationservice'
import BadgeDirective from 'primevue/badgedirective'
import Tooltip from 'primevue/tooltip'
import ToastService from 'primevue/toastservice'

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
  .directive('badge', BadgeDirective)
  .directive('tooltip', Tooltip)
  .mount('#app')
