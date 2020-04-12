import Vue from 'vue'
import VueRouter from 'vue-router';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

import App from './App.vue'
import Home from './components/Home.vue'
import Mash from './components/Mash.vue'
import MashCreate from './components/MashCreate.vue'
import MashResult from './components/MashResult.vue'

Vue.config.productionTip = false;
Vue.use(VueRouter);
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);

// const NotFound = { template: '<p>Page not found</p>' }

const router = new VueRouter({
  routes: [
    { path: '/', component: Home},
    { path: '/mash/create', component: MashCreate },
    { path: '/mash/:id', component: Mash },
    { path: '/mash/:id/result', component: MashResult },
  ]
});


new Vue({
  render: h => h(App),
  router
}).$mount('#app')