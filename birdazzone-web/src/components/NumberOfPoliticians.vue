<script lang="ts" setup>
const propss = defineProps<{ num: number }>();
console.log(propss.num);
</script>

<template>
  <div class="flex w-full bg-foreground items-center rounded-lg p-3" style="flex: 1 1 auto">
    <div class="text-white font-bold">Total number of ranked politicians:</div>
    <div class="text-lgreen font-bold mx-4" style="font-size: 200%">{{ propss.num }}</div>
  </div>
</template>
