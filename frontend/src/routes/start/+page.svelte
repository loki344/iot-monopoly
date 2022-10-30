<script>
	import Button from '$lib/components/atom/Button.svelte';

	import Title from '$lib/components/atom/Title.svelte';
	import Tags from '$lib/components/molecule/Tags.svelte';
</script>

<div class="flex flex-col items-center justify-around h-screen">
	<Title medium class="mb-[15rem]">HOW MANY ARE PLAYING?</Title>
	<Tags items={['1', '2', '3', '4']} />
	<Button onClick={() => {}}>Start</Button>
</div>
