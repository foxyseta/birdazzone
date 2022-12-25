<script setup lang="ts">
  import GuesserListItem from '../components/GuesserListItem.vue';
  import ApiRepository from '../api/api-repository';
  import { onBeforeMount, ref } from 'vue';
  import type { Tweet } from '../api/interfaces/tweet';
  import { SemipolarSpinner } from 'epic-spinners';
  import PaginationBar from '../components/PaginationBar.vue';
  import CardPerPage from './ItemPerPage.vue';
  import ErrorWidget from './ErrorWidget.vue';

  const loading = ref<boolean>(false);
  const list = ref<Tweet[]>([]);
  const props = defineProps<{ gameId: string; key: number; from: string | null; to: string | null }>();
  const max = ref<number>(0);
  const actualPage = ref<number>(1);
  const itemPerPage = ref<number>(5);
  const firstLoad = ref<boolean>(true);

  const from = ref<string | null>(props.from);
  const to = ref<string | null>(props.to);

  const error = ref<boolean>(false);
  const errorTitle = ref<string>();
  const errorText = ref<string>();

  const fetchList = async () => {
    loading.value = true;
    if (!from.value || !to.value) {
      const resp = await ApiRepository.getListOfGuesser(
        props.gameId,
        actualPage.value.toString(),
        itemPerPage.value.toString()
      );
      if (resp.esit) {
        list.value = resp.data!.entries;
        max.value = resp.data!.numberOfPages;
        loading.value = false;
        firstLoad.value = false;
      } else {
        if (resp.statusCode === 204) {
          error.value = true;
          errorTitle.value = 'There are no tweets!';
          errorText.value = 'No game instance has been played!';
        } else {
          error.value = true;
          errorTitle.value = 'Error!';
          errorText.value = 'something went wrong!';
        }
      }
    } else {
      const resp = await ApiRepository.getListOfGuesserFiltered(
        props.gameId,
        from.value,
        to.value,
        actualPage.value.toString(),
        itemPerPage.value.toString()
      );
      if (resp.esit) {
        list.value = resp.data!.entries;
        max.value = resp.data!.numberOfPages;
        loading.value = false;
        firstLoad.value = false;
      } else {
        if (resp.statusCode === 204) {
          error.value = true;
          errorTitle.value = 'There are no tweets!';
          errorText.value = 'No game instance has been played!';
        } else {
          error.value = true;
          errorTitle.value = 'Error!';
          errorText.value = 'something went wrong!';
        }
      }
    }
  };
  function changeActual(n: number) {
    actualPage.value = n;
    fetchList();
  }
  onBeforeMount(fetchList);
</script>

<template>
  <!-- Error -->
  <div
    v-if="error"
    class="flex justify-center items-center w-full"
  >
    <ErrorWidget
      :open="true"
      :title="errorTitle"
      :text="errorText"
    />
  </div>

  <!-- Non error -->
  <div class="flex flex-col">
    <CardPerPage
      class="mb-3 flex justify-start"
      v-show="!firstLoad"
      :itemPerPage="itemPerPage"
      @change-item-page="
        n => {
          itemPerPage = n;
          actualPage = 1;
          fetchList();
        }
      "
    />
    <div
      v-show="loading"
      class="h-screen flex items-center"
    >
      <semipolar-spinner
        :animation-duration="2000"
        :size="50"
        color="#1eb980"
      />
    </div>
    <div v-show="!loading">
      <div
        v-if="list.length === 0"
        class="flex flex-1 my-3 text-white font-semibold"
        style="flex: 1 1 auto"
      >
        Nobody was smart enough to guess &#127773;
      </div>
      <div
        v-else
        class="flex flex-col"
        v-for="(item, index) in list"
        :key="index"
      >
        <GuesserListItem
          :data="item"
          :index="index + (actualPage - 1) * itemPerPage + 1"
          class="flex"
          style="flex: 1 1 auto; width: 30rem"
        />
      </div>
    </div>
    <PaginationBar
      class="mt-2"
      v-show="!firstLoad"
      :actualPage="actualPage"
      :max="max"
      @change-actual="changeActual"
    />
  </div>
</template>
