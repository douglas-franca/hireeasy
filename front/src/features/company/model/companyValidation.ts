import { z } from 'zod'

export const createCompanySchema = z.object({
  name: z.string().trim().min(2, 'Enter your company name.'),
})

export type CreateCompanyFormValues = z.infer<typeof createCompanySchema>
