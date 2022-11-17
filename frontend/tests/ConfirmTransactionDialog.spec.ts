import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import ConfirmTransactionDialog from '$components/molecule/ConfirmTransactionDialog.svelte';

describe('ConfirmTransactionDialog ', () => {
	it('should be in the document', async () => {
		const { container } = render(ConfirmTransactionDialog, {
			props: {}
		});
		expect(container).toBeInTheDocument();
	});
})
