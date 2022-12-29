<script lang="ts">
    // Components
    import AddGrade from "./AddGrade.svelte";

    // Type
    import type { IGrade } from "../../types/grade";

    // Data
    export let grades
    export let totalWeight

    // Reactive
    $: amountOfGrades = grades.length;
    $: grades, totalWeight = calculateTotalWeight();
    function calculateTotalWeight() {
        let totalWeight = 0;

        grades.forEach(grade => {
            totalWeight += grade.weight;
        });

        return totalWeight;
    }

    function calculateGradePercentage(grade: number, totalPossible: number) {
        return Math.round((grade / totalPossible) * 10000) / 100;
    }

    function addGrade(grade: number, totalPossible: number, weight: number) {
        let newGrade: IGrade = {
            id: grades.length,
            grade: grade,
            totalPossible: totalPossible,
            percentage: calculateGradePercentage(grade, totalPossible),
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
                <tr>
                    <td>{grade.id + 1}</td>
                    <td>{grade.grade}</td>
                    <td>{grade.totalPossible}</td>
                    <td>{grade.percentage}%</td>
                    <td>{grade.weight}</td>    
                </tr>
            {/each}
        {/if}
    </table>
</div>

<AddGrade {addGrade} {calculateTotalWeight}/>

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

    th, td {
        text-align: left;
        padding: 8px;
    }
</style>
