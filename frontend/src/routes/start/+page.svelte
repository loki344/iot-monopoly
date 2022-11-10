<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/atom/Button.svelte';

	import Title from '$lib/components/atom/Title.svelte';
	import Tags from '$lib/components/molecule/Tags.svelte';
	import { BASE_URL, extractData } from '$lib/http/backendClient';

	async function startGame(playerCount: Number) {
		let response = await fetch(`${BASE_URL}/games`, {
			method: 'POST',
			body: JSON.stringify({ playerCount: playerCount }),
			headers: {
				'Content-Type': 'application/json'
			}
		});

		return await extractData(response);
	}

	let playerCount = '';
</script>

<div class="flex flex-col items-center justify-around h-screen">
	<Title type="medium">HOW MANY ARE PLAYING?</Title>
	<Tags bind:value={playerCount} items={['1', '2', '3', '4']} />
	<Button
		class={!playerCount ? 'invisible' : ''}
		onClick={async () => {
			await startGame(Number.parseInt(playerCount));
			goto('/game');
		}}
		disabled={!playerCount}>Start</Button
	>
</div>
