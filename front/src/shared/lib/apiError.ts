import { AxiosError } from 'axios'

interface ApiErrorResponse {
  error: {
    code: string
    message: string
  }
}

export interface ApiError {
  code: string
  message: string
  status?: number
}

export function toApiError(error: unknown): ApiError {
  if (error instanceof AxiosError) {
    const { response } = error
    const data = response?.data as ApiErrorResponse | undefined

    if (data?.error) {
      return { ...data.error, status: response?.status }
    }

    return {
      code: 'NETWORK_ERROR',
      message: 'We could not complete your request. Please try again.',
      status: response?.status,
    }
  }

  return {
    code: 'UNKNOWN_ERROR',
    message: 'We could not complete your request. Please try again.',
  }
}
