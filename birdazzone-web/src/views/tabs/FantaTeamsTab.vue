<script lang="ts" setup>
import FantaTeamsList from '../../components/FantaTeamsList.vue';
import FantaSearchUser from '../../components/FantaSearchUser.vue';
import type { FantaTeam } from '../../api/interfaces/fanta-team';
import { useRoute } from 'vue-router';
import { onBeforeMount, ref } from 'vue';
import ApiRepository from '@/api/api-repository';
import { SemipolarSpinner } from 'epic-spinners';

const username = ref<string>();
const teams = ref<FantaTeam[]>([]);
const isError = ref<boolean>(false);
const error = ref<string>();
const loading = ref<boolean>(false);

const fetchTeams = async () => {
  const mock = username.value ? teams.value.filter((x: FantaTeam) => x.username.startsWith(username.value!)) : teams;
  const response = await ApiRepository.getFantacitorioTeams(username.value);
  isError.value = response.esit;
  if (response.esit && response.data) {
    teams.value = response.data;
  } else {
    error.value = response.error?.message;
  }
};

onBeforeMount(async () => {
  loading.value = true;
  await fetchTeams();
  loading.value = false;
});

const onUsernameChanged = async (newUsername: string) => {
  loading.value = true;
  username.value = newUsername;
  await fetchTeams();
  loading.value = false;
};
</script>

<template>
  <div class="h-full">
    <FantaSearchUser class="m-10" @changed="onUsernameChanged" />
    <div v-show="loading" class="h-screen flex justify-center items-center">
      <semipolar-spinner :animation-duration="2000" :size="70" color="#1eb980" />
    </div>
    <FantaTeamsList v-show="!loading" :teams="teams" />
  </div>
</template>
