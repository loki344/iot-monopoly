<script lang="ts">
	import { BASE_URL, extractData } from "$lib/http/backendClient";
	import type { Property } from "$lib/model/PropertyDetail";
	import { event } from "$lib/store/event";
	import Button from "../atom/Button.svelte";
	import PropertyDetails from "./PropertyDetails.svelte";
	let clazz: String = '';
	export { clazz as class };
    export let property:Property

	const decline = () => {

	}

	const accept = async () => {

		let response = await fetch(`${BASE_URL}/transactions`, {
			method: 'POST',
			body: JSON.stringify({"senderId": $event.PlayerId, "recipientId": "bank", "amount": $event.Property.FinancialDetails.PropertyPrice}),
			headers: {
				'Content-Type': 'application/json'
			}
		});

		return await extractData(response);
	}
</script>

<div class={`${clazz}`}>
<PropertyDetails property={$event.Property}></PropertyDetails>
<div class="mt-8 flex flex-row  justify-around">
    <Button onClick={decline} type="NEGATIVE">X</Button><Button onClick={accept} type="POSITIVE">BUY</Button>
</div></div>
