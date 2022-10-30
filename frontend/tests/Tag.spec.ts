import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import Tag from '$components/atom/Tag.svelte';

describe('Tag ', () => {
	it('should be in the document', async () => {
		const { container } = render(Tag, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
