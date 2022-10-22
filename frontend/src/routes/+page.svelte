<script>
	import { goto } from "$app/navigation";
	import Button from "$lib/components/atom/Button.svelte";

</script>

<h1 class="text-xl">Welcome to IoT-Monopoly</h1>
<div class="flex flex-col justify-center grow">
    <Button label="Start" onClick={() => goto('/start')}/>
</div>
