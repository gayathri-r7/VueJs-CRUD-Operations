import Vue from 'vue'
import App from './App.vue'

// Import Bootstrap and BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// starts the b-table import.
import { TablePlugin } from 'bootstrap-vue'
Vue.use(TablePlugin)
// ends the b-table import.

// Import Vue Modal component...
import { ModalPlugin } from 'bootstrap-vue'
Vue.use(ModalPlugin)

// Toast Message
import { ToastPlugin } from 'bootstrap-vue'
Vue.use(ToastPlugin)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
