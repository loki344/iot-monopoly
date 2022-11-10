<script lang="ts">
	export let text = '';
	let clazz = '';
	export { clazz as class };
	export let type = '';

	function getClassForType(type: string) {
		switch (type) {
			case 'small':
				return 'text-[3rem] tracking-widest';
			case 'medium':
				return 'text-[4rem] tracking-widest';
			default:
				return 'text-[5rem] tracking-[1.5rem]';
		}
	}
</script>

<p class={`uppercase text-primary my-8 font-alfa  ${getClassForType(type)} ${clazz}`}>
	<slot />
</p>
