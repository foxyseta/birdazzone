<script setup lang=ts>
import { ref } from 'vue'

const error = ref<boolean>(false)

</script>

<template>
  <div v-if="error">
    <ErrorWidget />
  </div>

  <div class="bg-foreground font-semibold text-lg rounded-lg m-4 p-4 place-self-center">

    <!--<div v-if="loading">
      <semipolar-spinner :animation-duration="2000" :size="35" color="#1eb980" />
    </div>-->

  </div>
</template>
