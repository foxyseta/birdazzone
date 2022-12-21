<script lang="ts" setup>
import BirdazzoneButton from '../components/BirdazzoneButton.vue';
import { ref, onBeforeMount } from 'vue';
import FantaTeamsTab from './tabs/FantaTeamsTab.vue';
import PoliticiansTab from '@/views/tabs/PoliticiansTab.vue';
import NumberOfPoliticians from '../components/NumberOfPoliticians.vue';

const showPoliticiansTab = ref<boolean>(true);
const showTeamsTab = ref<boolean>(false);

const showPoliticians = () => {
  showPoliticiansTab.value = true;
  showTeamsTab.value = false;
};

const showTeams = () => {
  showPoliticiansTab.value = false;
  showTeamsTab.value = true;
};
</script>

<template>
  <div class="pl-4 w-full flex flex-col justify-start">
    <div class="flex flex-row">
      <div class="flex flex-col" style="flex: 1 1 auto"></div>
      <div
        class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold py-3 my-8 m-3"
        style="flex: 1 1 auto; text-align: center; max-width: 19rem"
      >
        FANTACITORIO
      </div>
      <div class="flex flex-col" style="flex: 1 1 auto"></div>
    </div>

    <div class="flex justify-start w-100 m-3">
      <BirdazzoneButton :text="'POLITICIANS'" :active="showPoliticiansTab" @click="showPoliticians"
        ></BirdazzoneButton
      >
      <BirdazzoneButton :text="'TEAMS'" :active="showTeamsTab" @click="showTeams"></BirdazzoneButton>
    </div>

    <div class="h-full">
      <div class="flex flex-row" v-show="showPoliticiansTab">
        <div>
          <PoliticiansTab class="flex" style="flex: 1 1 auto" />
        </div>
      </div>
      <div v-show="showTeamsTab">
        <FantaTeamsTab />
      </div>
    </div>
  </div>
</template>
