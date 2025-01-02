import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
// rewrite: (path) => path.replace(/^\/api/, '')
export default defineConfig({
    server: {
        proxy: {
            '/api': {
                target: 'http://localhost:6969',
                changeOrigin: true,
            }
        }
    },
    plugins: [react()],
})
