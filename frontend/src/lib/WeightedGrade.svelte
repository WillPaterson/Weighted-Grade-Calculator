<script lang="ts">
    // Components
    import CurrentGrades from "./WeightedGrade/CurrentGrades.svelte";
    import Info from "./WeightedGrade/Info.svelte";
    import RequiredGrades from "./WeightedGrade/RequiredGrades.svelte";
</script>

<main>
    <div class = "largeGrid">
        <section class="info">
            <Info />
        </section>
    
        <section class="RequiredGrades">
            <RequiredGrades />
        </section>
    </div>

    <div class="largeGrid">
        <section class="CurrentGrades">
            <CurrentGrades />
        </section>
    </div>
</main>

<style>
    main {
        color: #bdbab4;
        display: grid;
        grid-auto-flow: row;
        grid-template-columns: repeat(auto-fit, minmax(570px, 1fr));
        grid-gap: 0.5rem;
        font-family: 'Rubik', sans-serif;
    }

    .largeGrid {
        display: grid;
        grid-gap: 0.5rem;
    }

    section {
        margin: 0.1rem;
        margin-bottom: 1rem;
        background-color: #383a3c;
        border-radius: 5px;
        padding: 1rem;
    }
</style>