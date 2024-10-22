import { createRouter, createWebHashHistory } from 'vue-router';
import AboutPage from '../pages/AboutPage.vue';
import HomePage from '../pages/HomePage.vue';
import LicensePage from '../pages/LicensePage.vue';
import LoginPage from '../pages/LoginPage.vue';
import PopupPage from '../pages/PopupPage.vue';


const routes = [
  { path: '/', component: HomePage},
  { path: '/about', component: AboutPage},
  { path: '/license', component: LicensePage},
  { path: '/login', component: LoginPage, meta: { layout: 'BlankLayout' }},
  { path: '/popup', component: PopupPage}
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
