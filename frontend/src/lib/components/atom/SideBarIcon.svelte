<script lang="ts">
	export let onClick: any = undefined
		</script>

<span on:click={onClick} class="sidebar-icon">
	<slot/>
</span>
