<script setup lang="ts">
import ApiRepository from '../api/api-repository';
import type { TvGame } from '../api/interfaces/tv-game';
import { ref, onBeforeMount } from 'vue';
import GuesserMap from '../components/GuesserMap.vue';
import BirdazzoneButton from '../components/BirdazzoneButton.vue';
import ErrorWidget from '@/components/ErrorWidget.vue';
import ListTab from './tabs/ListTab.vue';
import ChartTab from './tabs/ChartTab.vue';
import MapTab from './tabs/MapTab.vue';

const props = defineProps<{ id: string }>();

const loading = ref<boolean>(true);
const error = ref<boolean>(false);
const errorTitle = ref<string>();
const errorText = ref<string>();
const game = ref<TvGame>();

const showListTab = ref<boolean>(true);
const showMapTab = ref<boolean>(false);
const showChartTab = ref<boolean>(false);

const fetchGame = async () => {
  loading.value = true;
  const resp = await ApiRepository.getTvGameById(props.id);
  if (resp.esit) {
    game.value = resp.data;
    loading.value = false;
  } else if (resp.statusCode === 404) {
    window.location.href = '/not-found';
  } else {
    error.value = true;
    errorTitle.value = 'Error!';
    errorText.value = 'something went wrong!';
  }
};

const showList = () => {
  showListTab.value = true;
  showChartTab.value = false;
  showMapTab.value = false;
};

const showChart = () => {
  showListTab.value = false;
  showChartTab.value = true;
  showMapTab.value = false;
};

const showMap = () => {
  showListTab.value = false;
  showChartTab.value = false;
  showMapTab.value = true;
};

onBeforeMount(fetchGame);
</script>
<template>
  <!-- Error -->
  <div v-if="error" class="flex justify-center items-center w-full">
    <ErrorWidget :open="true" :title="errorTitle" :text="errorText" />
  </div>
  <!-- Success -->
  <div v-else class="pl-4 w-full flex flex-col justify-start">
    <!-- Title -->
    <div class="flex flex-row">
      <div class="flex flex-col" style="flex: 1 1 auto"></div>
      <div
        class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold py-3 my-8 m-3"
        style="flex: 1 1 auto; text-align: center; max-width: 18rem"
      >
        {{ game?.name.toUpperCase() }}
      </div>
      <div class="flex flex-col" style="flex: 1 1 auto"></div>
    </div>

    <!-- Buttons -->
    <div class="flex justify-start w-100 m-3">
      <BirdazzoneButton :text="'LIST'" :active="showListTab" @click="showList">ciao</BirdazzoneButton>
      <BirdazzoneButton :text="'MAP'" :active="showMapTab" @click="showMap">ciao</BirdazzoneButton>
      <BirdazzoneButton :text="'CHARTS'" :active="showChartTab" @click="showChart">ciao</BirdazzoneButton>
    </div>

    <!-- Content -->
    <div class="h-full mb-5">
      <div v-show="showListTab">
        <ListTab :game-id="props.id" />
      </div>
      <div v-show="showMapTab">
        <MapTab :game-id="props.id" />
      </div>
      <div v-show="showChartTab">
        <ChartTab :game-id="props.id" />
      </div>
    </div>
  </div>
</template>
