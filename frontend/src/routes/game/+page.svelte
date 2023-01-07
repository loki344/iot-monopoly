<script lang="ts">
	import Tile from '$lib/components/wrapper/Tile.svelte';
	import Text from '$lib/components/atom/Text.svelte';
	import type { PageData } from './$types';
	import TiLocation from 'svelte-icons/ti/TiLocation.svelte'
	import GiCardJoker from 'svelte-icons/gi/GiCardJoker.svelte'
	import { goto } from '$app/navigation';
	import Button from '$lib/components/atom/Button.svelte';


	export let data: PageData;
	let game = JSON.parse(data.game)
	
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
	<Button class="text-center absolute left-[40%] top-10" onClick={() => goto('/menu')}>
		menu
	</Button>
	{#each game.players as player, i}
		<Tile active={game.currentPlayerIndex === i} class={`absolute ${getClass(i)}`}>
			<Text>Player {i + 1}</Text><Text class="text-center">{player.account.balance} $</Text>
			<div class="flex flex-row justify-around">
				<div class="flex flex-row">
					<div class="w-16">
						<TiLocation></TiLocation>
					</div>
				<Text class="inline">{player.propertyCount}</Text>

				</div>
				<div class="flex flex-row">
				<div class="w-16">
					<GiCardJoker></GiCardJoker>
				</div>
				<Text class="inline">{player.escapeFromPrisonCardCount}</Text>
				</div>
			</div>
		</Tile>
	{/each}
</div>
