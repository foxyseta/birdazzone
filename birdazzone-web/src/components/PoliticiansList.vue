<script lang="ts" setup>
    import { onBeforeMount } from "@vue/runtime-core"
    import { ref, onBeforeMount } from 'vue'
    import ErrorWidget from '@/components/ErrorWidget.vue'
    import ApiRepository from '../api/api-repository'
    import type { Politician } from '../api/interfaces/politician'

    const error = ref<boolean> (false)
    const loading = ref<boolean> (true)
    const list = ref<Politician[]>([])

    const fetchPoliticiansList = async () => {
        loading.value = true
        const resp = await ApiRepository.getPoliticians()
        if (resp.esit) {
            //list.value = resp.data!.entries
            loading.value = false
        } else if (resp.statusCode === 404) {
        window.location.href = "/not-found"   
        } else {
        error.value = true
        }
    }

    onBeforeMount(fetchPoliticiansList)
</script>

<template>
    
</template>