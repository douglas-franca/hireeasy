import {
  type ConfirmEmailVerificationPayload,
  type CurrentUser,
  type LoginPayload,
  type RegisterUserPayload,
  type RegisterUserResponse,
} from '@/features/auth/services/api/auth.types'
import { http } from '@/shared/lib/http'

export const authApi = {
  async register(payload: RegisterUserPayload) {
    const response = await http.post<RegisterUserResponse>('/users', payload)
    return response.data
  },

  async login(payload: LoginPayload) {
    await http.post<void>('/sessions', payload)
  },

  async getCurrentUser() {
    const response = await http.get<CurrentUser>('/users/me')
    return response.data
  },

  async confirmEmailVerification(payload: ConfirmEmailVerificationPayload) {
    await http.post<void>('/users/verification/confirm', payload)
  },
}
