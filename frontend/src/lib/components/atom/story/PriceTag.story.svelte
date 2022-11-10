<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import PriceTag from '$components/atom/PriceTag.svelte';

	export let Hst: Histoire;

	let disabled = false;
	let price = 1;
</script>

<Hst.Story title="Example/PriceTag">
	<PriceTag {price} on:click={(event) => logEvent('click', event)} />

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<Hst.Number bind:value={price} title="Price" />
		<pre>{JSON.stringify({ disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>
