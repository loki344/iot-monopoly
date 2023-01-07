<script>
	import { goto } from "$app/navigation";
	import Button from "$lib/components/atom/Button.svelte";
	import { BASE_URL } from "$lib/http/backendClient";

    const endGame = async () => {
        await fetch(`${BASE_URL}/games/current/`, {
			method: 'PATCH',
			body: JSON.stringify({ ended: true }),
			headers: {
				'Content-Type': 'application/json'
			}
		});
        goto('/')
    }

</script>

<div class="flex flex-col text-center justify-center space-y-12 h-full">
    <Button class="mb-12" onClick={()=> endGame()}>End Game</Button>
    <Button onClick={()=>goto('/game')}>Back to Game</Button>
</div>