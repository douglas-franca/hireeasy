import { type Company, type CreateCompanyPayload } from '@/features/company/services/api/company.types'
import { http } from '@/shared/lib/http'

export const companyApi = {
  async list() {
    const response = await http.get<Company[]>('/companies')
    return response.data
  },

  async create(payload: CreateCompanyPayload) {
    const response = await http.post<Company>('/companies', payload)
    return response.data
  },
}
