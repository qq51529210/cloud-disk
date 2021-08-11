import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
const { resolve } = require("path");
import ViteComponents, { VarletUIResolver } from "vite-plugin-components";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    ViteComponents({
      globalComponentsDeclaration: true,
      customComponentResolvers: [VarletUIResolver()]
    })
  ],
  build: {
    rollupOptions: {
      input: {
        sign_in: resolve(__dirname, "sign_in/index.html"),
        sign_up: resolve(__dirname, "sign_up/index.html")
      }
    }
  },
  resolve: {
    alias: [{ find: "vue-i18n", replacement: "vue-i18n/dist/vue-i18n.cjs.js" }]
  }
  // server: {
  //   proxy: {
  //     "/api": {
  //       target: "http://127.0.0.1:33966/api",
  //       changeOrigin: true
  //     }
  //   }
  // }
});
