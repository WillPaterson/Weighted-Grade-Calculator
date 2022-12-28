<script lang="ts">
    export let grade
</script>
<tr>
    <td>{grade.id}</td>
    <td>{grade.mark}</td>
    <td>{grade.totalPossible}</td>
    <td>{grade.markPercentage}%</td>
    <td>{grade.weight}%</td>    
</tr>


<style>
    td {
        text-align: left;
        padding: 8px;
    }
</style>