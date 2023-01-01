<script lang="ts">
	import '../app.postcss';
	import '../app.css';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	onMount(async () => {
		const socket = new WebSocket('ws://localhost:3000/ws');
		
		socket.onmessage = (backendEvent: MessageEvent<any>) => {
			const eventData = JSON.parse(backendEvent.data);
			console.log(eventData);
			

			switch(eventData.Type.split('.')[1]){

				case "TransactionCreatedEvent":
					let transaction = eventData.Transaction
					goto(`/transactions/${transaction.id}?recipientId=${transaction.recipientId}&senderId=${transaction.senderId}&amount=${transaction.amount}`)
					break

				case "PlayerOnUnownedFieldEvent":
					let p = eventData.Property
					let fd = p.financialDetails
					let r = fd.Revenue
					goto(`/properties/${p.id}?name=${p.name}&propertyPrice=${fd.PropertyPrice}&housePrice=${fd.HousePrice}&hotelPrice=${fd.HotelPrice}
					&revenueNormal=${r.Normal}&revenueOneHouse=${r.OneHouse}&revenueTwoHouses=${r.TwoHouses}&revenueThreeHouses=${r.ThreeHouses}
					&revenueFourHouses=${r.FourHouses}&revenueHotel=${r.Hotel}&buyerId=${eventData.PlayerId}`)
					break
				case "TransactionResolvedEvent":
					goto("/game")
					break
				case "CardDrewEvent":
					goto(`/card-event?title=${eventData.Title}&text=${eventData.Text}`)
					break
			}


			
		};
	});
</script>

<div class="w-screen h-screen bg-blue">
	<slot />
</div>
