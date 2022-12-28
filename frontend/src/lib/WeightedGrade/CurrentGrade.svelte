<script lang="ts">
    export let grade

    let markPercentage = 0

    // Calculate the percentage of the mark to 2 decimal places
    $: markPercentage = Math.round((grade.grade / grade.totalPossible) * 10000) / 100
</script>
<tr>
    <td>{grade.id + 1}</td>
    <td>{grade.grade}</td>
    <td>{grade.totalPossible}</td>
    <td>{markPercentage}%</td>
    <td>{grade.weight}</td>    
</tr>


<style>
    td {
        text-align: left;
        padding: 8px;
    }
</style>