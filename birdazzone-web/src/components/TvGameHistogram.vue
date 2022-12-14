<script lang="ts" setup>
import Histogram, { type HistogramValue } from './Histogram.vue';
import ApiRepository from '../api/api-repository';
import { onBeforeMount, ref } from 'vue';

const props = defineProps<{ gameId: string; from: string | null; to: string | null }>();

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
  <div>
    <Histogram
      v-if="!loading"
      :chart-title="'Played words with more than 3 tentatives'"
      :values="histogramValues"
      :serie-name="SERIE_NAME"
    />
  </div>
</template>
