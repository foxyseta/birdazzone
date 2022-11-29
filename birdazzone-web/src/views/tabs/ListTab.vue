<script lang="ts" setup>
import { ref } from 'vue';
import GuesserList from '../../components/GuesserList.vue';
import WordCloud from '../../components/WordCloud.vue';
import FilterList from '../../components/FilterList.vue';

const props = defineProps<{gameId: string}>()

const from = ref<string | null>(null)
const to = ref<string | null>(null)

</script>
<template>
    <div class="flex flex-col m-6 ml-20">
        <FilterList :from="from" :to="to" @change-from-to="(f, t) => {from = f; to = t;}" />
    </div>
    <div class="flex justify-evenly">
        <div>
            <GuesserList :game-id="props.gameId" :from="from" :to="to" />
        </div>
        <div>
            <WordCloud :tv-game-id="props.gameId" :from="from" :to="to" />
        </div>
    </div>
</template>