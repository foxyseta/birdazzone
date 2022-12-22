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
  <div class="place-content-center w-full mb-10 h-screen p-4">
    <div class="flex justify-center">
      <div class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold p-3 m-3">CHESS</div>
    </div>
    <NewChessGameAlert v-if="showAlert" @done="onValuesInserted" />
    <Chessboard
      v-if="selectedUser && selectedColor && !showAlert"
      :starting-color="selectedColor"
      :user="selectedUser"
    />
  </div>
</template>
