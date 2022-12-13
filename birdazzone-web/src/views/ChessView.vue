<script setup lang="ts">
    import BirdazzoneButton from '../components/BirdazzoneButton.vue'
</script>
<template>
  <main class="flex justify-center items-center" style="flex: 1 1 auto">
    <div class="flex ">
      <a class="flex flex-center" href="https://twitter.com/intent/tweet">
        <BirdazzoneButton :text="'Ask for some suggestions about a chess match'" :active="true">Ask for some suggestions about a chess match</BirdazzoneButton>
      </a>
    </div>
  </main>
</template>
