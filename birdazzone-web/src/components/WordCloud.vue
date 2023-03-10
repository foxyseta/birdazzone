<script lang="ts" setup>
  import ApiRepository from '../api/api-repository';
  import { ref, onBeforeMount } from 'vue';
  import type { WordCloudOptions } from '@/api/interfaces/wordcloud-options';
  import type { ChartEntry } from '@/api/interfaces/chart-entry';
  import type { Solution } from '@/api/interfaces/solution';
  import { SemipolarSpinner } from 'epic-spinners';

  const props = defineProps<{
    tvGameId: string;
    key: number;
    from: string | null;
    to: string | null;
  }>();
  const svgString = ref<string>();
  const loading = ref<boolean>(true);

  const error = ref<boolean>(false);

  const from = ref<string | null>(props.from);
  const to = ref<string | null>(props.to);
  const date = ref<string | null>();

  const explodeChartEntry = (entry: ChartEntry, solution: Solution): string => {
    let res = '';
    const stamp = entry.value === solution.solution ? `${entry.value}🏆` : entry.value;
    for (let i = 0; i < entry.absoluteFrequency; i++) res += stamp + ', ';
    return res;
  };

  const generateWordsList = (entries: ChartEntry[], solution: Solution) =>
    entries.reduce<string>((prev: string, curr: ChartEntry) => prev + explodeChartEntry(curr, solution), '');

  const generateCloudOptions = (entries: ChartEntry[], solution: Solution): WordCloudOptions => ({
    text: generateWordsList(entries, solution),
    useWordList: true,
    fontFamily: 'monospace',
    width: 350,
    height: 250,
    case: 'upper',
    maxNumWords: 40,
    fontScale: 20,
    padding: 1,
    scale: 'sqrt',
    colors: ['#999083', '#aaaaaa', '#ffffff'],
    rotation: 90
  });

  const fetchSolution = async (): Promise<Solution | undefined> => {
    if (!from.value || !to.value) {
      const response = await ApiRepository.getTvGameSolutionById(props.tvGameId);
      if (response.esit) {
        return response.data;
      } else {
        error.value = true;
      }
    } else {
      date.value = from.value.substring(0, 10); // just the relative date
      const response = await ApiRepository.getTvGameSolutionByIdFiltered(props.tvGameId, date.value);
      if (response.esit) {
        return response.data;
      } else {
        error.value = true;
      }
    }
  };

  const fetchStats = async (): Promise<ChartEntry[] | undefined> => {
    if (!from.value || !to.value) {
      const response = await ApiRepository.getTvGameAttemptsStat(props.tvGameId);
      if (response.esit) {
        return response.data;
      } else {
        error.value = true;
      }
    } else {
      const response = await ApiRepository.getTvGameAttemptsStat(props.tvGameId, from.value, to.value);
      if (response.esit) {
        return response.data;
      } else {
        error.value = true;
      }
    }
  };

  const fetchCloud = async (opts: WordCloudOptions) => {
    const response = await ApiRepository.postWordCloudData(opts);
    if (response.esit) {
      svgString.value = response.data;
    } else {
      svgString.value = 'Error';
    }
  };

  onBeforeMount(async () => {
    loading.value = true;
    const stats = await fetchStats();
    const solution = await fetchSolution();
    if (stats && solution) {
      await fetchCloud(generateCloudOptions(stats, solution));
    }
    loading.value = false;
  });
</script>
<template>
  <div class="shadow bg-foreground rounded-xl">
    <div
      v-if="loading"
      class="flex p-5 align-center justify-center"
    >
      <semipolar-spinner
        :animation-duration="2000"
        :size="35"
        color="#1eb980"
      />
    </div>
    <div
      v-show="!loading"
      v-if="svgString"
    >
      <span v-html="svgString"></span>
    </div>
  </div>
</template>
