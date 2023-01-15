<script lang="ts">
    import type { IWeightedClass } from "../types/weightedClass";

    // Type
    type SaveWeightedClass = () => void;

    // Props
    export let saveWeightedClass: SaveWeightedClass;

    export let classCode: string;
    export let totalWeight;
    export let finalWeightedGrade;

    // Calcuate remaining weight
    $: remainingWeight = 100 - totalWeight;

    // Class code changed
    function printClassCode() {
        console.log("Class Code: " + classCode);

        // TODO save
        saveWeightedClass();
    }
</script>

<div class="container">
    <div class="labelContainer">
        <h2 class="overview"> Overview </h2>
        <p class="overview"> A Simple Weighted Grade Calculator</p>
    </div>
    
    <div class="infoContainer">
        <box>
            <div class="classCode">Class Code</div>
            <input class="classCodeInput" type="text" placeholder="XXXX0000" bind:value={classCode} on:blur={printClassCode} />
        </box>

        <box>
            <div class="finalGrade">Final Weighted Grade</div>
            <div class="fakeInput"> {finalWeightedGrade}% </div>
        </box>

        <box>
            <div class="remainingWeight">Remaining Weight</div>
            <div class="fakeInput"> {remainingWeight} </div>
        </box>
    </div>
</div>

<style lang="scss">
    // Use sectionTitle style
    @use "../../style/sectionTitle";

    // Use input style
    @use "../../style/input";

    // Use common style
    @use "../../style/common";

    box {
        @include common.box;
    }

    input {
        @include input.inputStyle;
    }

    .fakeInput {
        @include input.fakeInput;
    }

    .container {
        display: grid;
        grid-auto-flow: row;
    }

    .infoContainer {
        display: grid;
        grid-auto-flow: column;
    }

    .overview {
        width: fit-content;
    }
</style>
