<script lang="ts">
    // TODO calculate required grades
    let requiredGradeLimits = {
        Pass: 50,
        Credit: 65,
        Distinction: 75,
        HighDistinction: 85
    }

    let requiredGrades = {
        Pass: "???",
        Credit: "???",
        Distinction: "???",
        HighDistinction: "???"
    }

    // Props
    export let totalWeight;
    export let finalWeightedGrade;

    // Calcuate remaining weight
    $: remainingWeight = 100 - totalWeight;

    // Calculate required grades for each grade
    $: {
        for (let grade in requiredGrades) {
            let requiredGrade = (requiredGradeLimits[grade] * 100 - (finalWeightedGrade) * totalWeight) / (remainingWeight)
            // Round to 2 decimal places unless round number
            requiredGrades[grade] = Math.round(requiredGrade * 100) / 100
        }
    }
</script>

<h2> Required Grades </h2>
<p> Minimum grade needed for desired scaled grade </p>  
<div class="requiredGrades">
    <div class="gridFitTwoSmall">
        <box class="boxPass">
            <div class="gradePassLabel">P</div>
            <div class="gradePass"> {requiredGrades.Pass} </div>
        </box>
        
        <box class="boxCredit">
            <div class="gradeCreditLabel">C</div>
            <div class="gradeCredit"> {requiredGrades.Credit} </div>
        </box>
    </div>
    <div class="gridFitTwoSmall">
        <box class="boxDistinction">
            <div class="gradeDistinctionLabel">D</div>
            <div class="gradeDistinction"> {requiredGrades.Distinction} </div>
        </box>

        <box class="boxHighDistinction">
            <div class="gradeHighDistinctionLabel">HD</div>
            <div class="gradeHighDistinction"> {requiredGrades.HighDistinction} </div>
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