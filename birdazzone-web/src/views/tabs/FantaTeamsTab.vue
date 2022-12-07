<script lang="ts" setup>
import FantaTeamsList from '../../components/FantaTeamsList.vue';
import FantaSearchUser from '../../components/FantaSearchUser.vue';
import type  { Tweet } from '../../api/interfaces/tweet';

const tweets : Tweet[]= [
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

</script>

<template>
    <h1> fanta teams tab</h1>
    <FantaSearchUser />
    <FantaTeamsList />
</template>