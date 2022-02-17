import { createRouter, createWebHashHistory, Router } from "vue-router"
export const router: Router = createRouter({
    history: createWebHashHistory(),
    routes: [{
        path: '/',
        component: () => import("../pages/main/main.vue"),
        children: [
            {
                path: '/',
                component: () => import("../pages/home/home.vue")
            },
            {
                path: '/banner',
                component: () => import("../pages/banner/Banner.vue")
            },
            {
                path: '/article',
                component: () => import("../pages/article/article.vue")
            },
            {
                path: '/user',
                component: () => import("../pages/user/user.vue")
            }
        ]
    }]
})

