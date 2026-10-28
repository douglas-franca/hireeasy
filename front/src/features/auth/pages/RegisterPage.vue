<script setup lang="ts">
import { useRouter } from 'vue-router'

import RegisterForm from '@/features/auth/components/RegisterForm.vue'
import { useRegister } from '@/features/auth/composables/useRegister'
import { type RegisterFormValues } from '@/features/auth/model/authValidation'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/shared/ui/card'

const router = useRouter()
const registration = useRegister()
const { errorMessage, isSubmitting } = registration

async function submit(values: RegisterFormValues) {
  const isRegistered = await registration.submit(values)

  if (isRegistered) await router.replace({ name: 'verify-email' })
}
</script>

<template>
  <main class="flex min-h-svh items-center justify-center bg-muted/40 p-4">
    <Card class="w-full max-w-md">
      <CardHeader>
        <CardTitle>Create your account</CardTitle>
        <CardDescription>Start managing your hiring workflow with Hire Easy.</CardDescription>
      </CardHeader>

      <CardContent>
        <RegisterForm :error-message="errorMessage" :is-submitting="isSubmitting" @submit="submit" />

        <p class="mt-6 text-center text-sm text-muted-foreground">
          Already have an account?
          <RouterLink class="font-medium text-foreground underline" :to="{ name: 'login' }"
            >Sign in</RouterLink
          >
        </p>
      </CardContent>
    </Card>
  </main>
</template>
