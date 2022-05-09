module.exports = {
	daisyui: {
		themes: [
			{
				dark: {
					primary: '#60a5fa',
					secondary: '#3730a3',
					accent: '#8a88fc',
					neutral: '#1D1929',
					'base-100': '#293042',
					info: '#A2C8F1',
					success: '#6EEDBA',
					warning: '#E7A01D',
					error: '#EC4B7E'
				}
			}
		]
	},
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {}
	},
	plugins: [require('@tailwindcss/typography'), require('daisyui')]
};
