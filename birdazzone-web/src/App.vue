<script setup lang="ts">
import { getHelloBird } from './api/api-repository';

const test = async ()=>{await getHelloBird()}
</script>

<template>
    <button @click="test">GET</button>
</template>
