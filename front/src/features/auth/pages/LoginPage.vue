<script setup lang="ts">
import { useRouter } from 'vue-router'

import LoginForm from '@/features/auth/components/LoginForm.vue'
import { useLogin } from '@/features/auth/composables/useLogin'
import { type LoginFormValues } from '@/features/auth/model/authValidation'
import { useAuthStore } from '@/features/auth/stores/auth'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/shared/ui/card'

const router = useRouter()
const authStore = useAuthStore()
const login = useLogin()
const { errorMessage, isSubmitting } = login

async function submit(values: LoginFormValues) {
  const isLoggedIn = await login.submit(values)

  if (!isLoggedIn) return

  const isAuthenticated = await authStore.refresh()

  if (!isAuthenticated) return

  if (!authStore.isEmailVerified) {
    await router.replace({ name: 'verify-email' })
    return
  }

  await router.replace({ name: 'home' })
}
</script>

<template>
  <main class="flex min-h-svh items-center justify-center bg-muted/40 p-4">
    <Card class="w-full max-w-md">
      <CardHeader>
        <CardTitle>Welcome back</CardTitle>
        <CardDescription>Sign in to continue to Hire Easy.</CardDescription>
      </CardHeader>

      <CardContent>
        <LoginForm :error-message="errorMessage" :is-submitting="isSubmitting" @submit="submit" />

        <p class="mt-6 text-center text-sm text-muted-foreground">
          New to Hire Easy?
          <RouterLink class="font-medium text-foreground underline" :to="{ name: 'register' }"
            >Create an account</RouterLink
          >
        </p>
      </CardContent>
    </Card>
  </main>
</template>
