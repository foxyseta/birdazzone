<script lang="ts" setup>
import AerogramCard from '../../components/AerogramCard.vue';

const props = defineProps<{gameId: string}>()
</script>

<template>
    <div class="flex pl-20">
        <AerogramCard :id="props.gameId"/>
    </div>
</template>