<script setup lang="ts">
    import guesserListItem from '@/components/guesserListItem.vue'
    import ApiRepository from '@/api/api-repository';
    import type { ListGuesser } from '@/api/interfaces/list-guesser';
    import {onBeforeMount, ref} from 'vue'

    const error = ref<boolean> (false)
    const list = ref<ListGuesser[]>([])
    const props = defineProps<{gameId: number}>()

    const fetchList = async () => {
    const resp = await ApiRepository.getListGuesser(props.gameId.toString())
    if (resp.esit) {
      list.value = resp.data!
    } else {
      error.value = true
    }
  }

    onBeforeMount(fetchList)
</script>

<template>
    <div v-for="item in list" :key="item.username">
        <guesserListItem :data="item" class="flex flex-1"/>
    </div>
</template>
