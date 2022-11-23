<script setup lang="ts">
  import GuesserListItem from '@/components/GuesserListItem.vue'
  import ApiRepository from '@/api/api-repository';
  import {onBeforeMount, ref} from 'vue'
  import type { Tweet } from '@/api/interfaces/tweet';
  import { SemipolarSpinner } from 'epic-spinners';
  import PaginationBar from '../components/PaginationBar.vue';

  const loading = ref<boolean> (false)
  const list = ref<Tweet[]>([])
  const props = defineProps<{gameId: number}>()
  const max = ref<number>(0)
  const actualPage = ref<number>(5)
  const itemPerPage = ref<number>(10)

  const fetchList = async () => {
    loading.value = true
    const resp = await ApiRepository.getListOfGuesser(props.gameId.toString(),actualPage.value.toString(), itemPerPage.value.toString())
      if (resp.esit) {
        list.value = resp.data!.entries
        max.value = resp.data!.numberOfPages
        console.log("resp")
        console.log(resp)
        loading.value = false
        console.log(resp.data!.numberOfPages)
      } 
  }

  onBeforeMount(fetchList)
</script>

<template>
  <div class="flex flex-col">
    <PaginationBar v-show="!loading" :actualPage="actualPage" :max="max"/>
    <div v-show="loading" class="h-screen flex items-center">
      <semipolar-spinner :animation-duration="2000" :size="50" color="#1eb980" />
    </div>
    <div v-show="!loading">
      <div v-if="list.length === 0" class="flex flex-1 my-3 text-white font-semibold" style="flex: 1 1 auto">
        Nobody was smart enough to guess &#127773;
      </div>  
      <div v-else class="flex flex-col" v-for="(item,index) in list" :key="index">
          <GuesserListItem :data="item" :index="index+(actualPage-1)*itemPerPage +1" class="flex flex-1"/>
      </div>
    </div>
  </div>
</template>