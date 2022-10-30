<script lang="ts">
	export let text = '';
	let clazz = '';
	export { clazz as class };
	export let medium = false;
</script>

<p
	class={`uppercase text-primary font-alfa  ${
		medium ? 'text-[4rem] tracking-widest' : 'text-[5rem] tracking-[1.5rem]'
	} ${clazz}`}
>
	<slot />
</p>
