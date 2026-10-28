<script setup lang="ts">
import { onMounted, shallowRef } from 'vue'

import AppSidebarContent from '@/app/layouts/App/components/AppSidebarContent.vue'
import AppSidebarHeader from '@/app/layouts/App/components/AppSidebarHeader.vue'
import CreateCompanyDrawer from '@/features/company/components/CreateCompanyDrawer.vue'
import { useCompanyOnboarding } from '@/features/company/composables/useCompanyOnboarding'
import { Alert, AlertDescription, AlertTitle } from '@/shared/ui/alert'
import { Separator } from '@/shared/ui/separator'
import { Sidebar, SidebarInset, SidebarProvider, SidebarRail, SidebarTrigger } from '@/shared/ui/sidebar'

const loadError = shallowRef<string | null>(null)
const onboarding = useCompanyOnboarding()
const { errorMessage, isSubmitting, shouldShowDrawer } = onboarding

onMounted(async () => {
  try {
    await onboarding.loadCompanies()
  } catch {
    loadError.value = 'We could not load your companies. Refresh the page to try again.'
  }
})
</script>

<template>
  <SidebarProvider>
    <Sidebar collapsible="icon">
      <AppSidebarHeader @create-company="onboarding.openDrawer" />
      <AppSidebarContent />
      <SidebarRail />
    </Sidebar>

    <SidebarInset>
      <header class="sticky top-0 border-b bg-background px-4 sm:px-6">
        <div class="mx-auto flex w-full items-center gap-4 py-3">
          <SidebarTrigger />
          <Separator orientation="vertical" class="h-4" />
          <span class="text-sm font-medium">Hire Easy</span>
        </div>
      </header>

      <main class="min-w-0 flex-1 px-4 py-6 sm:px-6">
        <Alert v-if="loadError" variant="destructive" class="mb-6">
          <AlertTitle>Companies unavailable</AlertTitle>
          <AlertDescription>{{ loadError }}</AlertDescription>
        </Alert>

        <RouterView />
      </main>
    </SidebarInset>

    <CreateCompanyDrawer
      :is-open="shouldShowDrawer"
      :is-submitting="isSubmitting"
      :error-message="errorMessage"
      @close="onboarding.closeDrawer"
      @submit="onboarding.createCompany"
    />
  </SidebarProvider>
</template>
