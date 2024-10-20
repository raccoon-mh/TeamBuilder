import { createRouter, createWebHashHistory } from 'vue-router';
import AboutPage from '../components/AboutPage.vue';
import HomePage from '../components/HomePage.vue';
import LicensePage from '../components/LicensePage.vue';

const routes = [
  { path: '/', component: HomePage },
  { path: '/about', component: AboutPage },
  { path: '/license', component: LicensePage }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
