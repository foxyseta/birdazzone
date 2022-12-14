<script lang="ts" setup>
import type { Tweet } from '../api/interfaces/tweet';

const props = defineProps<{ tweets: Tweet[] }>();
</script>

<template>
  <div v-if="tweets.length % 2 == 0" class="mx-10 md:mx-20 grid gap-4 md:gap-10 grid-cols-1 md:grid-cols-2">
    <div v-for="(item, index) in tweets" :key="index">
      <img :src="item.medias[0]" :alt="'@' + item.author.username + ' fantacitorio team'" />
    </div>
  </div>
  <div v-else class="mx-10 md:mx-20 grid gap-4 lg:gap-10 grid-cols-1 lg:grid-cols-3">
    <div v-for="(item, index) in tweets" :key="index">
      <img class="fantaTeam" :src="item.medias[0]" :alt="'@' + item.author.username + ' fantacitorio team'" />
    </div>
  </div>
</template>

<style>
.fantaTeam {
  filter:  hue-rotate(110deg) brightness(85%);
}
</style>
