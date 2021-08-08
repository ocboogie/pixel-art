import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  optimizeDeps: {
    exclude: ["@roxi/routify", "@urql/svelte", "@zerodevx/svelte-toast"],
  },
  define: {
    "process.env": { NODE_ENV: process.env.NODE_ENV },
  },
});
