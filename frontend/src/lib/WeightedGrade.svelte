<script lang="ts">
    // Components
    import CurrentGrades from "./WeightedGrade/CurrentGrades.svelte";
    import Info from "./WeightedGrade/Info.svelte";
    import RequiredGrades from "./WeightedGrade/RequiredGrades.svelte";

    // Type
    import type { IGrade } from "../types/grade";

    let grades: IGrade[] = [];
    let totalWeight: number = 0;
    let finalWeightedGrade: number = 0;

    $: grades, totalWeight, finalWeightedGrade = calculateFinalWeightedGrade();
    function calculateFinalWeightedGrade() {
        let totalWeightedGrade = 0;
        
        if (totalWeight === 0) {
            return 0;
        }

        grades.forEach(grade => {
            totalWeightedGrade += grade.percentage * grade.weight;
        });

        return Math.round(totalWeightedGrade / totalWeight);
    }

</script>

<main class="scrollable">
    <div class = "largeGrid">
        <section class="info">
            <Info {totalWeight} {finalWeightedGrade}/>
        </section>
    
        <section class="RequiredGrades">
            <RequiredGrades />
        </section>
    </div>

    <div class="largeGrid">
        <section class="CurrentGrades">
            <CurrentGrades bind:grades bind:totalWeight/>
        </section>
    </div>
</main>

<style>
    main {
        color: #bdbab4;
        display: grid;
        grid-auto-flow: row;
        grid-template-columns: repeat(auto-fit, minmax(600px, 1fr));
        grid-gap: 0.5rem;
        font-family: 'Rubik', sans-serif;

        overflow: hidden;
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