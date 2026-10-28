import { type CurrentUser } from '@/features/auth/services/api/auth.types'
import { type Company } from '@/features/company/services/api/company.types'

export function makeUser(overrides: Partial<CurrentUser> = {}): CurrentUser {
  return {
    id: 'user-1',
    name: 'Ana Silva',
    email: 'ana@example.com',
    email_verified: true,
    ...overrides,
  }
}

export function makeCompany(overrides: Partial<Company> = {}): Company {
  return {
    id: 'company-1',
    name: 'Acme Ltda',
    role: 'owner',
    ...overrides,
  }
}
