import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import PropertyDetails from '$components/molecule/PropertyDetails.svelte';

describe('PropertyDetails ', () => {
	it('should be in the document', async () => {
		const { container } = render(PropertyDetails, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
})
