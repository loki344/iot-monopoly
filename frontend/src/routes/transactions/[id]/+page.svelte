<script lang="ts">
	import ConfirmTransactionDialog from "$lib/components/molecule/ConfirmTransactionDialog.svelte";
    import { page } from '$app/stores';
	import Title from "$lib/components/atom/Title.svelte";
</script>

<div class="h-full flex flex-col justify-center text-center">
<Title type={'medium'}>Transaction</Title>
<ConfirmTransactionDialog transaction={{senderId: $page.data.senderId, receiverId: $page.data.receiverId, amount: $page.data.amount}}></ConfirmTransactionDialog>
</div>
