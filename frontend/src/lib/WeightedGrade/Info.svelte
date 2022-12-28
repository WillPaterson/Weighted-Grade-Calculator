<div class="container">
    <div class="labelContainer">
        <h2 class="overview"> Overview </h2>
        <p class="overview"> A Simple Weighted Grade Calculator</p>
    </div>
    
    <div class="infoContainer">
        <box>
            <div class="classCode">Class Code</div>
            <input class="classCodeInput" type="text" placeholder="XXXX0000">
        </box>
        
        <box>
            <div class="remainingWeight">Remaining Weight</div>
            <div class="fakeInput">100%</div>
        </box>
    </div>
</div>

<style lang="scss">
    // Use sectionTitle style
    @use "../../style/sectionTitle";

    // Use box style
    @use "../../style/box";

    // Use input style
    @use "../../style/input";

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
