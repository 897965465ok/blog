// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui';
import { api } from './api'
import store from './store/store'
import waterfall from 'vue-waterfall2'
import './common'
import util from './util/util'
// import "tailwindcss/tailwind.css"
Vue.use(ElementUI);
Vue.use(waterfall)
Vue.use(util)
Vue.config.productionTip = false
Vue.prototype.$api = api
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
