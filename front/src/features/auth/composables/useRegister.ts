import { shallowRef } from 'vue'

import { type RegisterFormValues } from '@/features/auth/model/authValidation'
import { authApi } from '@/features/auth/services/api/auth.api'
import { toApiError } from '@/shared/lib/apiError'

export function useRegister() {
  const errorMessage = shallowRef<string | null>(null)
  const isSubmitting = shallowRef(false)

  async function submit(values: RegisterFormValues) {
    if (isSubmitting.value) return false

    errorMessage.value = null
    isSubmitting.value = true

    try {
      const payload = { name: values.name, email: values.email, password: values.password }
      await authApi.register(payload)
      return true
    } catch (error) {
      errorMessage.value = toApiError(error).message
      return false
    } finally {
      isSubmitting.value = false
    }
  }

  return { errorMessage, isSubmitting, submit }
}
