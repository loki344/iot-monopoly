<script lang="ts">
    let clazz: String = ""
    export {clazz as class};
</script>

<p class={`w-1/5 self-center text-dark font-alpha bg-primary rounded-full text-[5rem] font-extrabold ${clazz}`}>
	<span>
		<slot></slot>
	</span>
</p>