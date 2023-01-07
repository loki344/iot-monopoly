<script lang="ts">
	import { goto } from "$app/navigation";
	import Button from "../atom/Button.svelte";
	import Title from "../atom/Title.svelte";
	import Tile from "../wrapper/Tile.svelte";

    let clazz: String = ""
    export {clazz as class};
    export let title = ""
    export let text = ""
    export let onClick = async () => {}
</script>

<div class={`${clazz} flex flex-col items-center h-full p-16`}>
    <Tile class="!border-green !h-3/4 !p-16 !max-w-[50%]">
        <Title class="!text-dark underline !text-center">{title}</Title>
        <Title class="!text-dark" type="small">{text}</Title>
    </Tile>
    <Button class="mt-16" type="POSITIVE" onClick={ async () =>{
        onClick()
        goto("/game")
    }}>OK</Button>
</div>

