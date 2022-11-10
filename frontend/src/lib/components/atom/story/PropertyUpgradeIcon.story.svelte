<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import PropertyUpgradeIcon from '$components/atom/PropertyUpgradeIcon.svelte';

	export let Hst: Histoire;

	let disabled = false;
	let count = 1;
	let type = 'HOUSE';
</script>

<Hst.Story title="Example/PropertyUpgradeIcon">
	<PropertyUpgradeIcon {count} {type} on:click={(event) => logEvent('click', event)} />

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<Hst.Number bind:value={count} title="Count" />
		<Hst.Select options={['HOUSE', 'HOTEL']} bind:value={type} title="Type" />
		<pre>{JSON.stringify({ disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>
