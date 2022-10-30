import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import { Button } from '../src/lib/';

describe('Button ', () => {
	it('should be in the document', async () => {
		const { container } = render(Button, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
