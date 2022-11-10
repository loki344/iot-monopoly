<script lang="ts">
	let scaleClass = '';
	export let onClick = () => {};
	export let disabled = false;
	let clazz = '';
	export { clazz as class };
	export let type = '';
	const getTypeClass = (type: String) => {
		switch (type) {
			case 'NEGATIVE':
				return 'bg-red red-shadow';
			case 'POSITIVE':
				return 'bg-green green-shadow';
			default:
				return 'bg-primary primary-shadow';
		}
	};
	let typeClass = getTypeClass(type);
</script>

<div>
	<button
		class={`px-[10rem] ${disabled ? 'bg-gray gray-shadow' : typeClass} rounded-full border-8
              border-dark  text-[5rem] uppercase ${clazz}
              text-dark tracking-widest font-roboto transition ease-in-out duration-300 ${scaleClass}`}
		on:click={() => {
			if (!disabled) {
				scaleClass = 'scale-90';
				setTimeout(() => (scaleClass = ''), 300);
				setTimeout(() => onClick(), 600);
			}
		}}
		{disabled}
	>
		<slot />
	</button>
</div>
