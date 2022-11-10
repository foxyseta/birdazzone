<script setup lang="ts">
  import {ref} from 'vue'
  let showAll = ref<boolean>(false)
  function prova(){
    showAll.value = !showAll.value
  }
</script>

<template>
    <div class="flex flex-col flex-1 my-3 ml-10" style="flex: 1 1 auto">
      <button class="btn p-4 rounded-lg  bg-foreground hover:bg-lgray/50" @click="prova()" >
        <div class="flex flex-row" style="flex: 1 1 auto;">
            <div class="flex" style="height:4rem">
                <img :src="'./public/defaultPropic.svg'" alt="propic"/>
            </div>
            <div class="flex flex-col mx-4" style="flex: 1 1 auto">
                <p class="flex flex-1 text-white font-bold">@username</p>
                <p class="flex flex-1 text-lgray">name&surname</p>
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
            <div id="paramTweet" class="flex flex-1 flex-row text-lgray">
                prova param
            </div>
        </div>
      </button>
    </div>
</template>