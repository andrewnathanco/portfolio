/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js}"],
  theme: {
    extend: {
          backgroundImage: {
            'profile-picture': "url('/static/img/profile-picture.png')",
          },
          spacing: {
            '128': '32rem',
            '144': '36rem',
          },
          colors: {
            'paris-white': {
              '50': '#f2f7f6',
              '100': '#d1e3dd',
              '200': '#c2d8d2',
              '300': '#98bdb5',
              '400': '#6b9c93',
              '500': '#4b7e76',
              '600': '#38635e',
              '700': '#2c504c',
              '800': '#25403d',
              '900': '#1f3533',
              '950': '#111d1c',
            },
            'scarpa-flow': {
              '50': '#f7f7f8',
              '100': '#eeedf1',
              '200': '#d9d8df',
              '300': '#b8b6c3',
              '400': '#918ea2',
              '500': '#747087',
              '600': '#575366',
              '700': '#4e4a5a',
              '800': '#423f4d',
              '900': '#3b3842',
              '950': '#27252c',
          },
          'japonica': {
            '50': '#fcf6f4',
            '100': '#f9ebe7',
            '200': '#f5dad3',
            '300': '#edc0b4',
            '400': '#e09b89',
            '500': '#d3806a',
            '600': '#bc5f46',
            '700': '#9d4d38',
            '800': '#834231',
            '900': '#6e3b2e',
            '950': '#3b1c14',
        },
      }
    },
  },
  plugins: [],
}

