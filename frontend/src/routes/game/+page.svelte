<script lang="ts">
	import Tile from '$lib/components/wrapper/Tile.svelte';
	import Text from '$lib/components/atom/Text.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
	let players = JSON.parse(data.players);

	const getClass = (index: Number): String => {
		console.log(index);

		switch (index) {
			case 0:
				return 'top-10 left-10';
			case 1:
				return 'top-10 right-10';
			case 2:
				return 'bottom-10 left-10';
			case 3:
				return 'bottom-10 right-10';
			default:
				return '';
		}
	};
</script>

<div>
	{#each players as player, i}
		<!-- TODO determine which players turn it is and set the others active = false -->
		<Tile active={true} class={`absolute ${getClass(i)}`}
			><Text>Player {i + 1}</Text><Text class="text-center">{player.balance} $</Text></Tile
		>
	{/each}
</div>
