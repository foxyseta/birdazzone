<script lang="ts" setup>
  import {ref} from 'vue';
  import FilterA from '../../components/FilterAerogram.vue';
  import FilterH from '../../components/FilterHistogram.vue';
  import FilterL from '../../components/FilterTimeChart.vue';
  import AerogramCard from '../../components/AerogramCard.vue';
  import TvGameTimeChart from '../../components/TvGameTimeChart.vue';
  import TvGameHistogram from '../../components/TvGameHistogram.vue';

  const props = defineProps<{gameId: string}>();

  const fromA = ref<string | null>(null);
  const toA = ref<string | null>(null);
  const dateL = ref<string | null>(null);
  const fromH = ref<string | null>(null);
  const toH = ref<string | null>(null);
  const aerogramKey = ref<number>(0);
  const lineKey = ref<number>(0);
  const histogramKey = ref<number>(0);

  const forceUpdateA = () => {
    aerogramKey.value += 1;
  };

  const forceUpdateL = () => {
    lineKey.value += 1;
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

  const l = (d: string | null) => {
    dateL.value = d;
    forceUpdateL();
  };
</script>

<template>
  <div class="mb-10 pl-20 pr-20 flex justify-between flex-wrap">
    <div>
      <!-- Aereogram -->
      <div class="mt-6 mb-4 place-self-start w-auto">
        <FilterA @change-from-to="a" />
      </div>
      <div class="m-3 mr-0 place-self-start w-auto">
        <AerogramCard
          :id="props.gameId"
          :key="aerogramKey"
          :from="fromA"
          :to="toA"
        />
      </div>
    </div>

    <!-- Line -->
    <div>
      <div class="mt-6 mb-4 place-self-start w-auto">
        <FilterL
          @change-date="l"
          class="justify-start"
        />
      </div>
      <div class="m-3 mr-0 place-self-start w-auto">
        <TvGameTimeChart
          :game-id="props.gameId"
          :key="lineKey"
          :date="dateL"
        />
      </div>
    </div>

    <!-- Histogram -->
    <div>
      <div class="mt-6 mb-4 place-self-start col-span-2 w-auto">
        <FilterH
          @change-from-to="h"
          class="justify-start"
        />
      </div>
      <div class="m-3 ml-0 mr-20 col-span-3">
        <TvGameHistogram
          style="margin-left: 0"
          :game-id="props.gameId"
          :key="histogramKey"
          :from="fromH"
          :to="toH"
        />
      </div>
    </div>
  </div>
</template>
