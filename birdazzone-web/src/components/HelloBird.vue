<script lang="ts" setup >
import {ref, onBeforeMount} from 'vue'
import ApiRepository from '../api/api-repository'

const message = ref<string>("")

onBeforeMount(async () => {
  const response = await ApiRepository.getHelloBird()
  if (response.esit)
    message.value = response.data!.message
  else
    message.value = response.error
})

</script>
<template>
  <h1> {{message}}</h1>
</template>
