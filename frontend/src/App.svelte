<script>
    import WeightedGrade from './lib/WeightedGrade.svelte';

    // TODO enable this when the app is ready
    let startUp = false
</script>

<!-- Button that sets startUp to false -->
<!-- Only be rendered if startUp is true -->
{#if startUp}
    <button on:click="{() => startUp = false}">Start</button>
{/if}

<!-- WeightedGrade will only be rendered if startUp is false -->
{#if !startUp}
    <WeightedGrade />
{/if}

<style>
    /* Put button at bottom with a modern style */
    button {
        position: fixed;
        bottom: 0;
        right: 0;
        margin: 1rem;
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 0.25rem;
        background-color: #0070f3;
        color: white;
        font-size: 1rem;
        font-weight: 600;
        cursor: pointer;
    }
</style>
