import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import LoginPage from '@/views/LoginPage.vue'
import RegisterPage from '@/views/RegisterPage.vue'
import Landing from '@/views/Landing.vue'
import Dashboard from '@/views/Dashboard.vue'
import VerifyEmailPage from '@/views/VerifyEmailPage.vue'

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
