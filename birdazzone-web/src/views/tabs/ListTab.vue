<script lang="ts" setup>
  import { ref } from 'vue';
  import GuesserList from '../../components/GuesserList.vue';
  import WordCloud from '../../components/WordCloud.vue';
  import FilterList from '../../components/FilterList.vue';

  const props = defineProps<{ gameId: string }>();

  const from = ref<string | null>(null);
  const to = ref<string | null>(null);
  const listKey = ref<number>(0);
  const wordcloudKey = ref<number>(0);

  const forceUpdate = () => {
    listKey.value += 1;
    wordcloudKey.value += 1;
  };
  const a = (f: string | null, t: string | null) => {
    from.value = f;
    to.value = t;
    forceUpdate();
  };
</script>
<template>
  <div class="flex flex-col m-6 ml-20">
    <FilterList @change-from-to="a" />
  </div>
  <div class="flex justify-evenly">
    <div>
      <GuesserList
        :game-id="props.gameId"
        :key="listKey"
        :from="from"
        :to="to"
      />
    </div>
    <div>
      <WordCloud
        :tv-game-id="props.gameId"
        :key="wordcloudKey"
        :from="from"
        :to="to"
      />
    </div>
  </div>
</template>
