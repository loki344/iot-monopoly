<script lang="ts">
	import { goto } from "$app/navigation";
	import { BASE_URL, extractData } from "$lib/http/backendClient";
	import type { Property } from "$lib/model/PropertyDetail";
	import Button from "../atom/Button.svelte";
	import PropertyDetails from "./PropertyDetails.svelte";
	let clazz: String = '';
	export { clazz as class };
    export let property:Property
    export let buyerId:string

	const decline = () => {
		goto("/game")
	}

	const accept = async () => {

		let response = await fetch(`${BASE_URL}/games/current/properties/${property.Id}`, {
			method: 'PATCH',
			body: JSON.stringify({"ownerId": buyerId}),
			headers: {
				'Content-Type': 'application/json'
			}
		});

		return await extractData(response);
	}
</script>

<div class={`${clazz}`}>
<PropertyDetails property={property}></PropertyDetails>
<div class="mt-8 flex flex-row  justify-around">
    <Button onClick={decline} type="NEGATIVE">X</Button><Button onClick={accept} type="POSITIVE">BUY</Button>
</div></div>
