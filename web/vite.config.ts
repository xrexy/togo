import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

//       "@components/*": ["./src/components/*"],
// "@views/*": ["./src/views/*"],
// "@composables/*": ["./src/composables/*"],
// "$composables": ["./src/composables"],

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue({reactivityTransform: true})],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@components': fileURLToPath(new URL('./src/components', import.meta.url)),
      '@views': fileURLToPath(new URL('./src/views', import.meta.url)),
      '@composables': fileURLToPath(new URL('./src/composables', import.meta.url)),
    }
  }
})
