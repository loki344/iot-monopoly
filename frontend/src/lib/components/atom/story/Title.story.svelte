<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import Title from '$components/atom/Title.svelte';

	export let Hst: Histoire;

	let disabled = false;
</script>

<Hst.Story title="Example/Title">
	<Title on:click={(event) => logEvent('click', event)}>This is a title</Title>

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<pre>{JSON.stringify({ disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>
