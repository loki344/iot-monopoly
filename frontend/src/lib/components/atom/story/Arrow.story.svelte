<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import Arrow from '$components/atom/Arrow.svelte';

	export let Hst: Histoire;

	let disabled = false

</script>

<Hst.Story title="Example/Arrow">
	<Arrow on:click={(event) => logEvent('click', event)} />

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<pre>{JSON.stringify({disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>