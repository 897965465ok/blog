// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui';
import { api } from './api'
import store from './store/store'
import './components'
import util from './util/util'
Vue.use(ElementUI);
Vue.use(util)
Vue.config.productionTip = false
Vue.prototype.$api = api
import "tailwindcss/tailwind.css"
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
