import { createApp } from "vue";
import App from "./App.vue";
import PrimeVue from "primevue/config";
import Aura from "@primevue/themes/aura";

createApp(App).mount("#app");

const app = createApp(App);
app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
});
