import { defineStore } from 'pinia'
import { computed, shallowRef } from 'vue'

import { companyApi } from '@/features/company/services/api/company.api'
import { type Company } from '@/features/company/services/api/company.types'

export const useCompanyStore = defineStore('company', () => {
  const companies = shallowRef<Company[]>([])
  const selectedCompanyId = shallowRef<string | null>(null)
  const isLoaded = shallowRef(false)

  const selectedCompany = computed(
    () => companies.value.find((company) => company.id === selectedCompanyId.value) ?? null,
  )

  async function load() {
    const listedCompanies = await companyApi.list()
    companies.value = listedCompanies
    selectedCompanyId.value = listedCompanies[0]?.id ?? null
    isLoaded.value = true
  }

  function select(company: Company) {
    selectedCompanyId.value = company.id
  }

  function add(company: Company) {
    companies.value = [...companies.value, company]
    selectedCompanyId.value = company.id
  }

  function reset() {
    companies.value = []
    selectedCompanyId.value = null
    isLoaded.value = false
  }

  return { companies, selectedCompany, isLoaded, load, select, add, reset }
})
