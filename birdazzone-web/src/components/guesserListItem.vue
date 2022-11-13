<script setup lang="ts">
  import type { Tweet } from '@/api/interfaces/tweet';
import {ref} from 'vue'

  let showAll = ref<boolean>(false)
  function prova(){
    showAll.value = !showAll.value
  }

  const data = defineProps<{data:Tweet}>()
</script>

<template>
    <div class="flex flex-col flex-1 my-3 ml-10" style="flex: 1 1 auto">
      <button class="btn p-4 rounded-lg  bg-foreground hover:bg-lgray/50" @click="prova()" >
        <div class="flex flex-row" style="flex: 1 1 auto;">
            <div class="flex" style="height:4rem">
                <img :src="'./public/icons/user.svg'" alt="propic"/>
            </div>
            <div class="flex flex-col mx-4" style="flex: 1 1 auto">
                <p class="flex flex-1 text-white font-bold">{{data.author.username}}</p>
                <p class="flex flex-1 text-lgray">{{data.author.name}}</p>
            </div>
            <div class="flex flex-col mx-4">
            <div class="flex" style="flex: 1 1 auto; height:2rem">
                <img :src="'./public/defaultPropic.svg'" alt="propic"/>
            </div>
                <div class="text-lgray text-xs">17:15</div>
            </div>
        </div>
        <div id="wholeTweet" class="flex flex-1 flex-col mt-3 p-2 border border-lgray rounded-lg" v-show="showAll">
            <div id="bodyTweet" class="flex flex-1 text-white">
                "body"
            </div>
            <div id="paramTweet" class="flex flex-row text-lgray" style="justify-content:left;max-height: 1.5rem;">
              <div class="flex flex-row mr-2"><img class="mr-1" style="max-height: 1.5rem;" src='/public/icons/comment.svg' /> 23 </div>  
              <div class="flex flex-row mr-2"><img class="mr-1" style="max-height: 1.5rem;" src='/public/icons/retweet.svg' /> 12 </div>
              <div class="flex flex-row"><img class="mr-1" style="max-height: 1.5rem;" src='/public/icons/heart.svg' /> 56 </div>
            </div>
        </div>
      </button>
    </div>
</template>
