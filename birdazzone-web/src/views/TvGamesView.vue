<script setup lang="ts">
  import { onBeforeMount, ref } from 'vue'
  import ApiRepository from '@/api/api-repository';
  import type { TvGame } from '@/api/interfaces/tv-game';
  import ErrorWidget from '../components/ErrorWidget.vue';

  //const BASE:string = `${import.meta.env.VITE_SERVER_URL}`
  const error = ref<boolean>(false)
  const games = ref<TvGame[]>([])
  const logo_list = ref<TvGame[]>([])

  const fetchTvGames = async () => {
    const response = await ApiRepository.getTvGames()
    
    if (response.esit) {
      error.value = false
      games.value = response.data!
      fetchLogo()
    } else {
      error.value = true
    }
  }

  const fetchLogo = async () => {
    for(var game in games.value) {
      logo_list.value.push((await ApiRepository.getTvGameById(games.value[game].id.toString())).data!) 
      
    }
  }

  onBeforeMount(fetchTvGames)
</script>

<template>
  <div class="flex flex-row"  style="flex: 1 1 auto">

    <div class="flex flex-col"  style="flex: 1 1 auto"></div>
    
    <div class="flex flex-col" style="flex: 1 1 auto">
      <div v-if="error">
          <ErrorWidget />   
      </div>
      <div v-else class="flex flex-col justify-center w-full text-center">
        <h1 class="text-4xl mt-10 mb-14 font-black text-lgreen">TV GAMES</h1>
      
        <div class="flex flex-col items-center" v-for="(game,index) in games" :key="index">
          <RouterLink :to="`/tv-games/${game.id}`" style="flex: 1 1 auto; width: 25rem;">
            <div class="flex flex-row rounded-lg m-10 bg-foreground hover:bg-lgray/50 p-4 flex-1">
              <div class="flex flex-row px-3">
                <img :src="(logo_list[game.id] ? 'http://localhost:8080' + logo_list[game.id].logo : '/icons/user.svg')" style="height: 6rem" />
              </div>
              
              <div class="flex flex-col flex-1"></div>
              <div class="flex flex-col justify-center align-center">
                <div class="text-white font-semibold p-3"> {{game.name.toUpperCase()}}</div>
                <div class="text-lgray"> {{game.hashtag}}</div>
              </div>
              <div class="flex flex-col flex-1"></div>
            </div>
          </RouterLink>
        </div>
      </div> 
    </div>

    <div class="felx flex-col"  style="flex: 1 1 auto"></div>
  </div>
</template>





<!--div class="flex flex-col justify-center w-full text-center">
  <div v-if="error">
    <ErrorWidget />   
  </div>

  <div v-else class="flex justify-center w-full text-center">
    <h1 class="text-4xl font-black text-lgreen">TV GAMES</h1>
    <img :src="(logo_list[0] ? 'http://localhost:8080' + logo_list[0].logo : '')" alt="ghigliottina img"/>
  </div>

  <div class="flex flex-col justify-evenly h-screen">
    <div v-for="game in games" :key="game.name">
      <RouterLink :to="`/tv-games/${game.id}`">
        <div class="rounded shadow py-9 px-9">
          <h1 class="text-white text-xl font-bold"> {{game.name}}</h1>
        </div>
      </RouterLink>
    </div>
  </div>
</div-->