import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import PriceTag from '$components/atom/PriceTag.svelte';

describe('PriceTag ', () => {
	it('should be in the document', async () => {
		const { container } = render(PriceTag, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
