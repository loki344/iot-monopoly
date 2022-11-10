<script lang="ts">
	import GiFamilyHouse from 'svelte-icons/gi/GiFamilyHouse.svelte';
	import GiHouse from 'svelte-icons/gi/GiHouse.svelte';

	let clazz: String = '';
	export { clazz as class };
	export let type = 'HOUSE';
	export let count = 1;
</script>

<div class={`flex flex-row justify-center gap-x-2 ${clazz}`}>
	{#each { length: count } as _}
		<div class="text-dark w-10">
			{#if type === 'HOUSE'}
				<GiHouse />
			{:else}
				<GiFamilyHouse />
			{/if}
		</div>
	{/each}
</div>
