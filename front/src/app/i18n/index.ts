import { createI18n } from 'vue-i18n'

import { messages } from '@/app/i18n/messages'
import { localStorageService } from '@/shared/lib/localStorage'

export const supportedLocales = ['en', 'pt', 'es'] as const
export type SupportedLocale = (typeof supportedLocales)[number]

const DEFAULT_LOCALE: SupportedLocale = 'pt'

function isSupportedLocale(value: string): value is SupportedLocale {
  return supportedLocales.includes(value as SupportedLocale)
}

function getInitialLocale(): SupportedLocale {
  const savedLocale = localStorageService.get('language')

  if (savedLocale && isSupportedLocale(savedLocale)) return savedLocale

  return DEFAULT_LOCALE
}

const i18n = createI18n({
  legacy: false,
  locale: getInitialLocale(),
  fallbackLocale: DEFAULT_LOCALE,
  messages,
})

export function setLocale(locale: SupportedLocale) {
  i18n.global.locale.value = locale
  localStorageService.set('language', locale)
}

export default i18n
