import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import Arrow from '$components/atom/Arrow.svelte';

describe('Arrow ', () => {
	it('should be in the document', async () => {
		const { container } = render(Arrow, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
})
