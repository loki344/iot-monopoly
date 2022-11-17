<script lang="ts">
	let scaleClass = '';
	export let onClick = () => {};
	let clazz: String = '';
	export { clazz as class };
	export let active = false;
	export let immediateOnClick = () => {};
	export let shadow = true
</script>

<div
	style="width:130px; height:130px;"
	class={`${
		active ? 'bg-primary' : 'bg-gray'
	} ${shadow && active && 'primary-shadow'} ${shadow && !active && 'gray-shadow'} rounded-full border-8 flex flex-col items-center justify-center
    border-dark  text-[5rem] uppercase hover:cursor-pointer
    text-dark tracking-widest font-roboto transition ease-in-out duration-300 ${scaleClass} ${clazz}`}
	on:click={() => {
		scaleClass = 'scale-90';
		immediateOnClick();
		setTimeout(() => onClick(), 600);
		setTimeout(() => (scaleClass = ''), 300);
	}}
>
	<slot />
</div>
