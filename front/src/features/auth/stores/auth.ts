import { defineStore } from 'pinia'
import { computed, shallowRef } from 'vue'

import { authApi } from '@/features/auth/services/api/auth.api'
import { type CurrentUser } from '@/features/auth/services/api/auth.types'
import { toApiError } from '@/shared/lib/apiError'

export const useAuthStore = defineStore('auth', () => {
  const user = shallowRef<CurrentUser | null>(null)
  const isInitialized = shallowRef(false)
  const isInitializing = shallowRef(false)
  const initializationError = shallowRef<string | null>(null)

  const isAuthenticated = computed(() => user.value !== null)
  const isEmailVerified = computed(() => user.value?.email_verified === true)

  async function initialize() {
    if (isInitialized.value || isInitializing.value) return

    isInitializing.value = true
    initializationError.value = null

    try {
      await refresh()
    } finally {
      isInitialized.value = true
      isInitializing.value = false
    }
  }

  async function refresh() {
    initializationError.value = null

    try {
      user.value = await authApi.getCurrentUser()
      return true
    } catch (error) {
      const apiError = toApiError(error)

      if (apiError.status === 404) {
        user.value = null
        return false
      }

      initializationError.value = apiError.message
      return false
    }
  }

  function setUser(currentUser: CurrentUser) {
    user.value = currentUser
    isInitialized.value = true
  }

  function clearUser() {
    user.value = null
    isInitialized.value = true
  }

  return {
    user,
    isInitialized,
    isInitializing,
    initializationError,
    isAuthenticated,
    isEmailVerified,
    initialize,
    refresh,
    setUser,
    clearUser,
  }
})
