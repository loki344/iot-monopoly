<script lang="ts">
	let scaleClass = '';
	export let onClick = () => {};
	let clazz: String = '';
	export { clazz as class };
</script>

<div
	style="width:130px; height:130px;"
	class={`bg-primary rounded-full border-8 flex flex-col items-center justify-center
    border-dark primary-shadow text-[5rem] uppercase
    text-dark tracking-widest font-roboto transition ease-in-out duration-300 ${scaleClass} ${clazz}`}
	on:click={() => {
		scaleClass = 'scale-90';
		setTimeout(() => (scaleClass = ''), 300);
		setTimeout(() => onClick(), 600);
	}}
>
	<slot />
</div>
