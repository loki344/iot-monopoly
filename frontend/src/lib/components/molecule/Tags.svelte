<script lang="ts">
	import Tag from '../atom/Tag.svelte';

	export let items: String[] = [];
	export let value: String = '';
</script>

<div class="flex flex-row">
	{#each items as item}
		<Tag
			class="mx-20"
			active={value === item}
			immediateOnClick={() => {
				if (value === item) {
					value = '';
				} else {
					value = item;
				}
			}}>{item}</Tag
		>
	{/each}
</div>
