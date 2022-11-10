<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import Tile from '$components/wrapper/Tile.svelte';

	export let Hst: Histoire;

	let disabled = false;
</script>

<Hst.Story title="Example/Tile">
	<Tile on:click={(event) => logEvent('click', event)}><span>Das ist ein Test</span></Tile>

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<pre>{JSON.stringify({ disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>
