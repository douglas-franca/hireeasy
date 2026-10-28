import { computed, shallowRef } from 'vue'

import { type CreateCompanyFormValues } from '@/features/company/model/companyValidation'
import { companyApi } from '@/features/company/services/api/company.api'
import { useCompanyStore } from '@/features/company/stores/company'
import { toApiError } from '@/shared/lib/apiError'

export function useCompanyOnboarding() {
  const companyStore = useCompanyStore()
  const isDrawerOpen = shallowRef(false)
  const isSubmitting = shallowRef(false)
  const errorMessage = shallowRef<string | null>(null)
  const shouldShowDrawer = computed(() => isDrawerOpen.value)

  async function loadCompanies() {
    await companyStore.load()

    if (companyStore.companies.length === 0) isDrawerOpen.value = true
  }

  function openDrawer() {
    errorMessage.value = null
    isDrawerOpen.value = true
  }

  function closeDrawer() {
    if (isSubmitting.value || companyStore.companies.length === 0) return

    isDrawerOpen.value = false
  }

  async function createCompany(values: CreateCompanyFormValues) {
    if (isSubmitting.value) return false

    errorMessage.value = null
    isSubmitting.value = true

    try {
      const payload = { name: values.name }
      const company = await companyApi.create(payload)
      companyStore.add(company)
      isDrawerOpen.value = false
      return true
    } catch (error) {
      errorMessage.value = toApiError(error).message
      return false
    } finally {
      isSubmitting.value = false
    }
  }

  return {
    errorMessage,
    isSubmitting,
    shouldShowDrawer,
    loadCompanies,
    openDrawer,
    closeDrawer,
    createCompany,
  }
}
