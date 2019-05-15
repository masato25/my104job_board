require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

import Vue from "vue";
// use element ui
require("./setup/initelement");

import VueRouter from 'vue-router'
import Vuex from 'vuex';
Vue.config.productionTip = false;
Vue.use(VueRouter)
Vue.use(Vuex);

import AppRow from "./components/app_row.vue";
import storep from '~/store/app';
const store = new Vuex.Store(storep);

const routes = [
  {path: '/', components: AppRow}
]
const router = new VueRouter({
  router,
})
/*new Vue({
  router,
}).$mount('#app')
*/

new Vue({
  el: '#app',
  store,
  render: h => h(AppRow)
})
