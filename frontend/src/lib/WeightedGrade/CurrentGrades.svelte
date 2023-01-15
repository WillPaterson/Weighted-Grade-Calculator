<script lang="ts">
    // Components
    import AddGrade from "./AddGrade.svelte";
    import Fa from 'svelte-fa'
    import { faXmark } from '@fortawesome/free-solid-svg-icons'

    // Type
    import type { IGrade } from "../types/grade";
    type SaveWeightedClass = () => void;

    // Props
    export let grades: IGrade[]
    export let totalWeight: number
    export let saveWeightedClass: SaveWeightedClass;

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
        let percentage = calculateGradePercentage(grade, totalPossible);

        console.log("Percentage: " + percentage + "%");
        

        let newGrade: IGrade = {
            grade: grade,
            totalPossible: totalPossible,
            percentage: percentage,
            weight: weight
        }

        grades = [...grades, newGrade];

        saveWeightedClass();
    }

    function removeGrade(index: number) {
        grades = grades.filter((_, i) => i !== index);

        saveWeightedClass();
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
                <th>Remove?</th>
            </tr>
            {#each grades as grade, index}
                <tr>
                    <td>{index + 1}</td>
                    <td>{grade.grade}</td>
                    <td>{grade.totalPossible}</td>
                    <td>{grade.percentage}%</td>
                    <td>{grade.weight}</td>
                    <!-- Remove button with svg cross -->
                    
                    <td>
                        <div class="centerButton">
                            <button class="close" on:click={() => removeGrade(index)}>
                                <div style="color: white">
                                    <Fa icon={faXmark} />
                                </div>
                            </button>
                        </div>
                    </td>
                </tr>
            {/each}
        {/if}
    </table>
</div>

<AddGrade {addGrade} {calculateTotalWeight} />

<style lang="scss">
    // Use sectionTitle style
    @use "../../style/sectionTitle";

    // Use common style
    @use "../../style/common";

    // Use button style
    @use "../../style/button";

    .currentGrades  {
        @include common.flexCenter;
    }

    .close {
        @include button.quickButtonStyle($use-slim: true);
    }

    .centerButton {
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
