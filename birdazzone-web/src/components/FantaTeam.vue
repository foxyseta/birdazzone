<script setup lang="ts">
import type { FantaTeam } from '../api/interfaces/fanta-team';
import { onBeforeMount, ref } from 'vue';

let showAll = ref<boolean>(false);
const props = defineProps<{ data: FantaTeam }>();

onBeforeMount(() => {
  showAll.value = false;
});
</script>

<template>
  <div class="flex flex-col flex-1 my-3 w-full" style="flex: 1 1 auto">
    <button class="btn rounded-lg bg-foreground hover:bg-lgray/50 pb-5 p-2">
      <a :href="props.data.postUrl" target="_blank" rel="noopener noreferrer nofollow">
        <div class="flex flex-col" style="border-radius: 50%; flex: 1 1 auto">
          <div class="flex flex-row">
            <img
              :src="props.data.profileImageUrl"
              class="aspect-square p-2"
              style="width: 6rem; border-radius: 50%"
              alt="propic"
              onerror="this.onerror = null; this.src='/icons/user.svg' "
            />
            <p class="flex flex-1 text-white font-bold p-2 items-center">@{{ props.data.username }}</p>

          </div>
          <div class="flex flex-col mx-4" style="flex: 1 1 auto">
            <img :src="props.data.imageUrl" alt="fantacitorio team" style="border-radius: 5%" />
          </div>
        </div>
        </a>
    </button>
  </div>
</template>
