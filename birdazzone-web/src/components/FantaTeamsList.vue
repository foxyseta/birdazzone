<script lang="ts" setup>
import type { FantaTeam } from '../api/interfaces/fanta-team';

const props = defineProps<{ teams: FantaTeam[]}>();

</script>

<template>
  <div v-if="teams.length == 1" class="flex justify-center">
    <div v-for="(item, index) in teams" :key="index">
      <img :src="item.imageUrl" class="fantaTeam" :alt="'@' + item.username + ' team'" width="800" />
    </div>
  </div>
  <div v-else>
    <div v-if="teams.length % 2 == 0" class="mx-10 md:mx-20 grid gap-4 md:gap-10 grid-cols-1 md:grid-cols-2 mb-10">
      <div v-for="(item, index) in teams" :key="index">
        <img :src="item.imageUrl" class="fantaTeam" :alt="'@' + item.username + ' team'" />
      </div>
    </div>
    <div v-else class="mx-10 md:mx-20 grid gap-4 lg:gap-10 grid-cols-1 lg:grid-cols-3 mb-10">
      <div v-for="(item, index) in teams" :key="index">
        <img :src="item.imageUrl" class="fantaTeam" :alt="'@' + item.username + ' team'" />
      </div>
    </div>
  </div>
</template>

<style>
.fantaTeam {
  filter: hue-rotate(110deg) brightness(85%);
}
</style>
