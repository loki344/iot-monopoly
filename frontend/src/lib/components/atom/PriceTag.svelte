<script lang="ts">
    let clazz: String = ""
    export {clazz as class};
    import IoMdPricetag from 'svelte-icons/io/IoMdPricetag.svelte'
	import Title from './Title.svelte';
    export let price:number
    let digits = price > 0 ? Math.ceil(price / 10).toString.length : 1
    $: positionCorrection = digits * 5
    $: topPosition = (90 + positionCorrection) + "px"

</script>


<div class="{`${clazz} text-primary w-72`}">
    <IoMdPricetag></IoMdPricetag>
</div>
<div class={`-rotate-45 absolute top-[${topPosition}] left-[150px]`}>
    <Title class="!text-dark text-2xl tracking-wider font-extrabold">{price} $</Title>
</div>
{digits}
{positionCorrection}
{topPosition}

