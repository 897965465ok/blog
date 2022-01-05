import { createRouter, createWebHashHistory, Router } from "vue-router"
export const router: Router = createRouter({
    history: createWebHashHistory(),
    routes: [{
        path: '/',
        component: () => import("../pages/main/main.vue"),
        children: [{
            path: '/banner',
            component: () => import("../pages/banner/banner.vue")
        }]
    }]
})

