<script lang="ts">
	import Icon from 'svelte-icon/Icon.svelte';
	import homeOutline from '$lib/components/atom/icon/eva-icon/home-outline.svg?raw';
	export let clazz = '';
</script>

<Icon class={clazz} data={homeOutline} />
