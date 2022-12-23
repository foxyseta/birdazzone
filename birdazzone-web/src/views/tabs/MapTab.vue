<script lang="ts" setup>
import { ref } from 'vue';
import FilterMap from '../../components/FilterList.vue';
import GuesserMap from '../../components/GuesserMap.vue';

const props = defineProps<{ gameId: string }>();

const from = ref<string | null>(null);
const to = ref<string | null>(null);
const mapKey = ref<number>(0);

const forceUpdate = () => {
  mapKey.value += 1;
};
const a = (f: string | null, t: string | null) => {
  from.value = f;
  to.value = t;
  forceUpdate();
};
</script>

<template>
  <div class="flex flex-col pl-20 justify-items-start mt-6 mb-4">
    <FilterMap @change-from-to="a" />
  </div>

  <div class="flex pl-20">
    <div class="justify-center">
      <GuesserMap :game-id="props.gameId" :key="mapKey" :from="from" :to="to" />
    </div>
  </div>
</template>
