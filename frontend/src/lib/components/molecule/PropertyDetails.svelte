<script lang="ts">
	import type { PropertyDetail } from "$lib/model/PropertyDetail";
	import Title from "../atom/Title.svelte";
    import GiHouse from 'svelte-icons/gi/GiHouse.svelte'
    import GiFamilyHouse from 'svelte-icons/gi/GiFamilyHouse.svelte'
	import PriceTag from "../atom/PriceTag.svelte";

    let clazz: String = ""
    export {clazz as class};
    export let propertyDetail: PropertyDetail;
</script>

<div class={`${clazz}`}>
	<Title type="medium" class='mb-6 text-center'>{propertyDetail.name}</Title>
    <PriceTag price={propertyDetail.price}></PriceTag>
    <div class="grid grid-cols-3 gap-x-16">
        <div></div>
        <div class="flex justify-center"><div class="text-dark w-24 text-center"><GiHouse></GiHouse></div></div>
        <div class="flex justify-center"><div class="text-dark w-24 mb-8"><GiFamilyHouse></GiFamilyHouse></div></div>

       
        <Title type="small">REVENUE:</Title>
        <Title class="cell" type="small">{propertyDetail.revenue} $</Title>
        <Title class="cell" type="small">{propertyDetail.revenue} $</Title>

        <Title  type="small">COSTS:</Title>
        <Title class="cell" type="small">{propertyDetail.smallHouseCost} $</Title>
        <Title class="cell" type="small">{propertyDetail.bigHouseCost} $</Title>

    </div>
    
</div>