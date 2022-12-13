import { createApp } from 'vue';
import { createPinia } from 'pinia';
import './index.css';

import App from './App.vue'
import router from './router'
import VueApexCharts from "vue3-apexcharts"

import OpenLayersMap from 'vue3-openlayers';
import 'vue3-openlayers/dist/vue3-openlayers.css';

const app = createApp(App);

app.use(createPinia())
app.use(VueApexCharts)
app.use(router)
app.use(OpenLayersMap)

app.mount('#app');
