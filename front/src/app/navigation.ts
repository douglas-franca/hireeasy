import { HomeIcon, type LucideIcon } from '@lucide/vue'

export interface NavigationRouteItem {
  icon: LucideIcon
  labelKey: string
  routeName: string
}

export interface NavigationGroup {
  labelKey: string
  items: NavigationRouteItem[]
}

export const navigationGroups: NavigationGroup[] = [
  {
    labelKey: 'app.menu.groups.main.title',
    items: [{ icon: HomeIcon, labelKey: 'app.menu.groups.main.home', routeName: 'home' }],
  },
]
