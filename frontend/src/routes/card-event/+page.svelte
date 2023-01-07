<script lang="ts">
    import { page } from '$app/stores';
	import CardDetails from '$lib/components/molecule/CardDetails.svelte';
	import { BASE_URL, extractData } from '$lib/http/backendClient';


	async function confirmCardRead() {
		await fetch(`${BASE_URL}/games/current/cards/current`, {
			method: 'PATCH',
			body: JSON.stringify({ accepted: true }),
			headers: {
				'Content-Type': 'application/json'
			}
		});
	}

</script>




<div class="h-full">
    <CardDetails text={$page.data.text} title={$page.data.title} onClick={confirmCardRead}></CardDetails>
</div>
