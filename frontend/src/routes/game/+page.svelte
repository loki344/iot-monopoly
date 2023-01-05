<script lang="ts">
	import Tile from '$lib/components/wrapper/Tile.svelte';
	import Text from '$lib/components/atom/Text.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
	let game = JSON.parse(data.game)
	console.log(game);
	
	const getClass = (index: Number): String => {

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
	{#each game.players as player, i}
		<Tile active={game.currentPlayerIndex === i} class={`absolute ${getClass(i)}`}>
			<Text>Player {i + 1}</Text><Text class="text-center">{player.account.balance} $</Text>
		</Tile>
	{/each}
</div>
