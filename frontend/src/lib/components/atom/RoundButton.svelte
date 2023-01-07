<script lang="ts">
	let scaleClass = '';
	export let onClick = () => {};
	let clazz: String = '';
	export { clazz as class };
	export let active = false;
	export let immediateOnClick = () => {};
	export let shadow = true
	export let big = false
	$: size = big ? 180 : 130
	$: textSize = big ? 7 : 5
</script>

<div
	style={`width:${size}px; height:${size}px;` }
	class={`${
		active ? 'bg-primary' : 'bg-gray'
	} ${shadow && active && 'primary-shadow'} ${shadow && !active && 'gray-shadow'} rounded-full border-8 flex flex-col items-center justify-center
    border-dark  text-[${textSize}rem] uppercase hover:cursor-pointer
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
