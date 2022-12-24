<script lang="ts" setup>
import { onBeforeMount, ref } from 'vue';

export interface HistogramValue {
  label: any;
  value: number;
}

export interface HistogramProps {
  chartTitle: string;
  serieName: string;
  values: HistogramValue[];
}

const CHART_HEIGHT = 500;
const CHART_WIDTH = 800;

const GREEN = '#1eb980';
const RED = '#ff937f';

const props = defineProps<HistogramProps>();

const categories = ref<string[]>([]);
const data = ref<number[]>([]);
const serieName = ref<string>('');

const chartOptions = ref();
const series = ref<any[]>([]);

const processProps = () => {
  serieName.value = props.serieName;
  categories.value = props.values.map((x) => x.label);
  data.value = props.values.map((x) => x.value);

  chartOptions.value = {
    colors: ['#1eb980'],
    chart: {
      id: 'vuechart-example',
    },
    xaxis: {
      position: 'top',
      categories: categories.value,
      labels: {
        style: {
          colors: categories.value.map((_) => '#fff'),
          fontSize: '14px',
        },
      },
    },
    plotOptions: {
      bar: {
        borderRadius: 10,
        dataLabels: {
          position: 'top',
        },
      },
    },
  };

  series.value = [
    {
      name: serieName.value,
      data: data.value,
    },
  ];
};

onBeforeMount(() => {
  processProps();
});
</script>
<template>
  <div class="bg-foreground shadow rounded-lg p-6" :style="`height: ${CHART_HEIGHT + 30}px`">
    <div v-if="props.values.length !== 0">
      <h1 class="text-white font-semibold text-lg">{{ props.chartTitle }}</h1>
      <apexchart
        v-show="chartOptions"
        type="bar"
        :width="CHART_WIDTH"
        :height="CHART_HEIGHT"
        :options="chartOptions"
        :series="series"
      ></apexchart>
    </div>
    <div v-else class="flex flex-col h-full items-center justify-center">
      <h1 class="text-center text-2xl text-lgray font-bold">No data</h1>
      <h1 class="text-center text-6xl text-lgray font-bold m-3">:(</h1>
    </div>
  </div>
</template>
