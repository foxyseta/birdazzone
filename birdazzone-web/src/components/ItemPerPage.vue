<script setup lang="ts">
    const props = defineProps<{itemPerPage:number}>()
    const options = [5,10,15,20]
    function log(value: any){
        console.log(value)
    }
</script>

<template>
    <div class="flex flex-row">

        <label for="cards" class="mr-2">Card per page:</label>
        <select name="cards" id="cards" class="flex justify-center" @change="(value)=>{$emit('changeItemPage',value.target?.value)}" type="dropdown" style="flex:0 0 0; border: solid 2px;">
            <option :selected="itemPerPage===5" value="5" >5</option>
            <option :selected="itemPerPage===10" value="10">10</option>
            <option :selected="itemPerPage===15" value="15">15</option>
            <option :selected="itemPerPage==20" value="20">20</option>"
        </select>    
    </div>
</template>