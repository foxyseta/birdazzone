<script setup lang="ts">
import type { User } from '@/api/interfaces/tweet';
import Chessboard from '@/components/Chessboard.vue'
import NewChessGameAlert from '@/components/NewChessGameAlert.vue'
import { ref } from 'vue'

const showAlert = ref<boolean>(true)
const selectedUser = ref<User>()

const onUsernameInserted = (user: User) => {
  showAlert.value = false
  selectedUser.value = user
}

</script>
<template>
  <div class="flex flex-col place-content-center flex-initial justify-start place-items-center w-full mb-10 w-screen h-screen">
    <div class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold p-3 my-12 m-3">CHESS</div>
    <NewChessGameAlert v-if="showAlert" @done="onUsernameInserted" />
    <Chessboard v-if="selectedUser && !showAlert" :user="selectedUser" />
  </div>
</template>
