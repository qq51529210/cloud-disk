import { createApp } from 'vue'
import App from './App.vue'
import i18n from './i18n'

import { Ripple } from '@varlet/ui'
import '@varlet/touch-emulator'

createApp(App)
    .use(i18n)
    .use(Ripple)
    .mount('#app')
