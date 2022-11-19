<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import ApiRepository from '@/api/api-repository';
import type { TvGame } from '@/api/interfaces/tv-game';
import ErrorWidget from '../components/ErrorWidget.vue';
import TimeFilter from '../components/TimeFilter.vue';

const error = ref<boolean>(false)
const games = ref<TvGame[]>([])

const fetchTvGames = async () => {
  const response = await ApiRepository.getTvGames()
  if (response.esit) {
    error.value = false
    games.value = response.data!
  } else {
    error.value = true
  }
}

onBeforeMount(fetchTvGames)
</script>
<template>
  <!-- Error -->
  <div v-if="error">
    <ErrorWidget />   
  </div>
  <div v-else class="text-center">
    <h1 class="text-4xl font-black text-lgreen">TV GAMES</h1>
  </div>
  <div class="flex flex-col justify-evenly">
    <div v-for="game in games">
      <RouterLink :to="`/tv-games/${game.id}`">
        <div class="rounded shadow bg-foreground py-9 px-9">
          <h1 class="text-white text-xl font-bold"> {{game.name}}</h1>
        </div>
      </RouterLink>
    </div>
  </div>
</template>
