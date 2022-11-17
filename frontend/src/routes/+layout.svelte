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
			console.log(eventData)
			switch(eventData.Type.split('.')[1]){
				case "TransactionAddedEvent":
				let transaction =eventData.Transaction
				goto(`transactions/${transaction.id}?recipientId=${transaction.recipientId}&senderId=${transaction.senderId}&amount=${transaction.amount}`)
				break;
				case "PropertyBuyQuestion":
				let property = eventData.Property
				goto(`properties/${property.id}`)
				break
			}


			
		};
	});
</script>

<div class="w-screen h-screen bg-blue">
	<slot />
</div>
