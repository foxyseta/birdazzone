<script setup lang="ts">
    import guesserListItem from '@/components/guesserListItem.vue'
    import {ref} from 'vue'
    import { unref } from 'vue'
    import { useFetch } from '@/use-fetch'


    const { data, error } = useFetch('tvgames/0')
    
    console.log(data.value , error.value)
    
    let showAll = ref<boolean>(false)
    function prova(){
        console.log(showAll)
        showAll.value = !showAll.value
        console.log(showAll)
    }
</script>

<template>
    <div v-for="n in 8" class="">
        <guesserListItem class="flex flex-1"/>
    </div>
</template>
