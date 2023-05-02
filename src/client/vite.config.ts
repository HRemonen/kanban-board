import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  base: '/',
  server: {
    proxy: {
      '/api/v1': {
        target: 'http://web:8080', // Web (API service name) instead of localhost because we are inside docker compose network: https://stackoverflow.com/questions/74281247/api-call-not-reaching-vite-proxy-target-localhost5000
        changeOrigin: true,
        secure: false,
      },
    },
    watch: {
      usePolling: true,
    },
    host: true,
    strictPort: true,
    port: 5173,
  },
})
