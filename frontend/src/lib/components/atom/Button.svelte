<script lang="ts">
	let scaleClass = '';
	export let onClick = () => {};
</script>

<div>
	<button
		class={`px-[10rem] bg-primary rounded-full border-8
              border-dark primary-shadow text-[5rem] uppercase
              text-dark tracking-widest font-roboto transition ease-in-out duration-300 ${scaleClass}`}
		on:click={() => {
			scaleClass = 'scale-90';
			setTimeout(() => (scaleClass = ''), 300);
			setTimeout(() => onClick(), 600);
		}}
	>
		<slot />
	</button>
</div>
