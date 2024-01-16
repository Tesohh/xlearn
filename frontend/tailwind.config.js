/** @type {import('tailwindcss').Config} */
const { addDynamicIconSelectors } = require('@iconify/tailwind');
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		colors: {
			primary: '#3963A3',
			notSelected: '#889BB9',
			background: '#FFFFFF'
		},
		extend: {}
	},
	plugins: [
		addDynamicIconSelectors({
			iconSets: 'tabler'
		})
	]
};
