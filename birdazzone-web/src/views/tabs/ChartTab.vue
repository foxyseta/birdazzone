<script lang="ts" setup>
import { ref } from 'vue';
import FilterA from '../../components/FilterAerogram.vue';
import FilterH from '../../components/FilterHistogram.vue';
import AerogramCard from '../../components/AerogramCard.vue';
import TvGameHistogram from '../../components/TvGameHistogram.vue';

const props = defineProps<{ gameId: string }>();

const fromA = ref<string | null>(null);
const toA = ref<string | null>(null);
const fromH = ref<string | null>(null);
const toH = ref<string | null>(null);
const aerogramKey = ref<number>(0);
const histogramKey = ref<number>(0);

const forceUpdateA = () => {
  aerogramKey.value += 1;
};

const forceUpdateH = () => {
  histogramKey.value += 1;
};

const a = (f: string | null, t: string | null) => {
  fromA.value = f;
  toA.value = t;
  forceUpdateA();
};

const h = (f: string | null, t: string | null) => {
  fromH.value = f;
  toH.value = t;
  forceUpdateH();
};
</script>

<template>
  <div class="pl-20 pr-20 grid grid-cols-5">
    <div class="mt-6 mb-4 place-self-start col-span-2 w-auto">
      <FilterA @change-from-to="a" />
    </div>
    <div class="mt-6 mb-4 place-self-start col-span-2 w-auto">
      <FilterH @change-from-to="h" class="justify-start" />
    </div>
    <div class="mt-6 mb-4"></div>
    <div class="m-3 mr-0 place-self-start col-span-2 w-auto">
      <AerogramCard :id="props.gameId" :key="aerogramKey" :from="fromA" :to="toA" />
    </div>
    <div class="m-3 ml-0 mr-20 col-span-3">
      <TvGameHistogram style="margin-left: 0" :game-id="props.gameId" :key="histogramKey" :from="fromH" :to="toH" />
    </div>
  </div>
</template>
