import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import Text from '$components/atom/Text.svelte';

describe('Text ', () => {
	it('should be in the document', async () => {
		const { container } = render(Text, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
