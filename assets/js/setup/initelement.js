import Vue from "vue";
import VueI18n from 'vue-i18n'
import zhLocale from 'element-ui/lib/locale/lang/zh-TW';
import Element from 'element-ui'

Vue.use(VueI18n)

const tlang = {
  zh: {
    hello: '你好',
    ...zhLocale,
  }
}
const i18n = new VueI18n({
  locale: 'zh',
  tlang
})

Vue.use(Element, {
  i18n: (key, value) => i18n.t(key, value)
})
import 'element-ui/lib/theme-default/index.css'
