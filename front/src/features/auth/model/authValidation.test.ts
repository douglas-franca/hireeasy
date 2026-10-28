import { describe, expect, it } from 'vitest'

import { loginSchema, registerSchema } from '@/features/auth/model/authValidation'

describe('auth validation', () => {
  it('rejects invalid login credentials', () => {
    const result = loginSchema.safeParse({ email: 'invalid', password: '' })

    expect(result.success).toBe(false)
  })

  it('accepts a valid registration payload', () => {
    const result = registerSchema.safeParse({
      name: 'Ana Silva',
      email: 'ana@example.com',
      password: 'password123',
    })

    expect(result.success).toBe(true)
  })
})
