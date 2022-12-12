
<script lang="ts" setup>
import { onBeforeMount, ref } from 'vue';

export interface HistogramValue {
  label: any
  value: number
}

export interface HistogramProps {
  chartTitle: string
  serieName: string
  values: HistogramValue[]
}

const CHART_HEIGHT = 500
const CHART_WIDTH = 800

const props = defineProps<HistogramProps>()

const categories = ref<string[]>([])
const data = ref<number[]>([])
const serieName = ref<string>("")

const chartOptions = ref()
const series = ref<any[]>([])

const processProps = () => {
  serieName.value = props.serieName
  categories.value = props.values.map(x => x.label)
  data.value = props.values.map(x => x.value)

  chartOptions.value = {
        chart: {
          id: "vuechart-example",
        },
        xaxis: {
          categories: categories.value
        },
        plotOptions: {
              bar: {
                borderRadius: 10,
                dataLabels: {
                  position: 'top',
                },
              }
            },
      }

    series.value =  [{
          name: serieName.value,
          data: data.value
    },]
  }

onBeforeMount(() => {
  processProps() 
})

</script>
<template>
  <div class="w-full bg-foreground shadow rounded-2xl p-6">
  <h1 class="text-white font-bold font-xl">{{props.chartTitle}}</h1>
    <apexchart v-show="chartOptions"
      type="bar" :width="CHART_WIDTH" :height="CHART_HEIGHT" :options="chartOptions" :series="series"></apexchart>
  </div>
</template>
