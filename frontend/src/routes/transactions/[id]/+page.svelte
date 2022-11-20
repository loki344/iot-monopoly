<script lang="ts">
	import ConfirmTransactionDialog from "$lib/components/molecule/ConfirmTransactionDialog.svelte";
    import { page } from '$app/stores';
</script>

<ConfirmTransactionDialog transaction={{senderId: $page.data.senderId, recipientId: $page.data.recipientId, amount: $page.data.amount}}></ConfirmTransactionDialog>
