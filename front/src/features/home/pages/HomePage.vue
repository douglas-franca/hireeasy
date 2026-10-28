<script setup lang="ts">
import { computed } from 'vue'

import { useCompanyStore } from '@/features/company/stores/company'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/shared/ui/card'

const companyStore = useCompanyStore()
const companyName = computed(() => companyStore.selectedCompany?.name ?? 'your company')
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Welcome to Hire Easy</CardTitle>
      <CardDescription>Your workspace for {{ companyName }} is ready.</CardDescription>
    </CardHeader>

    <CardContent>
      <p class="text-sm text-muted-foreground">Your hiring dashboard will appear here soon.</p>
    </CardContent>
  </Card>
</template>
