<script setup lang="ts">
import type { Tweet } from '@/api/interfaces/tweet';
import { onBeforeMount, ref } from 'vue';

let showAll = ref<boolean>(false);
function changeVisibility() {
  showAll.value = !showAll.value;
}
const props = defineProps<{ data: Tweet; index: number }>();

const hours: number = new Date(props.data.created_at).getHours();
const mins: number = new Date(props.data.created_at).getMinutes();

onBeforeMount(() => {
  showAll.value = false;
});
</script>

<template>
  <div class="flex flex-col flex-1 my-3 w-full" style="flex: 1 1 auto">
    <div class="flex flex-row flex-1 w-full">
      <span class="mr-3 text-lgray">{{ index }}.</span>
      <button class="btn rounded-lg bg-foreground hover:bg-lgray/50 p-4 flex-1" @click="changeVisibility()">
        <div class="flex flex-row" style="flex: 1 1 auto">
          <div class="flex" style="width: 4rem; height: 4rem">
            <img
              :src="props.data.author.profileImageUrl"
              class="aspect-square rounded-lg"
              style="width: 10em; border-radius: 50%"
              alt="propic"
              onerror="this.onerror = null; this.src='/icons/user.svg' "
            />
          </div>
          <div class="flex flex-col mx-4" style="flex: 1 1 auto">
            <p class="flex flex-1 text-white font-bold">@{{ props.data.author.username }}</p>
            <p class="flex flex-1 text-lgray">{{ props.data.author.name }}</p>
          </div>
          <div class="flex flex-col mx-4">
            <div class="flex" style="flex: 1 1 auto; height: 2rem">
              <img v-if="index === 1" :src="'/icons/coccarda1.svg'" alt="medal1" />
              <img v-if="index === 2" :src="'/icons/coccarda2.svg'" alt="medal2" />
              <img v-if="index === 3" :src="'/icons/coccarda3.svg'" alt="medal3" />
            </div>
            <div class="text-lgray text-xs">{{ hours }}:{{ mins > 10 ? mins : '0' + mins }}</div>
          </div>
        </div>
        <div
          id="wholeTweet"
          class="flex flex-1 flex-col mt-3 p-2 border border-2 border-lgray rounded-lg"
          v-show="showAll"
        >
          <div id="bodyTweet" class="flex flex-1 text-white">
            {{ props.data.text }}
          </div>
          <div id="paramTweet" class="flex flex-row mt-1 text-lgray" style="justify-content: left; max-height: 1.5rem">
            <div class="flex flex-row mr-2">
              <img class="mr-1" alt="replies" style="max-height: 1.5rem" src="/icons/comment.svg" />
              {{ props.data.metrics.reply_count }}
            </div>
            <div class="flex flex-row mr-2">
              <img class="mr-1" alt="retweets" style="max-height: 1.5rem" src="/icons/retweet.svg" />
              {{ props.data.metrics.retweet_count }}
            </div>
            <div class="flex flex-row">
              <img class="mr-1" alt="likes" style="max-height: 1.5rem" src="/icons/heart.svg" />
              {{ props.data.metrics.like_count }}
            </div>
          </div>
        </div>
      </button>
    </div>
  </div>
</template>
