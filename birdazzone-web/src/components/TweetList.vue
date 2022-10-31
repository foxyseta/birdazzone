<script lang="ts" setup>
import ApiRepository from '../api/api-repository'
import { onBeforeMount, ref } from 'vue'
import type { FakeTweet } from '@/api/interfaces/fake-tweet';
import TweetCard from './TweetCard.vue';

/* FAKE DATA */
const fake : FakeTweet[]= [
  {
  user: "mario_rossi",
  text: "I like this game!!",
  metrics: {
    like_count: 123,
    retweet_count: 2300
  },
  created_at: new Date() 
},
  {
  user: "luigi_verdi",
  text: "Oh yeah!!",
  metrics: {
    like_count: 233,
    retweet_count:232 
  },
  created_at: new Date() 
}
]

/* END FAKE DATA */
const isError = ref<boolean>(false)
const errorMessage = ref<string>("")
const tweets = ref<FakeTweet[]>([])

const fetchList = async () => {
  //const response = await ApiRepository.getTweetList("")
  const response = {esit: true, error: "lol"} // TODO remove mock
  if (response.esit) {
    tweets.value.push(...fake)
  } else {
    isError.value = true
    errorMessage.value = response.error
  }
}

onBeforeMount(fetchList)

</script>
<template>
  <!-- ERROR -->
  <h1 v-if="isError">{{errorMessage}}</h1>
  <div v-else>
    <div v-for="t in tweets"> 
      <TweetCard :tweet=t />
    </div>
  </div>

</template>
