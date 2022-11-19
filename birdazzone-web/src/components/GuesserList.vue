<script setup lang="ts">
  import GuesserListItem from '@/components/GuesserListItem.vue'
  import ApiRepository from '@/api/api-repository';
  import {onBeforeMount, ref} from 'vue'
  import type { Tweet } from '@/api/interfaces/tweet';

  const error = ref<boolean> (false)
  const list = ref<Tweet[]>([])
  const props = defineProps<{gameId: number}>()

  const fetchList = async () => {
  const resp = await ApiRepository.getListOfGuesser(props.gameId.toString())
    if (resp.esit) {
      list.value = resp.data!.entries
    } else {
      error.value = true
    }
  }
  onBeforeMount(fetchList)
</script>

<template>
  <div>
    <div v-if="list.length === 0" class="flex flex-1 my-3 p-10 text-white font-semibold" style="flex: 1 1 auto">
      Nobody was smart enough to guess &#127773;
    </div>  
    <div v-else class="flex flex-col" v-for="(item,index) in list" :key="index">
        <GuesserListItem :data="item" :index="index" class="flex flex-1"/>
    </div>
  </div>
</template>
