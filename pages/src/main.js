import { createApp } from 'vue';
import App from './App.vue';
import router from './router';


// Tabler CSS
import '@fontsource/inter';
import '@tabler/core/dist/css/demo.min.css';
import '@tabler/core/dist/css/tabler-flags.min.css';
import '@tabler/core/dist/css/tabler-payments.min.css';
import '@tabler/core/dist/css/tabler-vendors.min.css';
import '@tabler/core/dist/css/tabler.min.css';

// Tabler JS
import '@tabler/core/dist/js/demo-theme.min.js';
import '@tabler/core/dist/js/demo.min.js';
import '@tabler/core/dist/js/tabler.min.js';

// Lib
import 'apexcharts/dist/apexcharts.min.js';
import 'jsvectormap/dist/jsvectormap.min.js';
import 'jsvectormap/dist/maps/world-merc.js';

createApp(App).use(router).mount('#app');