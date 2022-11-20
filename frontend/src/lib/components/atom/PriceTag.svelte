<script lang="ts">
	let clazz: String = '';
	export { clazz as class };
	export let price: number;
</script>

<p class={`w-1/5 self-center text-dark font-alpha bg-primary rounded-full text-[5rem] font-extrabold ${clazz}`}>
	<span>
		{price.toLocaleString()} $
	</span>
</p>
