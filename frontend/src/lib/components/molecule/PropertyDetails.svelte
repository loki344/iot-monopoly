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
	<Title class="text-center !text-dark">{property.Name}</Title>
	<PriceTag class="text-center" price={property.FinancialDetails.PropertyPrice} />

	<div class="flex flex-row my-12 divide-x-4 divide-dark">
		<div class="w-1/2">
			<Title type="medium" class="text-center">REVENUE</Title>
			<div class="grid grid-cols-4 grid-rows-3">
				<Title type="small" class="text-center !text-dark">Empty</Title>
				<Title class="text-center !text-green" type="small"
				>{property.FinancialDetails.Revenue.Normal} $</Title>
		
				<PropertyUpgradeIcon type="HOUSE" count={3} />
				<Title class="text-center !text-green" type="small"
				>{property.FinancialDetails.Revenue.ThreeHouses} $</Title>

				<PropertyUpgradeIcon type="HOUSE" count={1} />
				<Title class="text-center !text-green" type="small"
				>{property.FinancialDetails.Revenue.OneHouse} $</Title>
		
				<PropertyUpgradeIcon type="HOUSE" count={4} />
				<Title class="text-center  !text-green" type="small"
					>{property.FinancialDetails.Revenue.FourHouses} $</Title>
		
					<PropertyUpgradeIcon type="HOUSE" count={2} />
					<Title class="text-center !text-green" type="small">{property.FinancialDetails.Revenue.TwoHouses} $</Title>

				<PropertyUpgradeIcon type="HOTEL" count={1} />
				<Title class="text-center !text-green" type="small"
					>{property.FinancialDetails.Revenue.Hotel} $</Title>
		</div>
	</div>


	<div class="w-1/2">
		<Title type="medium" class="text-center !text-primary">COST</Title>
		<div class="grid grid-cols-4 grid-rows-3">
			<PropertyUpgradeIcon type="HOUSE" count={1} />
			<Title class="text-center !text-red" type="small"
			>{property.FinancialDetails.HousePrice} $</Title
		>
			<PropertyUpgradeIcon type="HOTEL" count={1} />
			<Title class="text-center !text-red" type="small"
				>{property.FinancialDetails.HotelPrice} $</Title
			>
	</div>

	</div>


</div>
</div>
