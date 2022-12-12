<script lang="ts" setup>
import Histogram, {type HistogramValue} from './Histogram.vue';
import { onBeforeMount, ref } from 'vue'

const props = defineProps<{gameId: string, from: string|null, to: string|null}>()

const SERIE_NAME = "scores"
const STEP = 300

const histogramValues = ref<HistogramValue[]>([]) 
const error = ref<boolean>(false)
const errorMessage = ref<string>("")
const loading = ref<boolean>(false)

const mocked = [
  {name: "pino", score: 100},
  {name: "pino", score: 300},
  {name: "pino", score: 300},
  {name: "pino", score: 300},
  {name: "pano", score: 1000},
  {name: "peno", score: 1300},
  {name: "peno", score: 1300},
  {name: "peno", score: 1300},
  {name: "peto", score: 2000},
  {name: "peto", score: 2000},
  {name: "peto", score: 2000},
  {name: "peto", score: 2000},
]


const mapData = (scores) => {
  const categories = [0, 0, 0, 0, 0, 0]
  for (const score of scores) {
    const index = Math.trunc(score.score / STEP )
    categories[index < 5 ? index : 5] ++
  }

  histogramValues.value = categories.map((val, index) => ({label: STEP * index, value: val}))
}

// TODO
mapData(mocked)

</script>
<template>
    <Histogram 
      v-if="!loading" :chart-title="'Politician distribution TODO'" :values="histogramValues" :serie-name="SERIE_NAME" />
  
    <!-- Error message -->
    <h1 v-show="error">{{errorMessage}}</h1>
</template>
