<script setup lang="ts">
    import guesserListItem from '@/components/guesserListItem.vue'
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

    <div class="flex flex-col" v-for="(item,index) in list" :key="index">
        <guesserListItem :data="item" :index="index" class="flex flex-1"/>
    </div>
  </div>
</template>
