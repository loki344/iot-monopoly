<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import RoundButtons from '$lib/components/molecule/RoundButtons.svelte';

	export let Hst: Histoire;

	let disabled = false;
	let items = ['1', '2', '3', '4'];
</script>

<Hst.Story title="Example/Tags">
	<RoundButtons on:click={(event) => logEvent('click', event)} {items} />

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<pre>{JSON.stringify({ disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>
