import { defineConfig } from "vite";
import vue from '@vitejs/plugin-vue';
import AutoImport from 'unplugin-auto-import/vite'
import path from 'path'; // 确保引入 path 模块

export default defineConfig({
  plugins: [vue(),
    AutoImport({
      // 自动导入 vue 和 vue-router ，如：ref, computed, watch, useRouter, useStore等
      imports: ['vue', 'vue-router'],
      eslintrc:{
        // enabled: true
        // 默认 false, true 启用生成。生成一次就可以，避免每次工程启动都生成，一旦生成配置文件之后，最好把 enable 关掉，即改成 false。
        // 否则这个文件每次会在重新加载的时候重新生成，这会导致 eslint 有时会找不到这个文件。当需要更新配置文件的时候，再重新打开
        enabled: false,
      }
    }),
  ],
  // 防止 Vite 清除 Rust 显示的错误
  clearScreen: false,
  server: {
    open: true,
    // Tauri 工作于固定端口，如果端口不可用则报错
    // strictPort: true,
    // 如果设置了 host，Tauri 则会使用
    host: "0.0.0.0" ,
    port: 1420,
    proxy: {
      '^/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
    }
    },
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'), // 使用绝对路径
    },
  },
  // 添加有关当前构建目标的额外前缀，使这些 CLI 设置的 Tauri 环境变量可以在客户端代码中访问
  envPrefix: ["VITE_", "TAURI_ENV_*"],
  build: {
    // Tauri 在 Windows 上使用 Chromium，在 macOS 和 Linux 上使用 WebKit
    target:
      process.env.TAURI_ENV_PLATFORM == "windows" ? "chrome105" : "safari13",
    // 在 debug 构建中不使用 minify
    minify: !process.env.TAURI_ENV_DEBUG ? "esbuild" : false,
    // 在 debug 构建中生成 sourcemap
    sourcemap: !!process.env.TAURI_ENV_DEBUG,
  },
});
