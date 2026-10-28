<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

import { type NavigationRouteItem, navigationGroups } from '@/app/navigation'
import {
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/shared/ui/sidebar'

const route = useRoute()
const groups = computed(() => navigationGroups)

function isRouteActive(routeName: NavigationRouteItem['routeName']) {
  return route.name === routeName || route.meta.navigationRouteName === routeName
}
</script>

<template>
  <SidebarContent>
    <SidebarGroup v-for="group in groups" :key="group.labelKey">
      <SidebarGroupLabel>{{ $t(group.labelKey) }}</SidebarGroupLabel>

      <SidebarGroupContent>
        <SidebarMenu>
          <SidebarMenuItem v-for="item in group.items" :key="item.routeName">
            <SidebarMenuButton as-child :is-active="isRouteActive(item.routeName)">
              <RouterLink :to="{ name: item.routeName }">
                <component :is="item.icon" />
                <span>{{ $t(item.labelKey) }}</span>
              </RouterLink>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarGroupContent>
    </SidebarGroup>
  </SidebarContent>
</template>
