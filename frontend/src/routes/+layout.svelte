<script lang="ts">
	import '../app.postcss';
	import '../app.css';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	onMount(async () => {
		const socket = new WebSocket('ws://localhost:3000/ws');
		
		socket.onmessage = (backendEvent: MessageEvent<any>) => {
			const eventData = JSON.parse(backendEvent.data);		
			switch(eventData.Type.split('.')[1]){

				case "TransactionCreatedEvent":
					let transaction = eventData
					goto(`/transactions/${transaction.Id}?recipientId=${transaction.RecipientId}&
					senderId=${transaction.SenderId}&amount=${transaction.Amount}`)
					break
				case "PlayerOnUnownedFieldEvent":
					goto(`/properties/${eventData.PropertyIndex}?name=${eventData.PropertyName}
					&propertyPrice=${eventData.PropertyPrice}&housePrice=${eventData.HousePrice}
					&hotelPrice=${eventData.HotelPrice}&revenueNormal=${eventData.RevenueNormal}
					&revenueOneHouse=${eventData.RevenueOneHouse}&revenueTwoHouses=${eventData.RevenueTwoHouses}
					&revenueThreeHouses=${eventData.RevenueThreeHouses}&revenueFourHouses=${eventData.RevenueFourHouses}
					&revenueHotel=${eventData.RevenueHotel}&buyerId=${eventData.PlayerId}`)
					break
				case "TransactionResolvedEvent":
					goto("/game")
					break
				case "CardDrewEvent":
					goto(`/card-event?title=${eventData.Title}&text=${eventData.Text}`)
					break
				case "PlayerDataUpdatedEvent":
				case "AccountDataUpdatedEvent":
					goto("/game")
					break
			}
		};
	});
</script>

<div class="w-screen h-screen bg-blue">
	<slot />
</div>
