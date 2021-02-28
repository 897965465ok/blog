import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)
export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('@/layout/Home/Home'),
      children: [
        {
          path: '/',
          name: 'main',
          component: () => import('@/components/Main/Main')
        },
        {
          path: '/note',
          name: 'note',
          component: () => import('@/components/Note/Note')
        },
        {
          path: '/gossip',
          name: 'gossip',
          component: () => import('@/components/Gossip/Gossip')

        },
        {
          path: '/login',
          name: 'login',
          component: () => import('@/components/Login/Login')
        },

        {
          path: '/about',
          name: 'about',
          component: () => import('@/components/About/About')
        },
        {
          path: '/markdown',
          name: 'Markdown',
          component: () => import('@/components/Markdown/Markdown.vue')
        },
      ]
    },
  ]
})
