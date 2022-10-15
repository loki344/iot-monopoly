<script lang="ts">
	import { parse } from "postcss";


let playerCount = 0

</script>

<h1 class="text-3xl font-bold underline">Welcome to IoT-Monopoly</h1>
<h2>How many are playing?</h2>
<button class="py-5 px-10 my-5 block bg-gray-200" class:bg-green-200={playerCount == 1} on:click={() => playerCount = 1}>1</button>
<button class="py-5 px-10 my-5 block bg-gray-200" class:bg-green-200={playerCount == 2} on:click={() => playerCount = 2}>2</button>
<button class="py-5 px-10 my-5 block bg-gray-200" class:bg-green-200={playerCount == 3} on:click={() => playerCount = 3}>3</button>
<button class="py-5 px-10 my-5 block bg-gray-200" class:bg-green-200={playerCount == 4} on:click={() => playerCount = 4}>4</button>

<button class="py-5 px-10 my-5 block bg-green-200" on:click={() => fetch("http://localhost:3000/games", {
    method: "POST",
    body: JSON.stringify({playerCount: playerCount}),
    headers: {
                    'Content-Type': 'application/json'
                }
})}>Start a game</button>
