import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
export default defineConfig({
  server: {
    proxy: {
      '/admin': {
        target: 'http://localhost:3800',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/admin/, '/admin')
      },
      '/v1': {
        target: 'http://localhost:3800',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/v1/, '/v1')
      },
    }
  },
  plugins: [
    vue(),
    vueJsx({})
  ]
})
