<script lang="ts" setup>
  import Histogram, {type HistogramValue} from './Histogram.vue';
  import ApiRepository from '../api/api-repository';
  import {onBeforeMount, ref} from 'vue';

  const props = defineProps<{
    gameId: string;
    key: number;
    from: string | null;
    to: string | null;
  }>();

  const SERIE_NAME = 'attempts';
  const MAX_COLS = 5
  const LOWER_BOUND = 3

  const histogramValues = ref<HistogramValue[]>([]);
  const error = ref<boolean>(false);
  const errorMessage = ref<string>('');
  const loading = ref<boolean>(false);

  const fetchAttempts = async () => {
    loading.value = true;

    const response = await ApiRepository.getTvGameAttemptsStat(
      props.gameId,
      props.from,
      props.to
    );
    error.value = response.esit;
    if (response.esit) {
      histogramValues.value = response
        .data!.filter(x => response.data!.length < MAX_COLS || x.absoluteFrequency > LOWER_BOUND)
        .map(
          x =>
            ({
              label: x.value,
              value: x.absoluteFrequency
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

  function itaJetLagFrom() {
    return (
      ((Number(props.from!.substring(11, 13)) + 1) as unknown as string) +
      ':' +
      props.from!.substring(14, 16)
    );
  }

  function itaJetLagTo() {
    return (
      ((Number(props.to!.substring(11, 13)) + 1) as unknown as string) +
      ':' +
      props.to!.substring(14, 16)
    );
  }
</script>
<template>
  <div class="h-full">
    <div
      v-if="from && to"
      class="text-md text-white text-semibold mb-3"
    >
      Data refere to the following date-time range: from
      {{ from.substring(0, 10) }}, {{ itaJetLagFrom() }} to
      {{ itaJetLagTo() }}
    </div>
    <Histogram
      v-if="!loading"
      :chart-title="'Played words with more than 3 tentatives'"
      :values="histogramValues"
      :serie-name="SERIE_NAME"
    />
  </div>
</template>
