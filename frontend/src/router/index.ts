import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import Landing from '@/views/Landing.vue'
import Dashboard from '@/views/Dashboard.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/me/profile',
    name: 'Dashboard',
    component: Dashboard,
  },
  {
    path: '/',
    name: 'Langing',
    component: Landing,
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
