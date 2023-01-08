<script lang="ts">
	import Tag from "./Tag.svelte";


	let clazz: String = '';
	export { clazz as class };
	export let price: number;
</script>


<Tag class={clazz}>{price.toLocaleString()} $</Tag>
