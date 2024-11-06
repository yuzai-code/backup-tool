import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { loadEnv } from 'vite';

// 从环境变量中获取 Tauri 开发主机地址，如果没有设置则默认为 'localhost'
const host = process.env.TAURI_DEV_HOST || 'localhost';

export default defineConfig(async ({ command, mode }) => {
  // 根据当前模式加载环境变量
  const env = loadEnv(mode, process.cwd(), '');

  return {
    plugins: [vue()],

    // 针对 Tauri 开发的 Vite 选项，只在 `tauri dev` 或 `tauri build` 时应用
    clearScreen: false, // 防止 Vite 遮蔽 Rust 错误信息
    server: {
      port: 1420, // Tauri 期望一个固定的端口，如果该端口不可用则失败
      strictPort: true, // 确保端口号严格匹配
      host: host, // 使用环境变量中的主机地址，如果没有则使用 'localhost'
      hmr: host
        ? {
            protocol: 'ws',
            host,
            port: 1421, // 热更新端口号
          }
        : undefined, // 如果不在 Tauri 开发主机模式下，则不启用 HMR
      watch: {
        // 忽略 `src-tauri` 目录的文件变动，因为 Tauri 会处理这部分
        ignored: ['**/src-tauri/**'],
      },
    },

    // 定义全局变量 __APP_ENV__，用于在应用中访问环境变量
    define: {
      __APP_ENV__: JSON.stringify(env.APP_ENV),
    },
  };
});