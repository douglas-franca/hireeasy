import { shallowRef } from 'vue'

import { type LoginFormValues } from '@/features/auth/model/authValidation'
import { authApi } from '@/features/auth/services/api/auth.api'
import { toApiError } from '@/shared/lib/apiError'

export function useLogin() {
  const errorMessage = shallowRef<string | null>(null)
  const isSubmitting = shallowRef(false)

  async function submit(values: LoginFormValues) {
    if (isSubmitting.value) return false

    errorMessage.value = null
    isSubmitting.value = true

    try {
      const payload = { email: values.email, password: values.password }
      await authApi.login(payload)
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
