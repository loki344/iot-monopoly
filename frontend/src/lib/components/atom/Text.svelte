<script lang="ts">
	let clazz: String = '';
	export { clazz as class };
</script>

<p class={`uppercase text-dark font-alfa text-[3rem] ${clazz}`}>
	<slot />
</p>
