<script lang="ts">
	import type { Property } from '$lib/model/PropertyDetail';
	import Title from '../atom/Title.svelte';
	import PriceTag from '../atom/PriceTag.svelte';
	import PropertyUpgradeIcon from '$components/atom/PropertyUpgradeIcon.svelte';

	let clazz: String = '';
	export { clazz as class };
	export let property: Property;
</script>

<div class={`${clazz} flex flex-col`}>
	<Title type="medium" class="text-center">{property.Name}</Title>
	<PriceTag class="text-center" price={property.FinancialDetails.PropertyPrice} />
	<Title type="small" class="text-center">REVENUE</Title>
	<div class="grid grid-cols-6 grid-rows-2 gap-x-12">
		<div />
		<PropertyUpgradeIcon type="HOUSE" count={1} />
		<PropertyUpgradeIcon type="HOUSE" count={2} />
		<PropertyUpgradeIcon type="HOUSE" count={3} />
		<PropertyUpgradeIcon type="HOUSE" count={4} />
		<PropertyUpgradeIcon type="HOTEL" count={1} />

		<Title class="text-center text-green" type="small"
			>{property.FinancialDetails.Revenue.Normal} $</Title
		>
		<Title class="text-center text-green" type="small"
			>{property.FinancialDetails.Revenue.OneHouse} $</Title
		>
		<Title class="text-center text-green" type="small"
			>{property.FinancialDetails.Revenue.TwoHouses} $</Title
		>
		<Title class="text-center text-green" type="small"
			>{property.FinancialDetails.Revenue.ThreeHouses} $</Title
		>
		<Title class="text-center  text-green" type="small"
			>{property.FinancialDetails.Revenue.FourHouses} $</Title
		>
		<Title class="text-center text-green" type="small"
			>{property.FinancialDetails.Revenue.Hotel} $</Title
		>
	</div>

    <Title type="small" class="text-center">COST</Title>
	<div class="flex flex-row justify-center">
		<div class="grid grid-cols-2 w-1/2">
			<PropertyUpgradeIcon type="HOUSE" count={1} />
			<PropertyUpgradeIcon type="HOTEL" count={1} />

			<Title class="text-center text-green" type="small"
				>{property.FinancialDetails.HousePrice} $</Title
			>
			<Title class="text-center text-green" type="small"
				>{property.FinancialDetails.HotelPrice} $</Title
			>
		</div>
	</div>
</div>
