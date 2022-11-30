<script lang="ts" setup>
import GuesserList from '@/components/GuesserList.vue';
import WordCloud from '@/components/WordCloud.vue';

const props = defineProps<{gameId: string}>()
</script>
<template>
    <div class="flex justify-evenly">
        <div>
            <GuesserList :game-id="props.gameId" />
        </div>
        <div>
            <WordCloud :tv-game-id="props.gameId"/>
        </div>
    </div>
</template>
