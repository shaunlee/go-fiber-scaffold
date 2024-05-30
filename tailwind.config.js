module.exports = {
  content: ['./views/**/*.html'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        primary: '#8c52ff',
        secondary: '#ff914d'
      },
      aspectRatio: {
        pdf: '3 / 4'
      }
    },
  },
  variants: {
    extend: {
      opacity: ['disabled'],
      cursor: ['disabled'],
    },
  },
  plugins: [],
}
