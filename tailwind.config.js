/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'spring-green': '#5bf999',
        'accent': '#61AFEF',
        'dark': '#1E2127',
        'darker': '#181A1F',
        'lighter': '#282C34'
      }
    },
  },
  plugins: [],
}
