import { createApp } from "vue";
import App from "./App.vue";
import router from './router';
import './style/main.css'
import { VueFinalModal } from 'vue-final-modal';

const app = createApp(App);
app.use(router)
app.use(VueFinalModal)
app.mount('#app')
