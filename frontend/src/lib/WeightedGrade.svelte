<script lang="ts">
    // Version
    const version = __APP_VERSION__;

    // Components
    import CurrentGrades from "./WeightedGrade/CurrentGrades.svelte";
    import Info from "./WeightedGrade/Info.svelte";
    import RequiredGrades from "./WeightedGrade/RequiredGrades.svelte";

    // Importing from native methods
    import { LogPrint } from "../../wailsjs/runtime";
    LogPrint("Hello from Svelte!");
    
    // Type
    import type { IWeightedClass } from "./types/weightedClass";

    let weightedClass: IWeightedClass = {
        classCode: "CSC 101",
        gradeList: [
            {
                id: 0,
                grade: 50,
                totalPossible: 100,
                percentage: 50,
                weight: 20
            },
            {
                id: 1,
                grade: 100,
                totalPossible: 100,
                percentage: 100,
                weight: 10
            }
        ]
    }

    let totalWeight = 0;
    let finalWeightedGrade = 0;

    // Reactive
    $: classCode = weightedClass.classCode;
    $: grades = weightedClass.gradeList;

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
            <Info {classCode} {totalWeight} {finalWeightedGrade}/>
        </section>
    
        <section class="RequiredGrades">
            <RequiredGrades {totalWeight} {finalWeightedGrade}/>
        </section>
    </div>

    <div class="largeGrid">
        <section class="CurrentGrades">
            <CurrentGrades bind:grades bind:totalWeight/>
        </section>
    </div>
</main>

<!-- Pin version number from package.json to bottom -->
<section-small class="flexToPageEnd">
    <p>Version: {version}</p>
</section-small>

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

    .flexToPageEnd {
        display: flex;
        justify-content: flex-end;

        font-family: 'Rubik', sans-serif;
        color: #bdbab4;

        margin-top: 0.5rem;
        background-color: #383a3c;
        border-radius: 5px;
        padding-right: 1rem;
    }
</style>