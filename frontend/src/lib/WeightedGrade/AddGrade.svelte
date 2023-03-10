<script lang="ts">
    // Components
    import Modal from "./Modal.svelte";

    // Type
    type AddGrade = (grade: number, totalPossible: number, weight: number) => void;

    // Props
    export let addGrade: AddGrade;
    export let calculateTotalWeight: () => number;

    // State
    let grade = 0;
    let totalPossible = 0;
    let weight = 0;

    // Used to show the modal
    let showModal = false;

    // Used to store error messages
    let errorMessages = []

    function resetErrors() {
        errorMessages = [];
    }

    function handleAddGrade() {
        resetErrors();
        
        let totalWeight = calculateTotalWeight();

        console.log("Total Weight: " + totalWeight);
        

        if (totalPossible <= 0) {
            errorMessages.push("Total Possible must be greater than 0");
        }
        if (weight <= 0) {
            errorMessages.push("Weight must be greater than 0");
        }
        if (grade < 0) {
            errorMessages.push("Grade Achieved must be greater than or equal to 0");
        }
        if (grade > totalPossible) {
            errorMessages.push("Grade Achieved must be less than or equal to Total Possible");
        }
        if (totalWeight + weight > 100) {
            errorMessages.push("Total weight cannot be greater than 100");
        }
        if (errorMessages.length <= 0) {
            console.log("No errors");
            addGrade(grade, totalPossible, weight);
        }
    }
</script>

<div class="addGrade">
    <button class="addGradeButton" on:click="{() => showModal = true}">Add Grade</button>

    {#if showModal}
        <Modal on:close="{() => showModal = false}">
            <h2 slot="header">
                Add Grade
            </h2>

            <div class="modalBody" slot="body">
                <!-- Add Grade From -->
                <form on:submit|preventDefault={handleAddGrade}>
                    <div class="row">
                        <box>
                            <label for="grade">Grade Achieved</label>
                            <input type="number" placeholder="Grade Achieved" bind:value={grade} />
                        </box>
                    </div>
                    
                    <div class="row">
                        <box>
                            <label for="totalPossible">Total Possible</label>
                            <input type="number" placeholder="Total Possible" bind:value={totalPossible} />
                        </box>
                    </div>
                    
                    <div class="row">
                        <box>
                            <label for="weight">Weight</label>
                            <input type="number" placeholder="Weight" min="1" bind:value={weight} />
                        </box>
                    </div>
                </form>

                
                {#each errorMessages as message}
                    <div class="error">
                        {message}
                    </div>
                {/each}
            
            </div>
            
            
            <div class="modalFooter" slot="footer">
                <button class="submitGradeButton" on:click={handleAddGrade}>Add Grade</button>
            </div>
        </Modal>
    {/if}
</div>

<style lang="scss">
    // Use common style
    @use "../../style/common";

    // Use button style
    @use "../../style/button";

    // Use input style
    @use "../../style/input";

    box {
        @include common.box;
    }

    .error {
        @include common.box(#790000);
    }

    .addGrade, input, label, .modalFooter {
        @include common.flexCenter;
    }

    form div {float:left;width:33.3%;}

    .addGradeButton {
        @include button.quickButtonStyle;
    }

    .submitGradeButton {
        @include button.quickButtonStyle( #006800, #005200, $width:30%);
    }

    input {
        @include input.inputStyle;
    }
    
    input::-webkit-outer-spin-button,
    input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
    }
</style>