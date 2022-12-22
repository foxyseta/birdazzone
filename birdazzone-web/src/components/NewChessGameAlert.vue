<script lang="ts" setup>
import { ref } from 'vue'
import ApiRepository  from '@/api/api-repository';

const emit = defineEmits(['done'])
const username = ref<string>("")
const error = ref<string>("")
const isError = ref<boolean>(false)

const checkUsername = async () => {
    const response = await ApiRepository.getHelloUser(username.value)

    isError.value = response.esit
    if (response.esit) {
        emit('done', response.data)    
    } else {
        error.value = "User not found :/ Try again!"
    }
}
</script>
<template>
    <div class="shadow bg-foreground rounded-2xl h-screen flex justify-start items-center m-10 p-10 flex-col">
        <h1 class="font-bold text-4xl text-white">Welcome to the Birdazzone chess game!</h1>
        <div class="m-10">
            <p class="m-3 text-white text-xl">Tell me what's your name</p>
            <input 
            v-model="username"
            class="form-control block px-3 py-1.5 text-base font-normal text-white bg-foreground bg-clip-padding border border-solid border-dgreen rounded-xl transition ease-in-out m-0 focus:border-lgreen focus:outline-none"
            @change="checkUsername"
            type="text"/>

        <div v-show="isError" class="p-2 text-center text-lred rounded-2xl border-lred">
            {{ error }}
         </div>
        </div>
    </div>
</template>