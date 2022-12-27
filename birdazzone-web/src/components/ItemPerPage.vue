<script setup lang="ts">
  const props = defineProps<{itemPerPage: number}>();
  const options = [5, 10, 15, 20];
</script>

<template>
  <div
    class="flex flex-row p-2 rounded-lg text-lgray justify-center"
    style="width: 13rem; flex: 0 1 auto; border: 2px solid #1eb980"
  >
    <label
      for="cards"
      class="mr-2"
      >Twitters per page:</label
    >
    <select
      name="cards"
      id="cards"
      class="flex justify-center bg-background text-white"
      @change="
        value => {
          //@ts-ignore
          $emit('changeItemPage', value.target?.value);
        }
      "
      type="dropdown"
      style="flex: 0 0 0; outline: 0px"
    >
      <option
        :selected="itemPerPage === 5"
        value="5"
        >5</option
      >
      <option
        :selected="itemPerPage === 10"
        value="10"
        >10</option
      >
      <option
        :selected="itemPerPage === 15"
        value="15"
        >15</option
      >
      <option
        :selected="itemPerPage == 20"
        value="20"
        >20</option
      >
      "
    </select>
  </div>
</template>
