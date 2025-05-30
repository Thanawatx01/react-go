import path from 'path'
import { fileURLToPath } from 'url'

// แปลงจาก import.meta.url ไปเป็น __dirname
const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
})
