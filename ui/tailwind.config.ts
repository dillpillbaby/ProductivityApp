import type { Config } from "tailwindcss";
import flowbitePlugin from 'flowbite/plugin'
export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],

  theme: {
    extend: {
    }
  },

  plugins: [require("@tailwindcss/typography"), flowbitePlugin]
} as Config;
