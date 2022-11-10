<script lang="ts">
	// the alias fixes lint error 'Hst' is already defined
	import type { Hst as Histoire } from '@histoire/plugin-svelte';
	import { logEvent } from 'histoire/client';
	import PropertyDetails from '$components/molecule/PropertyDetails.svelte';
	import { PropertyUpgrades } from '$lib/model/PropertyDetail';

	export let Hst: Histoire;

	let disabled = false;
</script>

<Hst.Story title="Example/PropertyDetails">
	<PropertyDetails
		property={{
			Name: 'TestProperty',
			Id: 'abcd-wer-we93',
			OwnerId: 'ownerId',
			FinancialDetails: {
				PropertyPrice: 1000,
				HousePrice: 200,
				HotelPrice: 500,
				Revenue: {
					Normal: 100,
					OneHouse: 200,
					TwoHouses: 250,
					ThreeHouses: 300,
					FourHouses: 400,
					Hotel: 500
				}
			},
			Upgrades: PropertyUpgrades.TWO_HOUSES
		}}
		on:click={(event) => logEvent('click', event)}
	/>

	<svelte:fragment slot="controls">
		<Hst.Checkbox bind:value={disabled} title="Disabled" />
		<pre>{JSON.stringify({ disabled }, null, 2)}</pre>
	</svelte:fragment>
</Hst.Story>
