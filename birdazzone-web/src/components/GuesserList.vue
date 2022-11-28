<script setup lang="ts">
  import GuesserListItem from './GuesserListItem.vue'
  import ApiRepository from '../api/api-repository'
  import {onBeforeMount, ref} from 'vue'
  import type { Tweet } from '../api/interfaces/tweet'
  import FilterTandW from './FilterTweetsWordcloud.vue'

  const loading = ref<boolean> (false)
  const error = ref<boolean> (false)
  const list = ref<Tweet[]>([])
  const props = defineProps<{gameId: number}>()

  const from = ref<string>()
  const to = ref<string>()
  
  const fetchList = async () => {
    if(!from.value || !to.value){
      const resp = await ApiRepository.getListOfGuesser(props.gameId.toString())
      if (resp.esit) {
        list.value = resp.data!.entries
      } else {
        error.value = true
      }
    }
    else{
      const resp = await ApiRepository.getListOfGuesserFiltered(props.gameId.toString(),from.value.toString(), to.value.toString())
      if (resp.esit) {
        list.value = resp.data!.entries
      } else {
        error.value = true
      }
    }
  }
  onBeforeMount(fetchList)
</script>

<template>
  <div v-show="loading" class="h-screen flex items-center">
    <semipolar-spinner :animation-duration="2000" :size="50" color="#1eb980" />
  </div>
  <div v-show="!loading">
    <FilterTandW :from="from" :to="to" @change-from-to="(f, t) => {from = f; to = t; fetchList()}" />
    <div v-if="list.length === 0" class="flex flex-1 my-3 p-10 text-white font-semibold" style="flex: 1 1 auto">
      Nobody was smart enough to guess &#127773;
    </div> 
    <div v-else class="flex flex-col">
      <div class="flex flex-col" v-for="(item,index) in list" :key="index">
        <GuesserListItem :data="item" :index="index" class="flex flex-1"/>
      </div>
    </div>
  </div>
</template>
