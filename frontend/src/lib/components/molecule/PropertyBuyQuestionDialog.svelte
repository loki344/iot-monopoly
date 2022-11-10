<script lang="ts">
	import type { Property } from "$lib/model/PropertyDetail";
	import { event } from "$lib/store/event";
	import Button from "../atom/Button.svelte";
	import PropertyDetails from "./PropertyDetails.svelte";
	let clazz: String = '';
	export { clazz as class };
    export let property:Property
    event.subscribe((value)=> console.log(value))
    
</script>

<div class={`${clazz}`}>
<PropertyDetails property={$event.Property}></PropertyDetails>
<div class="mt-8 flex flex-row  justify-around">
    <Button type="NEGATIVE">X</Button><Button type="POSITIVE">BUY</Button>
</div></div>
