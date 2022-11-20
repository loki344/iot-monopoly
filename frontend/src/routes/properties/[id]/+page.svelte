<script lang="ts">
    import { page } from '$app/stores';
	import PropertyBuyQuestionDialog from "$lib/components/molecule/PropertyBuyQuestionDialog.svelte";
	
</script>

<PropertyBuyQuestionDialog property={$page.data.property} buyerId={$page.data.buyerId}></PropertyBuyQuestionDialog>
