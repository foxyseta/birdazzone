<script lang="ts" setup>
import { ref, onBeforeMount } from 'vue';
import ApiRepository from '../api/api-repository';
import type { Results } from '../api/interfaces/results';
import { SemipolarSpinner } from 'epic-spinners';

const props = defineProps<{ gameId: string; date: string | null }>();

const CHART_TITLE = 'Time chart';
const CHART_HEIGHT = 500;
const CHART_WIDTH = 800;
const DEFAULT_EACH = '60';
const GREEN = '#1eb980';
const RED = '#ff937f';

const loading = ref<boolean>(false);
const error = ref<boolean>(false);
const errorMessage = ref<string>('');

const chartOptions = ref({
  colors: [GREEN, RED],
  chart: {
    type: 'line',
  },
  stroke: {
    curve: 'smooth',
  },
  series: [],
  legend: {
    labels: {
      colors: ['#aaa'],
    },
  },
  dataLabels: {
    style: {
      colors: [GREEN, RED],
    },
  },
  xaxis: {
    tooltip: {
      shared: false,
      intersect: true,
      x: {
        show: false,
      },
    },
    type: 'datetime',
    axisTicks: {
      show: true,
    },
    labels: {
      style: {
        colors: '#aaa',
      },
    },
  },
});

const fromResultToChartPoint = (r: Results, positive: boolean) => ({
  x: new Date(r.label).getTime(),
  y: positive ? r.positives : r.negatives,
});

const fetchAttempts = async () => {
  loading.value = true;
  chartOptions.value.series = [];

  const response = props.date
    ? await ApiRepository.getResultsFiltered(props.gameId, props.date, null, DEFAULT_EACH)
    : await ApiRepository.getResults(props.gameId, DEFAULT_EACH);
  error.value = response.esit;
  if (response.esit) {
    chartOptions.value.series.push({
      // @ts-ignore
      name: 'Succeded',
      // @ts-ignore
      data: response.data!.map((x) => fromResultToChartPoint(x, true)),
    });
    chartOptions.value.series.push({
      // @ts-ignore
      name: 'Failed',
      // @ts-ignore
      data: response.data!.map((x) => fromResultToChartPoint(x, false)),
    });
  } else {
    errorMessage.value = response.error!.message;
    loading.value = false;
  }
  loading.value = false;
};

onBeforeMount(() => {
  fetchAttempts();
});
</script>
<template>
  <div
    class="bg-foreground shadow rounded-lg p-6"
    :style="`width: ${CHART_WIDTH + 50}px; height: ${CHART_HEIGHT + 80}px`"
  >
    <div v-if="loading" class="flex p-5 w-50 h-full items-center justify-center">
      <semipolar-spinner :animation-duration="2000" :size="50" color="#1eb980" />
    </div>

    <div v-if="!loading && chartOptions.series.length !== 0">
      <h1 class="text-white font-semibold text-lg">{{ CHART_TITLE }}</h1>
      <apexchart
        v-show="chartOptions"
        type="area"
        :width="CHART_WIDTH"
        :height="CHART_HEIGHT"
        :options="chartOptions"
        :series="chartOptions.series"
      ></apexchart>
    </div>
    <div
      v-if="!loading && chartOptions.series.length === 0"
      class="flex flex-col w-full h-full items-center justify-center"
    >
      <h1 class="text-center text-2xl text-lgray font-bold">No data</h1>
      <h1 class="text-center text-6xl text-lgray font-bold m-3">:(</h1>
    </div>
  </div>
</template>
