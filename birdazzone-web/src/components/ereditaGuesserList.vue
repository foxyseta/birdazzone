<script setup lang="ts">
    import guesserListItem from '@/components/guesserListItem.vue'
    import {ref} from 'vue'
    let showAll = ref<boolean>(false)
    function prova(){
        console.log(showAll)
        showAll.value = !showAll.value
        console.log(showAll)
    }
    interface listItem {
        username: string
        fullName: string
        propic: string
        time: string
        body: string
    }
    const listItems = ref<listItem[]>([
        {
        username: "paoooool",
        fullName: "Paolo Ceroni",
        propic: "./public/defaultPropic.svg",
        time: "12:03",
        body:"prova lungaaaaaaaaaaaaaaaaaaaaaaaaaaa"
        },
        {
        username: "marco",
        fullName: "Fede Grì",
        propic: "./public/defaultPropic.svg",
        time: "12:05",
        body:"prova corta"
        }
    ])
</script>

<template>
    <div v-for="n in 2" class="">
        <guesserListItem class="flex flex-1"/>
    </div>
</template>