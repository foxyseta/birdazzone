<script lang="ts" setup>
import { ref, onBeforeMount } from 'vue';
import ErrorWidget from '@/components/ErrorWidget.vue';
import ApiRepository from '../../api/api-repository';
import type { Politician } from '../../api/interfaces/politician';

const error = ref<boolean>(false);
const loading = ref<boolean>(true);
const list = ref<Politician[]>([
  /*{
            "name": "Giorgia Meloni",
            "score": 1000
        },
        {
            "name": "Matteo Salvini",
            "score": 900
        },
        {
            "name": "Mattia di Maio",
            "score": 800
        },
        {
            "name": "Carlo Calenda",
            "score": 700
        },
        {
            "name": "Matteo Renzi",
            "score": 600
        },
        {
            "name": "Silvio Berlusconi",
            "score": 500
        },
        {
            "name": "Gianni Agnelli",
            "score": 400
        },*/
]);

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
  }
};

function sortList() {
  list.value.sort((a, b) => b.score - a.score);
}

function isNumber(value: string): boolean {
  if (typeof value !== 'string') {
    return false;
  }

  if (value.trim() === '') {
    return false;
  }

  return !isNaN(Number(value));
}

function changeAndSort(index: number, newScore: string) {
  list.value[index].score = isNumber(newScore) ? parseInt(newScore) : list.value[index].score;
  sortList();
  return list.value[index].score.toString();
}

onBeforeMount(fetchPoliticiansList);
</script>

<template>
  <div class="flex flex-col flex-1 my-3" style="flex: 1 1 auto">
    <div class="flex flex-row w-full" style="flex: 1 1 auto">
      <div class="flex flex-col" style="flex: 2 1 auto"></div>
      <div class="flex flex-col justify-center align-center" style="flex: 1 1 auto; width: 22rem">
        <div v-for="(item, index) in list" :key="index">
          <div class="flex flex-row my-2 text-lgray">
            <div class="flex" style="flex: 1 1 auto; width: 1rem">{{ index + 1 }}.</div>
            <div class="flex flex-row rounded-lg bg-foreground p-4 ml-3 items-center">
              <div class="font-semibold text-white" style="flex: 1 1 auto; width: 100%">
                {{ item.name }}
              </div>
              <div
                class="flex font-bold text-white text-right align-center justify-end"
                style="flex: 1 1 auto; font-size: 180%"
              >
                <input
                  type="text"
                  @change="
                    (value) => {
                      //@ts-ignore
                      value = changeAndSort(index, value.target?.value);
                    }
                  "
                  class="flex bg-foreground justify-end flex-wrap text-end"
                  style="width: 100%"
                  :value="item.score"
                />

                <span class="flex text-lgray font-normal items-center ml-1" style="font-size: 50%">p.</span>
              </div>
              <div class="flex justify-center ml-3" style="flex: 1 0 auto; height: 2rem; width: 2rem">
                <img v-show="index === 0" :src="'/icons/coccarda1.svg'" alt="medal1" />
                <img v-show="index === 1" :src="'/icons/coccarda2.svg'" alt="medal2" />
                <img v-show="index === 2" :src="'/icons/coccarda3.svg'" alt="medal3" />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="flex flex-col" style="flex: 2 1 auto"></div>
    </div>
  </div>
</template>
