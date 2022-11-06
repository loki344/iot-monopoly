<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import Tag from '$components/atom/Tag.svelte';

	export let Hst: Histoire;

	let disabled = false;
</script>

<Hst.Story title="Example/Tag">
	<Tag on:click={(event) => logEvent('click', event)}>1</Tag>

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<pre>{JSON.stringify({ disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>
