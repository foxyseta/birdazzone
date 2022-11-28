<script lang="ts" setup>
import GuesserMap from '@/components/GuesserMap.vue';


const props = defineProps<{gameId: string}>()
</script>
<template>
    <div>
        <GuesserMap :game-id="props.gameId" />
    </div>
</template>