<script lang="ts">
	import '../app.postcss';
	import '../app.css';
	import { onMount } from 'svelte';

	onMount(async () => {
		const socket = new WebSocket('ws://localhost:3000/ws');
	socket.onmessage = (event: MessageEvent<any>) => {
		console.log(event)
		return event
	}
	});




</script>

<div class="w-screen h-screen bg-blue">
	<slot />
</div>
