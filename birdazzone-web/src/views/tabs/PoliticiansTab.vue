<script lang="ts" setup>
  import { ref, onBeforeMount } from 'vue';
  import ErrorWidget from '@/components/ErrorWidget.vue';
  import ApiRepository from '../../api/api-repository';
  import type { Politician } from '../../api/interfaces/politician';
  import FantacitorioHistogram from '@/components/FantacitorioHistogram.vue';
  import { SemipolarSpinner } from 'epic-spinners';
  import FantaRankChart from '@/components/FantaRankChart.vue';
  import NumberOfPoliticians from '@/components/NumberOfPoliticians.vue';
  const errorTitle = ref<string>();
  const errorText = ref<string>();

  const error = ref<boolean>(false);
  const loading = ref<boolean>(true);
  const list = ref<Politician[]>([]);

  const fetchPoliticiansList = async () => {
    loading.value = true;
    const resp = await ApiRepository.getPoliticians();
    if (resp.esit) {
      list.value = resp.data!;
      loading.value = false;
    } else if (resp.statusCode === 404) {
      window.location.href = '/not-found';
    } else {
      error.value = true;
      errorTitle.value = 'Error!';
      errorText.value = 'something went wrong!';
    }
  };

  function isNumber(value: string): boolean {
    if (typeof value !== 'string') {
      return false;
    }

    if (value.trim() === '') {
      return false;
    }

    return !isNaN(Number(value));
  }

  function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  const changeAndSort = async (index: number, newScore: string) => {
    loading.value = true;
    await sleep(1000);
    list.value[index].score = isNumber(newScore) ? parseInt(newScore) : list.value[index].score;
    list.value = list.value.sort((a, b) => b.score - a.score);
    loading.value = false;
    return list.value[index].score.toString();
  };

  onBeforeMount(() => {
    fetchPoliticiansList();
  });
</script>

<template>
  <div
    class="flex flex-col flex-1 my-3 w-full"
    style="flex: 1 1 auto"
  >
    <div v-if="error">
      <ErrorWidget
        :open="true"
        :title="errorTitle"
        :text="errorText"
      />
    </div>
    <div
      v-if="list.length > 0"
      class="flex flex-row w-full"
      style="flex: 1 1 auto"
    >
      <div
        class="flex flex-col"
        style="flex: 1 1 auto; width: 22rem"
      >
        <div
          v-for="(item, index) in list"
          :key="index"
        >
          <div class="flex flex-row my-3 text-lgray">
            <div
              class="flex"
              style="flex: 1 1 auto; width: 1rem"
              >{{ index + 1 }}.</div
            >
            <div class="flex flex-row rounded-lg bg-foreground p-4 ml-3 items-center">
              <div
                class="font-semibold text-white"
                style="flex: 1 1 auto; width: 100%"
              >
                {{ item.name }}
              </div>
              <div
                class="flex font-bold text-white text-right align-center justify-end"
                style="flex: 1 1 auto; font-size: 180%"
              >
                <input
                  type="text"
                  @change="
                    value => {
                      //@ts-ignore
                      value = changeAndSort(index, value.target?.value);
                    }
                  "
                  class="flex bg-foreground justify-end flex-wrap text-end"
                  style="width: 100%"
                  :value="item.score"
                />
                <span
                  class="flex text-lgray font-normal items-center ml-1"
                  style="font-size: 50%"
                  >p.</span
                >
              </div>
              <div
                class="flex justify-center ml-3"
                style="flex: 1 0 auto; height: 2rem; width: 2rem"
              >
                <img
                  v-show="index === 0"
                  :src="'/icons/coccarda1.svg'"
                  alt="medal1"
                />
                <img
                  v-show="index === 1"
                  :src="'/icons/coccarda2.svg'"
                  alt="medal2"
                />
                <img
                  v-show="index === 2"
                  :src="'/icons/coccarda3.svg'"
                  alt="medal3"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="flex ml-20 justify-center align-center flex-1">
        <div
          v-show="loading"
          class="flex justify-center items-center"
        >
          <semipolar-spinner
            :animation-duration="2000"
            :size="70"
            color="#1eb980"
          />
        </div>
        <div class="flex flex-col items-center">
          <NumberOfPoliticians
            class="flex align-center m-2 mb-3"
            style="flex: 1 1 auto"
            v-if="!loading"
            :num="list.length"
          />
          <div
            class="flex align-center m-2 mb-6"
            style="flex: 1 1 auto"
            v-if="!loading"
          >
            <FantaRankChart :list="list" />
          </div>
          <div
            class="flex align-center mt-2"
            style="flex: 1 1 auto"
            v-if="!loading"
          >
            <FantacitorioHistogram :list="list" />
          </div>
        </div>
      </div>
    </div>
    <div
      v-else
      class="m-5 text-white font-semibold"
    >
      <div v-if="!loading">No politicians ranked &#9203;</div>
    </div>
  </div>
</template>
