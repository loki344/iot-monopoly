<script lang="ts">
    let clazz: String = ""
    export {clazz as class};
	import Arrow from "../atom/Arrow.svelte";
	import PriceTag from "../atom/PriceTag.svelte";
	import RfidWaves from "../atom/RfidWaves.svelte";
	import RoundButton from "../atom/RoundButton.svelte";
    import GiBank from 'svelte-icons/gi/GiBank.svelte'
	import Title from "../atom/Title.svelte";
    import Text from "../atom/Text.svelte";

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


<div class="h-full flex flex-col justify-around text-center">

<Title class="!text-dark">Transaction</Title>
<PriceTag price={transaction.amount}></PriceTag>
<div class={`flex flex-row w-1/3 justify-between align-center items-center self-center ${clazz}`}>

    {#if sender === "Bank"}
    <RoundButton shadow={false} big={true} active><div class="w-2"><GiBank></GiBank></div></RoundButton>
    {:else}
    <RoundButton shadow={false} big={true} active><Text>{sender}</Text></RoundButton>
    {/if}

    <Arrow class={'w-80'}></Arrow>

    {#if recipient === "Bank"}
    <RoundButton shadow={false} big={true} active><div class="w-28"><GiBank></GiBank></div></RoundButton>
    {:else}
    <RoundButton shadow={false} big={true} active>{recipient}</RoundButton>
    {/if}</div>
<div class="flex justify-center">
    <RfidWaves class={'w-80'}></RfidWaves>
</div></div>