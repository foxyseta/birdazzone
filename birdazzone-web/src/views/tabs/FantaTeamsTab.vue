<script lang="ts" setup>
  import FantaTeamsList from '../../components/FantaTeamsList.vue';
  import FantaSearchUser from '../../components/FantaSearchUser.vue';
  import type { FantaTeam } from '../../api/interfaces/fanta-team';
  import { onBeforeMount, ref } from 'vue';
  import ApiRepository from '@/api/api-repository';
  import { SemipolarSpinner } from 'epic-spinners';

  const teams = ref<FantaTeam[]>([]);
  const cache = ref<FantaTeam[]>([]);
  const isError = ref<boolean>(false);
  const error = ref<string>();
  const loading = ref<boolean>(false);

  const cacheTeams = async () => {
    const response = await ApiRepository.getFantacitorioTeams();
    isError.value = response.esit;
    if (response.esit && response.data) {
      cache.value = response.data;
    } else {
      error.value = response.error?.message;
    }
  };

  onBeforeMount(async () => {
    loading.value = true;
    await cacheTeams();
    teams.value = cache.value;
    loading.value = false;
  });

  const onUsernameChanged = async (newUsername: string) => {
    loading.value = true;
    teams.value = cache.value.filter((x: FantaTeam) => x.username.toLowerCase().startsWith(newUsername.toLowerCase()));
    loading.value = false;
  };
</script>

<template>
  <div class="h-full">
    <FantaSearchUser
      class="mt-0 my-10"
      @changed="onUsernameChanged"
    />
    <div
      v-show="loading"
      class="h-screen flex justify-center items-center"
    >
      <semipolar-spinner
        :animation-duration="2000"
        :size="70"
        color="#1eb980"
      />
    </div>
    <FantaTeamsList
      v-show="!loading"
      :teams="teams"
    />
  </div>
</template>
