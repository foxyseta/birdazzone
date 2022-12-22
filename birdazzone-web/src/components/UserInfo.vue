<script lang="ts" setup>
import type { User } from '@/api/interfaces/tweet';

const props = defineProps<{ user: User; gameId: string | null; turn: number }>();
</script>
<template>
  <div class="flex justify-between items-center bg-foreground rounded-2xl p-6">
    <div class="flex items-center mr-5">
      <img :src="props.user.profileImageUrl" alt="propic" class="rounded-full w-20 h-20" />
      <div class="m-4">
        <h1 class="text-white font-bold text-xl">@{{ props.user.username }}</h1>
        <h1 class="text-lgray font-bold text-xl">{{ props.user.name }}</h1>
      </div>
    </div>
    <div class="flex flex-col justify-center mx-10">
      <h1 class="text-lgreen text-center">Turn</h1>
      <h1 class="text-lgreen font-bold text-4xl text-center">{{ props.turn }}</h1>
    </div>
    <div class="flex flex-col justify-center ml-5">
      <h1 class="text-white text-center">Game id</h1>
      <h1 class="text-white text-4xl font-bold text-center">{{ props.gameId }}</h1>
    </div>
  </div>
</template>
