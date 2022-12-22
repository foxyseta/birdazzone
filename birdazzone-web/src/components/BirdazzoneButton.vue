<script lang="ts" setup>
const props = defineProps<{ text: string; active: boolean }>();
</script>
<template>
  <!-- Active -->
  <div
    :class="`m-1 flex justify-center items-center border rounded-xl font-bold px-6 py-1 bg-foreground border-${
      props.active ? 'lgreen' : 'lgray'
    } text-${props.active ? 'lgreen' : 'lgray'} hover:bg-lgreen hover:text-white`"
  >
    <p class="text-center">{{ props.text }}</p>
  </div>
</template>
