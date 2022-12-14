<script setup lang="ts">
import BirdazzoneButton from '../components/BirdazzoneButton.vue';
</script>
<template>
  <main class="flex justify-center items-center" style="flex: 1 1 auto">
    <div class="flex">
      <a
        class="flex flex-center"
        href="https://twitter.com/intent/tweet?text=%23birdchess%20Quale%20mossa%20mi%20consigliate%3F&hashtags=chess%2Ctwitter%2Cbirdchess"
      >
        <BirdazzoneButton :text="'Ask for some suggestions about a chess match'" :active="true"
          >Ask for some suggestions about a chess match</BirdazzoneButton
        >
      </a>
    </div>
  </main>
</template>
