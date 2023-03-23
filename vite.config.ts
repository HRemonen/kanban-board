import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    modulePreload: false,
    manifest: true,
    rollupOptions: {
      input: './index.html',
    },
  },
  server: {
    proxy: {
      '/api/v1': 'http://localhost:8080',
    },
    watch: {
      usePolling: true,
    },
    host: true,
    strictPort: true,
    port: 5173,
  },
})
