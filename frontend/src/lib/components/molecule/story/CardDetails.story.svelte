<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import CardDetails from '$components/molecule/CardDetails.svelte';

	export let Hst: Histoire;

	let disabled = false
	let title = "This is a title"
	let text = "This is a text, so it's a bit longer and you can use it to inform in depth about what's happening here.."

</script>

<Hst.Story title="Example/CardDetails">
	<CardDetails on:click={(event) => logEvent('click', event)} text={text} title={title} />

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<Hst.Text bind:value={title} title="Title" />
		<Hst.Text bind:value={text} title="Text" />
		<pre>{JSON.stringify({disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>