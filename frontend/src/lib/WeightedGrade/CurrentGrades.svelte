<script lang="ts">
    // Components
    import CurrentGrade from "./CurrentGrade.svelte";
    
    // Type
    import type { IGrade } from "../../types/grade";
    import AddGrade from "./AddGrade.svelte";
    

    let grades: IGrade[] = [];

    $: amountOfGrades = grades.length;

    function addGrade(grade: number, totalPossible: number, weight: number) {
        let newGrade: IGrade = {
            id: grades.length,
            grade: grade,
            totalPossible: totalPossible,
            weight: weight
        }
        
        grades = [...grades, newGrade];
    }
</script>

<h2>Current Grades</h2>
<p class="pInfo">Current grades for the completed weight</p>
<div class="currentGrades">
    <table>
        {#if amountOfGrades <= 0}
            <tr>
                <th class="centerHeading">No grades added</th>
            </tr>
        {:else}
            <tr>
                <th></th>
                <th>Grade Achieved</th>
                <th>Total Possible</th>
                <th>Grade Percentage</th>
                <th>Weight</th>
            </tr>
            {#each grades as grade}
                <CurrentGrade {grade} />
            {/each}
        {/if}
    </table>
</div>

<AddGrade {addGrade}/>

<style lang="scss">
    // Use sectionTitle style
    @use "../../style/sectionTitle";

    // Use common style
    @use "../../style/common";

    .currentGrades  {
        @include common.flexCenter;
    }

    /* Table style */
    table {
        background-color: #2c2d2d;
        border-radius: 5px;
        padding: 1rem;
        margin: 0.5rem;

        border-collapse: collapse;
        width: 100%;
    }

    .centerHeading {
        text-align: center;
    }

    th {
        text-align: left;
        padding: 8px;
    }
</style>
