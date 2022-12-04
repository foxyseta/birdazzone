<script lang="ts" setup>
    import { onBeforeMount } from "@vue/runtime-core"
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

    console.log(list.value)
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
    <div class="flex flex-col" v-for="(item,index) in list" :key="index">
        <div class="rounded-lg  bg-foreground p-4">
            {{item.name}}
        </div>
    </div>
</template>