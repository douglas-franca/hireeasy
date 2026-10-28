import { AxiosError, type AxiosResponse } from 'axios'
import { createPinia, setActivePinia } from 'pinia'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { authApi } from '@/features/auth/services/api/auth.api'
import { useAuthStore } from '@/features/auth/stores/auth'
import { makeUser } from '@/test-utils/factories'

vi.mock('@/features/auth/services/api/auth.api', () => ({
  authApi: {
    register: vi.fn(),
    login: vi.fn(),
    getCurrentUser: vi.fn(),
    confirmEmailVerification: vi.fn(),
  },
}))

describe('auth store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  it('stores the authenticated user during initialization', async () => {
    const user = makeUser()
    vi.mocked(authApi.getCurrentUser).mockResolvedValue(user)

    const store = useAuthStore()
    await store.initialize()

    expect(store.user).toEqual(user)
    expect(store.isAuthenticated).toBe(true)
  })

  it('treats a 404 current-user response as anonymous', async () => {
    const response = {
      status: 404,
      data: { error: { code: 'UNAUTHENTICATED', message: 'Missing session.' } },
    } as AxiosResponse
    vi.mocked(authApi.getCurrentUser).mockRejectedValue(
      new AxiosError('Not found', undefined, undefined, undefined, response),
    )

    const store = useAuthStore()
    await store.initialize()

    expect(store.user).toBeNull()
    expect(store.isAuthenticated).toBe(false)
  })
})
