<script lang="ts">
	import Tile from '$lib/components/wrapper/Tile.svelte';
	import Text from '$lib/components/atom/Text.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
	
	let players = JSON.parse(data.gameData).players;
	let currentPlayerId = JSON.parse(data.gameData).currentPlayerId;
	let accounts = JSON.parse(data.accounts).reduce((map:any, account:any) => (map[account.id] = account.balance, map), {});

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
	{#each players as player, i}
		<Tile active={currentPlayerId === player.id} class={`absolute ${getClass(i)}`}>
			<Text>Player {i + 1}</Text><Text class="text-center">{accounts[player.accountId]} $</Text>
		</Tile>
	{/each}
</div>
