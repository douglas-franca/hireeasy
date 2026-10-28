export interface RegisterUserPayload {
  name: string
  email: string
  password: string
}

export interface LoginPayload {
  email: string
  password: string
}

export interface ConfirmEmailVerificationPayload {
  token: string
}

export interface CurrentUser {
  id: string
  name: string
  email: string
  email_verified: boolean
}

export interface RegisterUserResponse {
  message: string
}
