import { createApp } from "vue";
import App from "./App.vue";

import './index.css'
import { createPinia } from "pinia";
import { PiniaColada } from "@pinia/colada";

const app = createApp(App)

const pinia = createPinia()
app.use(pinia)

app.use(PiniaColada)

app.mount("#app");
