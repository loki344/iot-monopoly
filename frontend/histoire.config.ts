import { defineConfig } from 'histoire';
import { HstSvelte } from '@histoire/plugin-svelte';

export default defineConfig({
	setupFile: '/histoire-setup.ts',
	plugins: [HstSvelte()],
	storyMatch: ['**/*.story.svelte'],

	tree: {
		groups: [
			{
				id: 'top',
				title: ''
			}
		]
	}
});
