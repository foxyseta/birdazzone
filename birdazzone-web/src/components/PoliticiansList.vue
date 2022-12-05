<script lang="ts" setup>
    import { ref, onBeforeMount } from 'vue'
    import ErrorWidget from '@/components/ErrorWidget.vue'
    import ApiRepository from '../api/api-repository'
    import type { Politician } from '../api/interfaces/politician'

    const error = ref<boolean> (false)
    const loading = ref<boolean> (true)
    const list = ref<Politician[]>([
        {
            "name": "Nome Cognome1",
            "score": 1000
        },
        {
            "name": "Nome Cognome2",
            "score": 900
        },
        {
            "name": "Nome Cognome3",
            "score": 800
        },
        {
            "name": "Nome Cognome4",
            "score": 700
        },
        {
            "name": "Nome Cognome5",
            "score": 600
        },
        {
            "name": "Nome Cognome6",
            "score": 500
        },
        {
            "name": "Nome Cognome7",
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
    <div class="flex flex-col flex-1 my-3" style="flex: 1 1 auto">
        <div class="flex flex-row w-full" style="flex: 1 1 auto">
            
            <div class="flex flex-col" style="flex: 2 1 auto"></div>
            
            <div class="flex flex-col justify-center align-center" style="flex:1 1 auto; width:13rem ">
                <div v-for="(item,index) in list" :key="index">
                    <div class="flex flex-row rounded-lg  bg-foreground p-4 m-2 items-center">
                        <div class="font-semibold text-white ml-2">
                            {{item.name}}
                        </div>
                        <div class="font-bold text-white ml-10 flex-1 text-center" style="font-size:180%">
                            {{item.score}}
                        </div>
                        <div class="flex justify-end" style="flex: 1 1 auto; height:2rem">
                            <img v-show="(index===0)" :src="'/icons/coccarda1.svg'" alt="medal1" class="mr-2"/>
                            <img v-show="(index===1)" :src="'/icons/coccarda2.svg'" alt="medal2" class="mr-2"/>
                            <img v-show="(index===2)" :src="'/icons/coccarda3.svg'" alt="medal3" class="mr-2"/>
                        </div>
                    </div>
                </div>
            </div>
            <div class="flex flex-col" style="flex: 2 1 auto"></div>
        </div>
    </div>
</template>