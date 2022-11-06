import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import Tags  from '$components/molecule/Tags.svelte';

describe('Tags ', () => {
	it('should be in the document', async () => {
		const { container } = render(Tags, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
