import { shallowRef } from 'vue'

import { authApi } from '@/features/auth/services/api/auth.api'
import { toApiError } from '@/shared/lib/apiError'

export function useEmailVerification() {
  const errorMessage = shallowRef<string | null>(null)
  const isConfirming = shallowRef(false)
  const isConfirmed = shallowRef(false)

  async function confirm(token: string) {
    if (isConfirming.value || isConfirmed.value) return false

    errorMessage.value = null
    isConfirming.value = true

    try {
      const payload = { token }
      await authApi.confirmEmailVerification(payload)
      isConfirmed.value = true
      return true
    } catch (error) {
      errorMessage.value = toApiError(error).message
      return false
    } finally {
      isConfirming.value = false
    }
  }

  return { errorMessage, isConfirmed, isConfirming, confirm }
}
