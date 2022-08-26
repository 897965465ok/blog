import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    { path: '/', component: '@/pages/index' },
    { path: '/point', component: '@/pages/point/point' },
    { path: '/event', component: '@/pages/event/Event' },
    { path: '/popo', component: '@/pages/popo/popo' },
    { path: '/children', component: '@/pages/children/children' },
    {
      path: '/home', component: '@/pages/home/home',
      children: {
        path: '/home', component: '@/pages/home/home',
      }
    },
  ],
  fastRefresh: {},
  "sass": {}
});
