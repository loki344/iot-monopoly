const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			fontFamily: {
				roboto: ['"Roboto Condensed"', 'sans-serif'],
				alfa: ['"Alfa Slab One"', 'sans-serif']
			},
			colors: {
				primary: '#fcd24f',
				dark: '#1f1347',
				blue: '#004aad',
				gray: '#9ca3af'
			},
		}
	},

	plugins: []
};

module.exports = config;
