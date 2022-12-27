<script lang="ts" setup>
  import type { Politician } from '@/api/interfaces/politician';
  import Histogram, { type HistogramValue } from './Histogram.vue';

  const props = defineProps<{ list: Politician[] }>();
  const chartTitle = 'Politicians Ranking';
  const serieName = 'Politician';

  const politicianToValue = (p: Politician): HistogramValue => ({
    label: p.name,
    value: p.score
  });
</script>
<template>
  <Histogram
    :serie-name="serieName"
    :chart-title="chartTitle"
    :values="props.list.map(politicianToValue)"
  />
</template>
