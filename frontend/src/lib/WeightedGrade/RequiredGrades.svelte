<script lang="ts">
    // TODO calculate required grades
</script>

<h2> Required Grades </h2>
<p> Grades possible from the remaining weight </p>
<div class="requiredGrades">
    <div class="gridFitTwoSmall">
        <box>
            <div class="gradePassLabel">P</div>
            <div class="gradePass">???</div>
        </box>
        
        <box>
            <div class="gradeCreditLabel">C</div>
            <div class="gradeCredit">???</div>
        </box>
    </div>
    <div class="gridFitTwoSmall">
        <box>
            <div class="gradeDistinctionLabel">D</div>
            <div class="gradeDistinction">???</div>
        </box>

        <box>
            <div class="gradeHighDistinctionLabel">HD</div>
            <div class="gradeHighDistinction">???</div>
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

    /* input style. Rounded border, centered text, gray background, white text, small amount of padding, not sunken, text size normal  */
    .gradePass, .gradeCredit, .gradeDistinction, .gradeHighDistinction {
        @include input.fakeInput;
    }

    .requiredGrades, .gridFitTwoSmall {
        display: grid;
        align-items: center;
    }

    .requiredGrades, .gridFitTwoSmall {
        grid-template-columns: repeat(auto-fit, minmax(100px, 1fr));
    }
</style>