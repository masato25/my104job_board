require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

import Vue from "vue";
import VueI18n from 'vue-i18n';
import zhLocale from 'element-ui/lib/locale/lang/zh-TW';
import enLocale from 'element-ui/lib/locale/lang/en';
import Element from 'element-ui';
import 'element-ui/lib/theme-default/index.css';

import VueRouter from 'vue-router'
import Vuex from 'vuex';
Vue.config.productionTip = false;
Vue.use(VueRouter)
Vue.use(Vuex);

// setup i18n
Vue.use(VueI18n)
const messages = {
  zh: zhLocale,
  en: enLocale,
}
const i18n = new VueI18n({
  locale: 'zh',
  messages
})
// use element ui
Vue.use(Element, {
  i18n: (key, value) => i18n.t(key, value)
})

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
  i18n,
  render: h => h(AppRow)
})
