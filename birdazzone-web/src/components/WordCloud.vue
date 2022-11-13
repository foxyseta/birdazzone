<script lang="ts" setup>
import ApiRepository from '../api/api-repository'
import { ref, onBeforeMount } from 'vue'
import type { WordCloudOptions } from '@/api/interfaces/wordcloud-options';
import type { ChartEntry } from '@/api/interfaces/chart-entry';
import type { Solution } from '@/api/interfaces/solution';
import { SemipolarSpinner } from 'epic-spinners';

const props = defineProps<{tvGameId: number}>()
const svgString = ref<string>()
const loading = ref<boolean>(true)

const explodeChartEntry = (entry: ChartEntry, solution: Solution): string => {
  let res = ""
  const stamp = (entry.value === solution.solution)? `${entry.value}🏆`: entry.value
  for(let i = 0; i < entry.absoluteFrequency; i++) 
    res += stamp + ", "
  return res
}

const generateWordsList = (entries: ChartEntry[], solution: Solution) => 
    entries.reduce<string>((prev: string, curr: ChartEntry) => prev + explodeChartEntry(curr, solution), "")

const generateCloudOptions = (entries: ChartEntry[], solution: Solution): WordCloudOptions => ({
  text: generateWordsList(entries, solution),
  useWordList: true,
  fontFamily: 'monospace',
  width: 350,
  height: 250,
  case: 'upper',
  maxNumWords: 40,
  fontScale: 40,
  padding: 1,
  scale: 'sqrt',
  colors: ["#999083","#53534f"],
  rotation: 90
})

const fetchSolution = async (): Promise<Solution|undefined> => 
  (await ApiRepository.getTvGameSolutionById(props.tvGameId.toString())).data

const fetchStats = async (): Promise<ChartEntry[]|undefined> => 
  (await ApiRepository.getTvGameAttemptsStat(props.tvGameId.toString())).data

const fetchCloud = async (opts: WordCloudOptions) => {
  const response = await ApiRepository.postWordCloudData(opts)
  if (response.esit) {
    svgString.value = response.data
  } else {
    svgString.value = "Error"
  }
}

onBeforeMount(async () => {
  loading.value = true
  const stats = await fetchStats()
  const solution = await fetchSolution()
  if (stats && solution) {
    await fetchCloud(generateCloudOptions(stats, solution))
  }
  loading.value = false 
})

</script>
<template>
  <div  class="shadow bg-foreground rounded-xl m-5">
  <div v-if="loading" class="flex p-5 align-center justify-center">
      <semipolar-spinner :animation-duration="2000" :size="35" color="#1eb980" />
  </div>
    <div v-show="!loading" v-if="svgString">
      <span v-html="svgString"></span>
    </div>
  </div>
</template>
