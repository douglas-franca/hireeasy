import { createPinia, setActivePinia } from 'pinia'
import { beforeEach, describe, expect, it, vi } from 'vitest'

import { companyApi } from '@/features/company/services/api/company.api'
import { useCompanyStore } from '@/features/company/stores/company'
import { makeCompany } from '@/test-utils/factories'

vi.mock('@/features/company/services/api/company.api', () => ({
  companyApi: { list: vi.fn(), create: vi.fn() },
}))

describe('company store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  it('selects the first company after loading', async () => {
    const firstCompany = makeCompany({ id: 'company-1' })
    const secondCompany = makeCompany({ id: 'company-2' })
    vi.mocked(companyApi.list).mockResolvedValue([firstCompany, secondCompany])

    const store = useCompanyStore()
    await store.load()

    expect(store.selectedCompany).toEqual(firstCompany)
  })

  it('selects a newly created company when added', () => {
    const store = useCompanyStore()
    const company = makeCompany()

    store.add(company)

    expect(store.selectedCompany).toEqual(company)
  })
})
