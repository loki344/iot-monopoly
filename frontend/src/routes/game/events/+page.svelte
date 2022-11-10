<script lang="ts">
	import PropertyBuyQuestionDialog from '$lib/components/molecule/PropertyBuyQuestionDialog.svelte';
	import { event } from '$lib/store/event';

	let component: any;
	switch ($event.Type.split('.')[1]) {
		case 'PropertyBuyQuestion':
			component = PropertyBuyQuestionDialog;
			break;
		default:
			component = 'Default';
			break;
	}
</script>

<div>
	<svelte:component this={component} />
</div>
