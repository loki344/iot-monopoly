<script>
	import { goto } from '$app/navigation';

	import Button from '$lib/components/atom/Button.svelte';
	import Title from '$lib/components/atom/Title.svelte';
</script>

<div class="flex flex-col items-center justify-center h-screen">
	<Title class="mb-[15rem]">MONOPOLY</Title>
	<Button onClick={() => goto('/start')}>Start</Button>
</div>
