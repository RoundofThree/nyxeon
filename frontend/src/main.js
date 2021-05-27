import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import { library } from "@fortawesome/fontawesome-svg-core"
import { faPhone } from "@fortawesome/free-solid-svg-icons"
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome"
import './assets/index.css'

library.add(faPhone)
//import vuetify from "@/plugins/vuetify"
const app = createApp(App)

app.component("fas", FontAwesomeIcon)
app.use(router, axios)
//app.use(vuetify)

app.mount('#app')
