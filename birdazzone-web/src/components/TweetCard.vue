<script lang="ts" setup >
import type { FakeTweet } from '@/api/interfaces/fake-tweet';
import { defineProps} from 'vue'

const props = defineProps<{tweet: FakeTweet}>()
</script>

<template>
  <div class="rounded shadow block p-6 m-3">
    <div class="flex justify-between">
      <h2 class="font-bold text-xl">{{ props.tweet.user}}</h2>
      <h2 class=""> {{props.tweet.created_at.toDateString()}}</h2>
    </div>

    <h3 class="my-3">{{ props.tweet.text}}</h3>

    <div class="flex justify-start">
      💙: {{props.tweet.metrics.like_count}}
      📨: {{props.tweet.metrics.retweet_count}}
      📫: {{props.tweet.metrics.reply_count}}
    </div>
  </div>
</template>
