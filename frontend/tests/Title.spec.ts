import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import { Title } from '../src/lib/';

describe('Title ', () => {
	it('should be in the document', async () => {
		const { container } = render(Title, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
