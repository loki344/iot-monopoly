<script lang="ts">
    let clazz: String = ""
    export {clazz as class};
	import { event } from "$lib/store/event";
	import Arrow from "../atom/Arrow.svelte";
	import PriceTag from "../atom/PriceTag.svelte";
	import RfidWaves from "../atom/RfidWaves.svelte";
	import Tag from "../atom/Tag.svelte";


    export let transaction: {senderId: string, receiverId: string, amount: number}
</script>

<div class="px-32">
<PriceTag class={'text-center'} price={transaction.amount}></PriceTag>
<div class={`flex flex-row justify-around align-center items-center ${clazz}`}>

    <Tag shadow={false} active>{transaction.receiverId.split('_')[1]}</Tag>
    <Arrow class={'w-52'}></Arrow>
    <Tag shadow={false} active>{transaction.senderId.split('_')[1]}</Tag>
</div>
<div class="flex justify-center">
    <RfidWaves class={'w-80'}></RfidWaves>

</div>
</div>