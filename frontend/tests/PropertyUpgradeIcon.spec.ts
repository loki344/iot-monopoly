import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import PropertyUpgradeIcon from '$components/atom/PropertyUpgradeIcon.svelte';

describe('PropertyUpgradeIcon ', () => {
	it('should be in the document', async () => {
		const { container } = render(PropertyUpgradeIcon, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
