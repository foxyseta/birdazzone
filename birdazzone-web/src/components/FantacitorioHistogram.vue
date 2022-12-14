<script lang="ts" setup>
import Histogram, { type HistogramValue } from './Histogram.vue';
import { onBeforeMount, onBeforeUpdate, ref } from 'vue';
import type { Politician } from '@/api/interfaces/politician';

const SERIE_NAME = 'Politicians';
const CHART_TITLE = "Politician distribution TODO"
const COLUMN_NO = 10
const UNIT_BIN = 100

const props = defineProps<{list: Politician[]}>()
const histogramValues = ref<HistogramValue[]>([]);
const error = ref<boolean>(false);
const errorMessage = ref<string>('');
const loading = ref<boolean>(false);

const mapData = (scores: Politician[]) => {
  const step = Math.round((Math.max(...scores.map(x => x.score)) / COLUMN_NO) / UNIT_BIN) * UNIT_BIN
  const categories = new Array<number>(COLUMN_NO).fill(0)
  for (const score of scores) {
    const index = Math.trunc(score.score / step);
    categories[index < COLUMN_NO ? index : COLUMN_NO - 1]++;
  }

  histogramValues.value = categories.map((val, index) => ({ label: 
    index === COLUMN_NO -1
    ? `>${step * index}`
    :`${step * index}-${step * (index + 1)-1}`, value: val }));
};

onBeforeMount(() => {
  mapData(props.list)
})
</script>
<template>
  <Histogram
    v-if="!loading"
    :chart-title="CHART_TITLE"
    :values="histogramValues"
    :serie-name="SERIE_NAME"
  />

  <!-- Error message -->
  <h1 v-show="error">{{ errorMessage }}</h1>
</template>
