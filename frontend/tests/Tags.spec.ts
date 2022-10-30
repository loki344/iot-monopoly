import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import { Tags } from '../src/lib/';

describe('Tags ', () => {
	it('should be in the document', async () => {
		const { container } = render(Tags, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
