<script setup lang="ts">
  import {ref} from 'vue'
  let showAll = ref<boolean>(false)
  function prova(){
    console.log(showAll)
    showAll.value = !showAll.value
    console.log(showAll)
  }
  function isShowAll(){
    return showAll
  }
  interface listItem {
    username: string
    fullName: string
    propic: string
    time: string
    body: string
  }
  const listItems = ref<listItem[]>([
    {
      username: "@paoooool",
      fullName: "Paolo Ceroni",
      propic: "./public/defaultPropic.svg",
      time: "12:03",
      body:"prova lungaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    },
    {
      username: "@marco",
      fullName: "Fede Grì",
      propic: "./public/defaultPropic.svg",
      time: "12:05",
      body:"prova corta"
    }
  ])
</script>

<template>
    <div class="flex flex-col flex-1 my-3 ml-10" v-for="(item, index) in listItems" style="width:50%">
      <button class="btn p-4 rounded-lg  bg-foreground hover:bg-lgray/50" @click="prova()" >
        <div class="flex flex-row" style="flex: 1 1 auto;">
            <div class="flex" style="height:4rem">
                <img :src="'./public/defaultPropic.svg'" alt="propic"/>
            </div>
            <div class="flex flex-col mx-4" style="flex: 1 1 auto">
                <p class="flex flex-1 text-white font-bold">{{item.username}}</p>
                <p class="flex flex-1 text-lgray">{{item.fullName}}</p>
            </div>
            <div class="flex flex-col mx-4">
            <div class="flex" style="flex: 1 1 auto; height:2rem">
                <img :src="'./public/defaultPropic.svg'" alt="propic"/>
            </div>
                <div class="text-lgray text-xs">{{item.time}}</div>
            </div>
        </div>
        <div id="wholeTweet" class="flex flex-1 flex-col mt-3 p-2 border border-lgray rounded-lg" v-show="showAll">
            <div id="bodyTweet" class="flex flex-1 text-white">
                "{{item.body}}"
            </div>
            <div id="paramTweet" class="flex flex-1 flex-row text-lgray">
                prova param
            </div>
        </div>
      </button>
    </div>
</template>