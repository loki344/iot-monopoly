const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		colors: {
			primary: '#fcd24f',
			dark: '#1f1347',
			blue: '#004aad'
		},
		extend: {
			fontFamily: {
				roboto: ['"Roboto Condensed"', 'sans-serif'],
				alfa: ['"Alfa Slab One"', 'sans-serif']
			}
		}
	},

	plugins: []
};

module.exports = config;
