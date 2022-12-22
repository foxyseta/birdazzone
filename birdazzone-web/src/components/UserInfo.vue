<script lang="ts" setup>
import type { User } from '@/api/interfaces/tweet';

const props = defineProps<{user: User, gameId: string|undefined, turn: number}>()
</script>
<template>
    <div class="flex justify-between items-center bg-foreground rounded-2xl p-6">
        <div class="flex items-center">
            <img :src="props.user.profileImageUrl" alt="propic" class="rounded-full w-20 h-20" />
            <div class="m-4">
                <h1 class="text-white font-bold text-xl" >@{{ props.user.username}}</h1>
                <h1 class="text-lgray font-bold text-xl" >{{ props.user.name }}</h1>
            </div>
        </div>
        <h1 class="text-lgreen font-bold"> Turn: {{ props.turn}}</h1>
        <h1 class="text-white font-bold"> Game id: {{ props.gameId }}</h1>
    </div>
</template>