import path from 'node:path'
import { fileURLToPath } from 'node:url'
import vue from '@vitejs/plugin-vue'
import { defineConfig } from 'vitest/config'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: { '@': path.resolve(fileURLToPath(new URL('.', import.meta.url)), './src') },
  },
  test: {
    environment: 'jsdom',
    globals: true,
  },
})
