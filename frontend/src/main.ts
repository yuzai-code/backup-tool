import { createApp } from "vue";
import App from "./App.vue";
import router from './router';
import './style/main.css'
import { createVfm } from 'vue-final-modal';
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'

const app = createApp(App);
const vfm = createVfm()

app.use(router)
app.use(vfm)
app.use(Toast)
app.mount('#app')
