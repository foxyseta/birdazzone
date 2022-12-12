<script lang="ts" setup>
import { ref } from 'vue';
import FilterA from '../../components/FilterAerogram.vue';
import AerogramCard from '../../components/AerogramCard.vue';

const props = defineProps<{ gameId: string }>();

const from = ref<string | null>(null);
const to = ref<string | null>(null);
const aerogramKey = ref<number>(0);

const forceUpdate = () => {
  aerogramKey.value += 1;
};
const a = (f: string | null, t: string | null) => {
  from.value = f;
  to.value = t;
  forceUpdate();
};
</script>

<template>
  <div class="flex flex-col pl-20 justify-items-start mt-6 mb-4">
    <FilterA @change-from-to="a" />
  </div>
  <div class="flex pl-20">
    <div class="justify-items-start">
      <AerogramCard :id="props.gameId" :key="aerogramKey" :from="from" :to="to" />
    </div>
  </div>
</template>
