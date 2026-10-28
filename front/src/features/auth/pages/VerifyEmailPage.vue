<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import { useEmailVerification } from '@/features/auth/composables/useEmailVerification'
import { Alert, AlertDescription, AlertTitle } from '@/shared/ui/alert'
import { Button } from '@/shared/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/shared/ui/card'
import { Spinner } from '@/shared/ui/spinner'

const route = useRoute()
const router = useRouter()
const verification = useEmailVerification()
const { errorMessage, isConfirming } = verification
const token = computed(() => (typeof route.query.token === 'string' ? route.query.token : null))

onMounted(async () => {
  if (!token.value) return

  const isConfirmed = await verification.confirm(token.value)

  if (isConfirmed) await router.replace({ name: 'login', query: { verified: 'true' } })
})
</script>

<template>
  <main class="flex min-h-svh items-center justify-center bg-muted/40 p-4">
    <Card class="w-full max-w-md">
      <CardHeader>
        <CardTitle>Verify your email</CardTitle>
        <CardDescription>We need to verify your email before you create a company.</CardDescription>
      </CardHeader>

      <CardContent class="flex flex-col gap-6">
        <div v-if="token && isConfirming" class="flex items-center gap-3 text-sm text-muted-foreground">
          <Spinner />
          Confirming your email address…
        </div>

        <Alert v-else-if="errorMessage" variant="destructive">
          <AlertTitle>Verification failed</AlertTitle>
          <AlertDescription>{{ errorMessage }}</AlertDescription>
        </Alert>

        <p v-else>Check your inbox and open the verification link to activate your account.</p>

        <Button as-child>
          <RouterLink :to="{ name: 'login' }">Go to sign in</RouterLink>
        </Button>
      </CardContent>
    </Card>
  </main>
</template>
