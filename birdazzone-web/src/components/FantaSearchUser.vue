<script lang="ts" setup>
import { ref } from 'vue';

const username = ref("")

</script>
<template>
    <div class="p-5 w-full flex justify-center">
        <input
        type="text"
        v-model="username"
        class="
            form-control
            block
            px-3
            py-1.5
            text-base
            font-normal
            text-white
            bg-foreground 
            bg-clip-padding
            border border-solid border-dgreen 
            rounded-xl
            transition
            ease-in-out
            m-0
            focus:border-lgreen focus:outline-none
        "
        @keypress="$emit('changed', username)"
        placeholder="username"
        />
    </div>
</template>