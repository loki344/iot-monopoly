import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import RfidWaves from '$components/atom/RfidWaves.svelte';

describe('RfidWaves ', () => {
	it('should be in the document', async () => {
		const { container } = render(RfidWaves, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
})
