import { z } from 'zod'

const MINIMUM_PASSWORD_LENGTH = 8

export const loginSchema = z.object({
  email: z.email('Enter a valid email address.'),
  password: z.string().min(1, 'Enter your password.'),
})

export const registerSchema = z.object({
  name: z.string().trim().min(2, 'Enter your name.'),
  email: z.email('Enter a valid email address.'),
  password: z.string().min(MINIMUM_PASSWORD_LENGTH, 'Use at least 8 characters.'),
})

export type LoginFormValues = z.infer<typeof loginSchema>
export type RegisterFormValues = z.infer<typeof registerSchema>
