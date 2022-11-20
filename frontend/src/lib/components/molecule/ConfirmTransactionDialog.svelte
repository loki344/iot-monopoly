<script lang="ts">
    let clazz: String = ""
    export {clazz as class};
	import { event } from "$lib/store/event";
	import Arrow from "../atom/Arrow.svelte";
	import PriceTag from "../atom/PriceTag.svelte";
	import RfidWaves from "../atom/RfidWaves.svelte";
	import Tag from "../atom/Tag.svelte";
    import GiBank from 'svelte-icons/gi/GiBank.svelte'

    export let transaction: {senderId: string, recipientId: string, amount: number}
    const transformIdToTagText = (id: string) => {

        if (id.includes("_")){
            return id.split('_')[1]
        }

        return id;
    }

    let sender = transformIdToTagText(transaction.senderId)
    let recipient = transformIdToTagText(transaction.recipientId)
</script>

<div class="px-32">
<PriceTag class={'text-center'} price={transaction.amount}></PriceTag>
<div class={`flex flex-row justify-around align-center items-center ${clazz}`}>

    {#if sender === "Bank"}
    <Tag shadow={false} active><div class="w-2"><GiBank></GiBank></div></Tag>
    {:else}
    <Tag shadow={false} active>{sender}</Tag>
    {/if}

    <Arrow class={'w-52'}></Arrow>

    {#if recipient === "Bank"}
    <Tag shadow={false} active><div class="w-20"><GiBank></GiBank></div></Tag>
    {:else}
    <Tag shadow={false} active>{recipient}</Tag>
    {/if}</div>
<div class="flex justify-center">
    <RfidWaves class={'w-80'}></RfidWaves>

</div>
</div>