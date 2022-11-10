import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import PropertyBuyQuestionDialog from '$components/molecule/PropertyBuyQuestionDialog.svelte';

describe('PropertyBuyQuestionDialog ', () => {
	it('should be in the document', async () => {
		const { container } = render(PropertyBuyQuestionDialog, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
});
