<script setup lang="ts">
import ApiRepository from '@/api/api-repository';
import type { TvGame } from '@/api/interfaces/tv-game';
import { ref, onBeforeMount } from 'vue'
import ErrorWidget from '../components/ErrorWidget.vue'
import WordCloud from '../components/WordCloud.vue'
import AerogramCard from '../components/AerogramCard.vue'
import GuesserList from '@/components/GuesserList.vue'
import { SemipolarSpinner } from 'epic-spinners';

const loading = ref<boolean> (true)
const game = ref<TvGame>()
const props = defineProps<{id: number}>()

const fetchGame = async () => {
    loading.value = true
    const resp = await ApiRepository.getTvGameById(props.id.toString())
    if (resp.esit) {
      game.value = resp.data
      loading.value = false
    } else {
      window.location.href = "/not-found"   
    }
  }


onBeforeMount(fetchGame)
</script>
<template>
  <!-- Loading -->
  <div class="w-full h-full flex justify-center items-center" v-if="loading">
    <semipolar-spinner :animation-duration="2000" :size="50" color="#1eb980" />
  </div>
  <!-- Success -->
  <div v-else class="w-full flex flex-col justify-start items-center">
    <div class="rounded-lg text-white bg-lgreen text-4xl font-semibold py-3 px-9 my-8 m-3">
      {{game?.name.toUpperCase()}}
    </div>
    <div class="w-full flex justify-evenly">
        <GuesserList :game-id="props.id"/>
      <div class="flex flex-col justify-start">
        <div class="mt-3">
          <AerogramCard :id="props.id" />
        </div>
        <div class="mt-5">
          <WordCloud :tv-game-id="props.id" />
        </div>  
      </div>
    </div>
  </div>
</template>
