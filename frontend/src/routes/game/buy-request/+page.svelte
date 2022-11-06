<script lang="ts">
	import Button from "$lib/components/atom/Button.svelte";
	import Title from "$lib/components/atom/Title.svelte";
	import PropertyDetails from "$lib/components/molecule/PropertyDetails.svelte";




</script>

<div class="flex flex-col items-center justify-around h-screen">
	<Title>BUY OPPORTUNITY</Title>
	<!-- TODO fetch from backend -->
	<PropertyDetails propertyDetail={{name:"RandomName", price: 5000, revenue:5000, bigHouseCost: 50000, smallHouseCost: 9000}}></PropertyDetails>

	<div class="flex flex-row w-[80%] justify-around">
		<Button type="NEGATIVE">X</Button><Button type="POSITIVE">BUY</Button>
	</div>
</div>
