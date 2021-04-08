import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)
const router = new Router({
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
          meta: { requiresAuth: true },
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
//  路由拦击代码
// router.beforeEach((to, from, next) => {
//   let result = to.matched.some(record => record.meta.requiresAuth)
//   // 如果是login就是放行
//   if (result) {
//     next()
//   } else {
//     // 如果不是login
//     let token = localStorage.getItem('token')
//     //如果token 存在
//     if (token) {
//       next()
//     } else {
//       // 否则让用户登录
//       next({ path: "/login" })
//     }
//   }
// })


export default router