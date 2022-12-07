<script lang="ts" setup>
import ApiRepository from '../api/api-repository'
import type { Tweet } from '../api/interfaces/tweet';
import ref from 'vue'

const props = defineProps<{ tweets: Tweet[] }>()
</script>
<template>

  <div class="flex flex-col justify-center align-center" style="flex:1 1 auto; width:13rem ">
    <div v-for="(item, index) in tweets" :key="index">
      <div class="font-semibold text-white ml-2">
        <img :src="item.medias[0]" alt="fantacitorioTeam" />
      
      </div>
    </div>
  </div>

</template>