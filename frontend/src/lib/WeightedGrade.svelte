<script lang="ts">
    // Version
    const version = __APP_VERSION__;

    // Components
    import CurrentGrades from "./WeightedGrade/CurrentGrades.svelte";
    import Info from "./WeightedGrade/Info.svelte";
    import RequiredGrades from "./WeightedGrade/RequiredGrades.svelte";
    
    // Type
    import type { IWeightedClass } from "./types/weightedClass";
    import type { IGrade } from "./types/grade";
    
    // Native Methods
    import { main } from "../../wailsjs/go/models";
    import { SaveWeightedClass } from "../../wailsjs/go/main/App";
    import { LoadJSONFile } from "../../wailsjs/go/main/App";
    

    let weightedClass: IWeightedClass = {
        classCode: "",
        gradeList: Array<IGrade>()
    }

    LoadJSONFile().then((data) => {
        console.log("Data: " + data);
        weightedClass = JSON.parse(data);

    });

    let totalWeight = 0;
    let finalWeightedGrade = 0;

    $: weightedClass.gradeList, totalWeight, finalWeightedGrade = calculateFinalWeightedGrade();
    function calculateFinalWeightedGrade() {
        let totalWeightedGrade = 0;

        if (totalWeight === 0) {
            return 0;
        }
    
        weightedClass.gradeList.forEach(grade => {
            totalWeightedGrade += grade.percentage * grade.weight;
        });

        return Math.round(totalWeightedGrade / totalWeight);
    }

    $: weightedClass.gradeList;
    function saveWeightedClass() {
        // Create output weighted class
        let outputWeightedClass = main.WeightedClass.createFrom(weightedClass);

        // Print pretty json
        console.log(JSON.stringify(weightedClass, null, 4));

        SaveWeightedClass(outputWeightedClass);
    }

</script>

<main class="scrollable">
    <div class = "largeGrid">
        <section class="info">
            <Info bind:classCode={weightedClass.classCode} {totalWeight} {finalWeightedGrade} {saveWeightedClass}/>
        </section>
    
        <section class="RequiredGrades">
            <RequiredGrades {totalWeight} {finalWeightedGrade}/>
        </section>
    </div>

    <div class="largeGrid">
        <section class="CurrentGrades">
            <CurrentGrades bind:grades={weightedClass.gradeList} bind:totalWeight/>
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