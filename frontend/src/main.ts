import { createApp } from "vue";
import App from "./App.vue";
import router from './router';
import './style/main.css'
import { VueFinalModal } from 'vue-final-modal';
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'

const app = createApp(App);
app.use(router)
app.use(VueFinalModal)
app.use(Toast)
app.mount('#app')
