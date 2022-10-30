<script lang="ts">
	import Tag from '../atom/Tag.svelte';

	export let items: String[] = [];
</script>

<div class="flex flex-row">
	{#each items as item}
		<Tag class="mx-20">{item}</Tag>
	{/each}
</div>
