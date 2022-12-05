<script setup lang="ts">
  import { onBeforeMount, ref } from 'vue'
  import ApiRepository from '@/api/api-repository';
  import type { TvGame } from '@/api/interfaces/tv-game';
  import ErrorWidget from '../components/ErrorWidget.vue';

const errorTitle = ref<string>()
const errorText = ref<string>()

const games = ref<TvGame[]>([])

  const BASE:string = `http://${import.meta.env.VITE_SERVER_URL}`
  const error = ref<boolean>(false)
  const logo_list = ref<TvGame[]>([])

  const fetchTvGames = async () => {
    const response = await ApiRepository.getTvGames()
    
    if (response.esit) {
      error.value = false
      games.value = response.data!
      fetchLogo()
    } else {
      error.value = true
    errorTitle.value = 'Error!'
    errorText.value = 'something went wrong!'
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
  <!-- Error -->
  <div v-if="error">
    <ErrorWidget :open="true" :title="errorTitle" :text="errorText"/>   
  </div>
  <div class="flex flex-row mb-10 h-screen"  style="flex: 1 1 auto;">

    <div class="flex flex-col"  style="flex: 1 1 auto"></div>
    
    <div class="flex flex-col" style="flex: 1 1 auto">
      <div v-if="error">
        <ErrorWidget :open="true" :title="errorTitle" :text="errorText" />
      </div>
      <div v-else class="flex flex-col justify-center w-full text-center">
        <div class="flex flex-row">
          <div class="flex flex-col"  style="flex: 1 1 auto"></div>
          <div class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold py-3 my-8 m-3" style="flex: 1 1 auto; max-width: 15rem;">TV GAMES</div>
          <div class="flex flex-col"  style="flex: 1 1 auto"></div>
        </div>
        <div class="flex flex-col items-center" v-for="(game,index) in games" :key="index">
          <RouterLink :to="`/tv-games/${game.id}`" style="flex: 1 1 auto; width: 25rem;">
            <div class="flex flex-row rounded-lg m-10 bg-foreground hover:bg-lgray/50 p-4 flex-1">
              <div class="flex flex-row px-3">
                <img :src="(logo_list[game.id] ? BASE + logo_list[game.id].logo : '/icons/user.svg')" style="height: 6rem" alt="gameIcon"/>
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
    <div class="flex flex-col"  style="flex: 1 1 auto"></div>
  </div>
</template>
ì
