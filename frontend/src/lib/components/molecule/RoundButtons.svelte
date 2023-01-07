<script lang="ts">
	import RoundButton from '../atom/RoundButton.svelte';

	export let items: String[] = [];
	export let value: String = '';
</script>

<div class="flex flex-row">
	{#each items as item}
		<RoundButton
			class="mx-20"
			active={value === item}
			immediateOnClick={() => {
				if (value === item) {
					value = '';
				} else {
					value = item;
				}
			}}>{item}</RoundButton
		>
	{/each}
</div>
