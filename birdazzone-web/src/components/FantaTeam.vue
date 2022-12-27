<script setup lang="ts">
  import type {FantaTeam} from '../api/interfaces/fanta-team';
  import {onBeforeMount, ref} from 'vue';

  let showAll = ref<boolean>(false);
  const props = defineProps<{data: FantaTeam}>();

  onBeforeMount(() => {
    showAll.value = false;
  });
</script>

<template>
  <div
    class="flex flex-col flex-1 my-3 w-full"
    style="flex: 1 1 auto"
  >
    <button class="btn rounded-lg bg-foreground hover:bg-lgray/30">
      <a
        :href="props.data.postUrl"
        target="_blank"
        rel="noopener noreferrer nofollow"
      >
        <div
          class="flex flex-row items-start gap-x-3 p-4"
          style="border-radius: 50%"
        >
          <img
            :src="props.data.profileImageUrl"
            class="align-top rounded-full object-scale-down aspect-square"
            alt="propic"
            onerror="this.onerror = null; this.src='/icons/user.svg' "
          />
          <div class="flex flex-col grow">
            <p class="flex flex-1 text-white font-bold mb-2"
              >@{{ props.data.username }}</p
            >
            <img
              class="fanta-team rounded-lg"
              :src="props.data.imageUrl"
              alt="fantacitorio team"
            />
          </div>
        </div>
      </a>
    </button>
  </div>
</template>
<style>
  .fanta-team {
    filter: hue-rotate(110deg) brightness(85%);
  }
</style>
