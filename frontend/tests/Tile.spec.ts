import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import Tile from '$components/wrapper/Tile.svelte';

describe('Tile ', () => {
	it('should be in the document', async () => {
		const { container } = render(Tile, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
