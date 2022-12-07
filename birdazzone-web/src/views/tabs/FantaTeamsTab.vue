<script lang="ts" setup>
import FantaTeamsList from '../../components/FantaTeamsList.vue';
import FantaSearchUser from '../../components/FantaSearchUser.vue';
import type  { Tweet } from '../../api/interfaces/tweet';
import { useRoute } from 'vue-router';
import { onBeforeMount, ref } from 'vue';
import ApiRepository from '@/api/api-repository';

const fakeTweets : Tweet[]= [
    {
        author: {
            name: "a",
            profile_image_url: "",
            username: "username"
        },
        created_at: "now",
        text: "this is a tweet",
        metrics: {
            like_count: 0,
            reply_count: 0,
            retweet_count: 0
        },
        medias: ["https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small"]
    },
    {
        author: {
            name: "a",
            profile_image_url: "",
            username: "username"
        },
        created_at: "now",
        text: "this is a tweet",
        metrics: {
            like_count: 0,
            reply_count: 0,
            retweet_count: 0
        },
        medias: ["https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small"]
    },
    {
        author: {
            name: "a",
            profile_image_url: "",
            username: "username"
        },
        created_at: "now",
        text: "this is a tweet",
        metrics: {
            like_count: 0,
            reply_count: 0,
            retweet_count: 0
        },
        medias: ["https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small"]
    },
    {
        author: {
            name: "a",
            profile_image_url: "",
            username: "username"
        },
        created_at: "now",
        text: "this is a tweet",
        metrics: {
            like_count: 0,
            reply_count: 0,
            retweet_count: 0
        },
        medias: ["https://pbs.twimg.com/media/Fi_4v1OWAAMCmxd?format=jpg&name=small"]
    },

]

const username = ref<string>()
const tweets = ref<Tweet[]>([]) 
const isError = ref<boolean>(false)
const error = ref<string>()
const loading = ref<boolean>(false)

const fetchTweets = async () => {
    loading.value = true
    const response = await ApiRepository.getFantacitorioTeams(username.value)
    isError.value = response.esit

    if (response.esit) {
        tweets.value = fakeTweets // TODO response.data
    } else {
        error.value = response.error?.message
    }
    loading.value = false
}

onBeforeMount(()=>fetchTweets())

const onUsernameChanged = async (newUsername: string) => {
    username.value = newUsername
    await fetchTweets()
}

</script>

<template>
    <h1> fanta teams tab</h1>
    <FantaSearchUser @changed="onUsernameChanged" />
    <FantaTeamsList />
</template>