<script lang="ts">
    // Components
    import CurrentGrade from "./CurrentGrade.svelte";
    import Modal from "./Modal.svelte";
    
    // Type
    import type { IGrade } from "../../types/grade";
    

    let grades: IGrade[] = [
        {id: 0, mark: 100, totalPossible: 100, weight: 10},
        {id: 1, mark: 100, totalPossible: 100, weight: 10},
        {id: 2, mark: 100, totalPossible: 100, weight: 10},
        {id: 3, mark: 100, totalPossible: 100, weight: 10},
        {id: 4, mark: 100, totalPossible: 100, weight: 10},
    ];

    $: amountOfGrades = grades.length;

	let showModal = false;


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
                <th>Mark</th>
                <th>Total Possible</th>
                <th>Mark Percentage</th>
                <th>Weight</th>
            </tr>
            {#each grades as grade}
                <CurrentGrade {grade} />
            {/each}
        {/if}
    </table>
</div>

<div class="addGrade">
    <button class="addGradeButton" on:click="{() => showModal = true}">Add Grade</button>

    {#if showModal}
        <Modal on:close="{() => showModal = false}">
            <h2 slot="header">
                Add Grade
            </h2>

            <p>Grade</p>
            <input type="number" placeholder="Grade" />
            <p>Total Possible</p>
            <input type="number" placeholder="Total Possible" />
        </Modal>
    {/if}
</div>

<style lang="scss">
    // Use sectionTitle style
    @use "../../style/sectionTitle";

    // Use common style
    @use "../../style/common";

    // Use button style
    @use "../../style/button";

    .currentGrades, .addGrade  {
        @include common.flexCenter;
    }

    .addGradeButton {
        @include button.buttonStyle;
    }

    .addGradeButton:hover {
        @include button.buttonHover;
    }
    .addGradeButton:active {
        @include button.buttonActive;
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
