<script lang="ts">
	import '../app.postcss';
	import '../app.css';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { event } from '$lib/store/event';

	onMount(async () => {
		const socket = new WebSocket('ws://localhost:3000/ws');
		socket.onmessage = (backendEvent: MessageEvent<any>) => {
			const eventData = JSON.parse(backendEvent.data);
			event.set(eventData);
			goto('/game/events');
		};
	});
</script>

<div class="w-screen h-screen bg-blue">
	<slot />
</div>
