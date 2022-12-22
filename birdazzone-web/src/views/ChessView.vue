<script setup lang="ts">
import type { User } from '@/api/interfaces/tweet';
import Chessboard, { type ChessColor } from '@/components/Chessboard.vue';
import NewChessGameAlert from '@/components/NewChessGameAlert.vue';
import { ref } from 'vue';

const showAlert = ref<boolean>(true);
const selectedUser = ref<User>();
const selectedColor = ref<ChessColor>();

const onValuesInserted = (user: User, color: ChessColor) => {
  showAlert.value = false;
  selectedUser.value = user;
  selectedColor.value = color;
};
</script>
<template>
  <div class="place-content-center place-items-center w-full mb-10 h-screen p-10">
    <div class="flex justify-center">
      <div class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold py-3 my-8 m-3 mx-0" style="flex: 1 1 auto; text-align: center; max-width: 10rem">CHESS</div>
    </div>
    <NewChessGameAlert v-if="showAlert" @done="onValuesInserted" />
    <Chessboard
      v-if="selectedUser && selectedColor && !showAlert"
      :starting-color="selectedColor"
      :user="selectedUser"
    />
  </div>
</template>
