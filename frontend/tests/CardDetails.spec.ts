import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import CardDetails from '$components/molecule/CardDetails.svelte';

describe('CardDetails ', () => {
	it('should be in the document', async () => {
		const { container } = render(CardDetails, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
})
