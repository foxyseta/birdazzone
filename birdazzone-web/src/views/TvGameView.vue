<script setup lang="ts">
import ApiRepository from '@/api/api-repository';
import type { TvGame } from '@/api/interfaces/tv-game';
import { ref, onBeforeMount } from 'vue'
import ErrorWidget from '../components/ErrorWidget.vue'
import WordCloud from '../components/WordCloud.vue'
import AerogramCard from '../components/AerogramCard.vue'

const error = ref<boolean> (false)
const game = ref<TvGame>()
const props = defineProps<{id: number}>()

const fetchGame = async () => {
    const resp = await ApiRepository.getTvGameById(props.id.toString())
    if (resp.esit) {
      game.value = resp.data
    } else {
      error.value = true
    }
  }


onBeforeMount(fetchGame)
</script>
<template>
  <!-- Error -->
  <div v-if="error">
    <ErrorWidget/>
  </div>
  <!-- Success -->
  <div v-else class="w-full flex flex-col justify-start items-center">
    <div class="rounded-lg text-white bg-lgreen py-3 px-9 my-8 m-3">
      {{game?.name.toUpperCase()}}
    </div>
    <div class="flex">
      <div>
        <!-- Insert here the list -->
      </div>
      <div class="flex flex-col justify-start">
        <AerogramCard :id="props.id" />
        <WordCloud :tv-game-id="props.id" />
      </div>
    </div>
  </div>
</template>
