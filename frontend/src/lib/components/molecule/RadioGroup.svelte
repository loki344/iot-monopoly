<script lang="ts">
	import Label from "../atom/Label.svelte";



    export let options: string[]
</script>

<div class="flex flex-col gap-y-40">
{#each options as option}
<div class="flex flex-row gap-x-20 content-center">
<Label _for={option}>{option}</Label>
<input type="radio" id={option} name={`optionGroup`} value={option} class="scale-[2.5]">
</div>
{/each}
</div>

