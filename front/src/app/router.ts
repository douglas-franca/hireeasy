import { createRouter, createWebHistory } from 'vue-router'

import AppLayout from '@/app/layouts/App/AppLayout.vue'
import LoginPage from '@/features/auth/pages/LoginPage.vue'
import RegisterPage from '@/features/auth/pages/RegisterPage.vue'
import VerifyEmailPage from '@/features/auth/pages/VerifyEmailPage.vue'
import { useAuthStore } from '@/features/auth/stores/auth'
import HomePage from '@/features/home/pages/HomePage.vue'
import NotFoundPage from '@/features/not-found/pages/NotFoundPage.vue'

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
    guestOnly?: boolean
  }
}

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: AppLayout,
      meta: { requiresAuth: true },
      children: [{ path: '', name: 'home', component: HomePage }],
    },
    { path: '/login', name: 'login', component: LoginPage, meta: { guestOnly: true } },
    { path: '/register', name: 'register', component: RegisterPage, meta: { guestOnly: true } },
    { path: '/verify-email', name: 'verify-email', component: VerifyEmailPage },
    { path: '/:pathMatch(.*)*', name: 'not-found', component: NotFoundPage },
  ],
})

router.beforeEach(async (to) => {
  const authStore = useAuthStore()
  await authStore.initialize()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) return { name: 'login' }

  if (to.meta.requiresAuth && !authStore.isEmailVerified) return { name: 'verify-email' }

  if (to.meta.guestOnly && authStore.isAuthenticated) {
    if (!authStore.isEmailVerified) return { name: 'verify-email' }

    return { name: 'home' }
  }

  return true
})

export default router
