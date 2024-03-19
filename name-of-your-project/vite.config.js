import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      "/assets": {
        target: "http://localhost:5500",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^/, "/name-of-your-project/src"),
      },
    },

  }
})
