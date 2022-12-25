<script lang="ts" setup>
  const props = defineProps<{
    text: string;
    clickable: boolean;
  }>();
  const emit = defineEmits(['click']);

  const onClick = () => {
    emit('click');
  };
</script>
<template>
  <!-- Active -->
  <div
    v-if="clickable"
    @click="onClick"
    class="m-1 flex justify-center items-center border rounded-xl font-bold px-6 py-1 bg-foreground border-lgreen text-lgreen hover:bg-lgreen hover:text-white"
  >
    <p class="text-center">{{ props.text }}</p>
  </div>
  <!-- Inactive -->
  <div
    v-else
    class="m-1 flex justify-center items-center border rounded-xl font-bold px-6 py-1 bg-foreground border-lgray text-lgray"
  >
    <p class="text-center">{{ props.text }}</p>
  </div>
</template>
