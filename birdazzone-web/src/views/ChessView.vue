<script setup lang="ts">
import BirdazzoneButton from '../components/BirdazzoneButton.vue';
</script>
<template>
  <main class="flex flex-col w-full" style="flex: 1 1 auto">
    <div class="flex flex-row">
      <div class="flex flex-col" style="flex: 1 1 auto"></div>
      <div
        class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold p-3 my-8 m-3"
        style="flex: 1 1 auto; text-align: center; max-width: 19rem"
      >
        CHESS
      </div>
      <div class="flex flex-col" style="flex: 1 1 auto"></div>
    </div>
    <div class="flex justify-center h-full items-center">
      <a
        class="flex flex-center"
        href="https://twitter.com/intent/tweet?text=Quale%20mossa%20mi%20consigliate%3F&hashtags=birdchess%2Cchess%2Chelp%2Cbirdazzone"
        target="_blank"
        rel="noopener noreferrer"
      >
        <BirdazzoneButton :text="'Ask for some suggestions about a chess match'" :active="true"
          >Ask for some suggestions about a chess match</BirdazzoneButton
        >
      </a>
    </div>
  </main>
</template>
