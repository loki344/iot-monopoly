<script lang="ts">
	import { goto } from "$app/navigation";
	import Button from "$lib/components/atom/Button.svelte";




const startGame = async () => {
    let response = await fetch("http://localhost:3000/games", {
    method: "POST",
    body: JSON.stringify({playerCount: playerCount}),
    headers: {
                    'Content-Type': 'application/json'
                }

})
let players = await response.json()}



let playerCount = 0

</script>

<h2 class="text-md">How many are playing?</h2>
<div class="flex flex-col justify-center grow gap-y-10">
    <Button label={2} onClick={() => playerCount = 2} active={playerCount === 2}></Button>
    <Button label={3} onClick={() => playerCount = 3} active={playerCount === 3}></Button>
    <Button label={4} onClick={() => playerCount = 4} active={playerCount === 4}></Button>
    <Button label="Start game" onClick={async () => {
        await startGame();
        goto('/game')
    }}></Button>
</div>
