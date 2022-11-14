/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      'background': '#33333d',
      'foreground': '#3f3e48',
      'lgreen': '#1eb980', // light green
      'dgreen': '#005d57', // dark green
      'lgray': '#7f7f85',  // light gray
      'white': '#fff',   // white
      'lblack': '#1a1a1a',   // light black
      'lred': '#ff937f'  // light red
    },
    extend: {},
  },
  plugins: [],
}

