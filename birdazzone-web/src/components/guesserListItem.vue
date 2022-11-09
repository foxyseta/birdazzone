<script setup lang="ts">
  import {ref} from 'vue'

  interface listItem {
    username: string
    fullName: string
    propic: string
    time: string
  }
  const listItems = ref<listItem[]>([
    {
      username: "@paoooool",
      fullName: "Paolo Ceroni",
      propic: "./public/defaultPropic.svg",
      time: "12:03" 
    },
    {
      username: "@fede",
      fullName: "Fede Grì",
      propic: "./public/defaultPropic.svg",
      time: "12:45" 
    }
  ])
</script>

<template>
    <div class="flex flex-col flex-1 my-3 ml-10" v-for="(item, index) in listItems" style="flex: 1 1 auto;">
        <div class="flex flex-row bg-foreground rounded-lg p-4" style="flex: 1 1 auto;">
            <div class="flex" style="height:4rem">
                <img :src="'./public/defaultPropic.svg'" alt="propic"/>
            </div>
            <div class="flex flex-col mx-4" style="flex: 1 1 auto">
                <div class="text-white">{{item.username}}</div>
                <div class="text-lgray">{{item.fullName}}</div>
            </div>
            <div class="flex flex-col mx-4">
            <div class="flex" style="flex: 1 1 auto; height:2rem">
                <img :src="'./public/defaultPropic.svg'" alt="propic"/>
                </div>
                <div class="text-lgray text-xs">{{item.time}}</div>
            </div>
        </div>
    </div>
</template>