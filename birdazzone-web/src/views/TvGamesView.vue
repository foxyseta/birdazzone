<script setup lang="ts">
import { EREDITA_PATH, CATENA_PATH} from '../router/index'

interface TvGame {
  name: string
  url: string
}
const games = [
  {
    name: "L'eredità",
    url: EREDITA_PATH
  },
  {
    name: "Reazione a catena",
    url: CATENA_PATH
  }
]
</script>
<template>
  <div class="text-center">
    <h1 class="text-4xl font-black text-lgreen">TV GAMES</h1>
  </div>
  <div class="flex flex-col justify-evenly">
    <div v-for="game in games">
      <a :href="game.url">
        <div class="rounded shadow bg-foreground py-9 px-9">
          <h1 class="text-white text-xl font-bold"> {{game.name}}</h1>
        </div>
      </a>
    </div>
  </div>
</template>
