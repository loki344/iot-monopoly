<script lang="ts">
	let clazz: String = '';
	export { clazz as class };
	export let price: number;
</script>

<p class={`text-dark font-alpha  text-[5rem] font-extrabold ${clazz}`}>
	{price.toLocaleString()} $
</p>
