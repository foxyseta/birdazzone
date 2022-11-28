<script setup lang="ts">
import ApiRepository from '@/api/api-repository';
import type { TvGame } from '@/api/interfaces/tv-game';
import { ref, onBeforeMount } from 'vue'
import WordCloud from '../components/WordCloud.vue'
import GuesserMap from '../components/GuesserMap.vue'
import AerogramCard from '../components/AerogramCard.vue'
import BirdazzoneButton from '../components/BirdazzoneButton.vue'
import ErrorWidget from '@/components/ErrorWidget.vue'
import GuesserList from '@/components/GuesserList.vue'
import { SemipolarSpinner } from 'epic-spinners';

const props = defineProps<{id: string}>()

const loading = ref<boolean> (true)
const error = ref<boolean> (false)
const game = ref<TvGame>()
const showList = ref<boolean>(true)

const fetchGame = async () => {
    loading.value = true
    const resp = await ApiRepository.getTvGameById(props.id)
    if (resp.esit) {
      game.value = resp.data
      loading.value = false
    } else if (resp.statusCode === 404) {
      window.location.href = "/not-found"   
    } else {
      error.value = true
    }
  }

const showMap = () => { showList.value = false }
const hideMap = () => { showList.value = true }

onBeforeMount(fetchGame)
</script>
<template>
  <!-- Error -->
  <div v-if="error" class="flex justify-center items-center w-full">
    <ErrorWidget />
  </div>
  <!-- Success -->
  <div v-else class="pl-4 w-full flex flex-col justify-start">
    <!-- Title -->
    <div class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold py-3 px-9 my-8 m-3">
      {{game?.name.toUpperCase()}}
    </div>

    <!-- Buttons -->
    <div class="flex justify-start w-100 m-3">
      <BirdazzoneButton :text="'LIST'" :active="showList" @click="hideMap">ciao</BirdazzoneButton>
      <BirdazzoneButton :text="'MAP'" :active="!showList" @click="showMap">ciao</BirdazzoneButton>
    </div>

    <!-- Content -->
    <div class="w-full flex justify-evenly flex-wrap">
      <div v-show="showList">
        <GuesserList :game-id="props.id"/>
      </div>
      <div v-show="!showList" >
        <!--<GuesserMap :game-id="props.id" />-->
        <GuesserMap/>
      </div>
      <div class="flex flex-col justify-start">
        <div class="mt-3">
          <AerogramCard :id="props.id" />
          <WordCloud :tv-game-id="props.id" />
        </div>
      </div>
    </div>
  </div>
</template>
