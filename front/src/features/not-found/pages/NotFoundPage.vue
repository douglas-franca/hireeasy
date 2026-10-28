<script setup lang="ts">
import { computed } from 'vue'

import { useAuthStore } from '@/features/auth/stores/auth'
import { Button } from '@/shared/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/shared/ui/card'

const authStore = useAuthStore()
const destination = computed(() => (authStore.isAuthenticated ? { name: 'home' } : { name: 'login' }))
</script>

<template>
  <main class="flex min-h-svh items-center justify-center bg-muted/40 p-4">
    <Card class="w-full max-w-md text-center">
      <CardHeader>
        <CardTitle>Page not found</CardTitle>
        <CardDescription>The page you requested does not exist.</CardDescription>
      </CardHeader>

      <CardContent>
        <Button as-child>
          <RouterLink :to="destination">Return to safety</RouterLink>
        </Button>
      </CardContent>
    </Card>
  </main>
</template>
