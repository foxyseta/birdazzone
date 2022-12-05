<script lang="ts" setup>
    import { ref, onBeforeMount } from 'vue'
    import ErrorWidget from '@/components/ErrorWidget.vue'
    import ApiRepository from '../api/api-repository'
    import type { Politician } from '../api/interfaces/politician'

    const error = ref<boolean> (false)
    const loading = ref<boolean> (true)
    const list = ref<Politician[]>([
        {
            "name": "Matteo Salvini1",
            "score": 1000
        },
        {
            "name": "Matteo Salvini2",
            "score": 900
        },
        {
            "name": "Matteo Salvini3",
            "score": 800
        },
        {
            "name": "Matteo SalvinI4",
            "score": 700
        },
        {
            "name": "Matteo Salvini5",
            "score": 600
        },
        {
            "name": "Matteo Salvini6",
            "score": 500
        },
        {
            "name": "Matteo Salvini7",
            "score": 400
        },
    ])
    
    /*const fetchPoliticiansList = async () => {
        loading.value = true
        const resp = await ApiRepository.getPoliticians()
        if (resp.esit) {
            list.value = resp.data!
            loading.value = false
        } else if (resp.statusCode === 404) {
            window.location.href = "/not-found"   
        } else {
            error.value = true
        }
    }*/

    //onBeforeMount(fetchPoliticiansList)
</script>

<template>
    <div class="flex flex-row w-full" style="flex: 1 1 auto">
        
        <div class="flex flex-col" style="flex: 2 1 auto"></div>
        
        <div class="flex flex-col justify-center align-center" style="flex:1 1 auto">
            <div v-for="(item,index) in list" :key="index">
                <div class="flex flex-row rounded-lg  bg-foreground p-4 m-2 items-center">
                    <div class="font-semibold text-white ml-2">
                        {{item.name}}
                    </div>
                    <div class="font-bold text-white ml-10" style="font-size:180%">
                        {{item.score}}
                    </div>
                </div>
            </div>
        </div>
        <div class="flex flex-col" style="flex: 2 1 auto"></div>
    </div>
</template>