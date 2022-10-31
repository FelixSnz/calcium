import { type RouteRecordRaw, createRouter, createWebHistory } from "vue-router";

const routes: RouteRecordRaw[] = [
    

    {
        path: '/',
        name: 'home',
        component: () => import('../components/Navbar.vue')

    },
    {
        path: '/table',
        name: 'table',
        component: () => import('../components/EmailMatchesTable.vue')

    }

]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router;