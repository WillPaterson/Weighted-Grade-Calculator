<script lang="ts">
    // TODO calculate required grades
</script>

<h2> Required Grades </h2>
<p> Minimum grade needed for desired scaled grade </p>  
<div class="requiredGrades">
    <div class="gridFitTwoSmall">
        <box class="boxPass">
            <div class="gradePassLabel">P</div>
            <div class="gradePass">???</div>
        </box>
        
        <box class="boxCredit">
            <div class="gradeCreditLabel">C</div>
            <div class="gradeCredit">???</div>
        </box>
    </div>
    <div class="gridFitTwoSmall">
        <box class="boxDistinction">
            <div class="gradeDistinctionLabel">D</div>
            <div class="gradeDistinction">???</div>
        </box>

        <box class="boxHighDistinction">
            <div class="gradeHighDistinctionLabel">HD</div>
            <div class="gradeHighDistinction">???</div>
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

    $colors: #790000, #803c00, #006800, #00497e;
    $grades: Pass, Credit, Distinction, HighDistinction;
    
    @each $color, $grade in zip($colors, $grades) {
        .grade#{$grade} {
            @include input.fakeInput($color)
        }
    }

    .requiredGrades, .gridFitTwoSmall {
        display: grid;
        align-items: center;
    }

    .requiredGrades, .gridFitTwoSmall {
        grid-template-columns: repeat(auto-fit, minmax(100px, 1fr));
    }
</style>