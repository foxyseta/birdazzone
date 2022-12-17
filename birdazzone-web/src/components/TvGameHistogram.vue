<script lang="ts" setup>
import Histogram, { type HistogramValue } from './Histogram.vue';
import ApiRepository from '../api/api-repository';
import { onBeforeMount, ref } from 'vue';

const props = defineProps<{ gameId: string; key: number; from: string | null; to: string | null }>();

const SERIE_NAME = 'attempts';

const histogramValues = ref<HistogramValue[]>([]);
const error = ref<boolean>(false);
const errorMessage = ref<string>('');
const loading = ref<boolean>(false);

const fetchAttempts = async () => {
  loading.value = true;

  const response = await ApiRepository.getTvGameAttemptsStat(props.gameId, props.from, props.to);
  error.value = response.esit;
  if (response.esit) {
    histogramValues.value = response
      .data!.filter((x) => x.absoluteFrequency > 3)
      .map(
        (x) =>
          ({
            label: x.value,
            value: x.absoluteFrequency,
          } as HistogramValue)
      );
  } else {
    errorMessage.value = response.error!.message;
  }
  loading.value = false;
};

onBeforeMount(() => {
  fetchAttempts();
});
</script>
<template>
  <div class="h-full">
    <div v-if="from && to" class="text-md text-white text-semibold mb-3">
      Data refere to the following date-time range: from {{ from.substring(0, 10) }}, {{ from.substring(11, 16) }} to
      {{ to.substring(0, 10) }}, {{ to.substring(11, 16) }}
    </div>
    <Histogram
      v-if="!loading && histogramValues.length > 0"
      :chart-title="'Played words with more than 3 tentatives'"
      :values="histogramValues"
      :serie-name="SERIE_NAME"
    />
    <div v-if="!loading && histogramValues.length === 0" class="bg-foreground shadow rounded p-6">
      <h1 class="font-bold text-lgray text-xl text-center">No data :(</h1>
    </div>
  </div>
</template>
