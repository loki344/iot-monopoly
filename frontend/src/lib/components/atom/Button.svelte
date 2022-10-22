<script lang="ts">

    export let label: string = ''
    export let onClick: any
    export let clazz: string = ''
    export let active: boolean = false
</script>

<button on:click={onClick} class={`${clazz} text-md rounded-lg border p-3 px-8 ${active && "bg-white"} active:bg-white`}>{label}<slot></slot></button>