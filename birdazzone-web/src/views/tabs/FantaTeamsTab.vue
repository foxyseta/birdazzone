<script lang="ts" setup>
import FantaTeamsList from '../../components/FantaTeamsList.vue';
import FantaSearchUser from '../../components/FantaSearchUser.vue';
import type { Tweet } from '../../api/interfaces/tweet';
import { useRoute } from 'vue-router';
import { onBeforeMount, ref } from 'vue';
import ApiRepository from '@/api/api-repository';
import { SemipolarSpinner } from 'epic-spinners';

const fakeTweets: Tweet[] = [
  {
    author: {
      name: 'mario rossi',
      profile_image_url: '',
      username: 'mariorossi',
    },
    created_at: 'now',
    text: 'this is a tweet',
    metrics: {
      like_count: 0,
      reply_count: 0,
      retweet_count: 0,
    },
    medias: ['https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small'],
  },
  {
    author: {
      name: 'luigi verdi',
      profile_image_url: '',
      username: 'gigiverdi',
    },
    created_at: 'now',
    text: 'this is a tweet',
    metrics: {
      like_count: 0,
      reply_count: 0,
      retweet_count: 0,
    },
    medias: ['https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small'],
  },
  {
    author: {
      name: 'mattia girolimetto',
      profile_image_url: '',
      username: 'mattiagiro',
    },
    created_at: 'now',
    text: 'this is a tweet',
    metrics: {
      like_count: 0,
      reply_count: 0,
      retweet_count: 0,
    },
    medias: ['https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small'],
  },
  {
    author: {
      name: 'federica grisendi',
      profile_image_url: '',
      username: 'fedegri',
    },
    created_at: 'now',
    text: 'this is a tweet',
    metrics: {
      like_count: 0,
      reply_count: 0,
      retweet_count: 0,
    },
    medias: ['https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small'],
  },
  {
    author: {
      name: 'a',
      profile_image_url: '',
      username: 'chilegge',
    },
    created_at: 'now',
    text: 'this is a tweet',
    metrics: {
      like_count: 0,
      reply_count: 0,
      retweet_count: 0,
    },
    medias: ['https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small'],
  },
];

const username = ref<string>();
const tweets = ref<Tweet[]>([]);
const isError = ref<boolean>(false);
const error = ref<string>();
const loading = ref<boolean>(false);

const fetchTweets = async () => {
  const mock = username.value
    ? fakeTweets.filter((x: Tweet) => x.author.username.startsWith(username.value!))
    : fakeTweets;
  const response = { esit: true, data: mock, error: { message: '' } }; //await ApiRepository.getFantacitorioTeams(username.value)
  isError.value = response.esit;
  if (response.esit) {
    tweets.value = response.data;
  } else {
    error.value = response.error?.message;
  }
};

onBeforeMount(async () => {
  loading.value = true;
  await fetchTweets();
  loading.value = false;
});

const onUsernameChanged = async (newUsername: string) => {
  loading.value = true;
  username.value = newUsername;
  await fetchTweets();
  loading.value = false;
};
</script>

<template>
  <div class="h-screen">
    <FantaSearchUser class="m-10" @changed="onUsernameChanged" />
    <div v-show="loading" class="h-screen flex justify-center items-center">
      <semipolar-spinner :animation-duration="2000" :size="70" color="#1eb980" />
    </div>
    <FantaTeamsList v-show="!loading" :tweets="tweets" />
  </div>
</template>
