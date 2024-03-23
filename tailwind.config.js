const neutral = {
  50: '#f9f8f7',
  100: '#f1f0ef',
  200: '#e7e4e1',
  300: '#d6d1cd',
  400: '#bcb4ad',
  500: '#a39890',
  600: '#8b8077',
  700: '#736a62',
  800: '#615953',
  900: '#534d49',
  950: '#2b2724',
}

const primary = {
  50: '#f5f2f1',
  100: '#e5e0dc',
  200: '#ccc4bc',
  300: '#aea196',
  400: '#978578',
  500: '#897469',
  600: '#756159',
  700: '#5e4d49',
  800: '#524341',
  900: '#483c3b',
  950: '#282020',
}
import colors from "tailwindcss/colors"

/** @type {import('tailwindcss').Config} */
export default {
  content: ["views/**/*.jet", "static/css/global.css"], 
  theme: {
    extend: {
        colors: {
          neutral, 
          primary,
          accent:colors.yellow,
        "background":neutral[100],
        "foreground":neutral[950],
        "background-secondary":neutral[300],
        "foreground-secondary":neutral[800],
        "background-muted":neutral[50],
        "foreground-muted":neutral[600],
        "border":neutral[300],
        "danger":colors.red[500]
        },
      },
  },
    plugins: [],
  }
  
  