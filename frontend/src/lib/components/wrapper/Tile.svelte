<script lang="ts">
	let clazz: String = '';
	export { clazz as class };
	export let active = true;
</script>

<div
	class={`flex flex-col ${
		active ? 'bg-primary border-8 border-green' : 'bg-gray border-8 border-gray'
	} py-10 px-5 rounded-lg   ${clazz}`}
>
	<slot />
</div>
