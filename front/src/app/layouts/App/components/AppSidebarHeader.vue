<script setup lang="ts">
import { Building2Icon, CheckIcon, ChevronsUpDownIcon, PlusIcon } from '@lucide/vue'
import { computed } from 'vue'

import { type Company } from '@/features/company/services/api/company.types'
import { useCompanyStore } from '@/features/company/stores/company'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/shared/ui/dropdown-menu'
import { SidebarHeader, SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '@/shared/ui/sidebar'

const emit = defineEmits<{ createCompany: [] }>()
const companyStore = useCompanyStore()
const company = computed(() => companyStore.selectedCompany)
const companies = computed(() => companyStore.companies)

function selectCompany(selectedCompany: Company) {
  companyStore.select(selectedCompany)
}
</script>

<template>
  <SidebarHeader>
    <SidebarMenu>
      <SidebarMenuItem v-if="company">
        <DropdownMenu>
          <DropdownMenuTrigger as-child>
            <SidebarMenuButton size="lg">
              <Building2Icon />
              <span>{{ company.name }}</span>
              <ChevronsUpDownIcon class="ml-auto" />
            </SidebarMenuButton>
          </DropdownMenuTrigger>

          <DropdownMenuContent side="right" align="start" class="w-64">
            <DropdownMenuLabel>Companies</DropdownMenuLabel>
            <DropdownMenuGroup>
              <DropdownMenuItem
                v-for="listedCompany in companies"
                :key="listedCompany.id"
                @select="selectCompany(listedCompany)"
              >
                <Building2Icon />
                <span>{{ listedCompany.name }}</span>
                <CheckIcon v-if="listedCompany.id === company.id" class="ml-auto" />
              </DropdownMenuItem>
            </DropdownMenuGroup>

            <DropdownMenuSeparator />

            <DropdownMenuGroup>
              <DropdownMenuItem @select="emit('createCompany')">
                <PlusIcon />
                Create company
              </DropdownMenuItem>
            </DropdownMenuGroup>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarMenuItem>
    </SidebarMenu>
  </SidebarHeader>
</template>
