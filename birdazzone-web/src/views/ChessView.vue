<script setup lang="ts">
import BirdazzoneButton from '../components/BirdazzoneButton.vue';
import { ref } from 'vue';
import { TheChessboard } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';
import type { ChessboardAPI, BoardConfig } from 'vue3-chessboard';
import { createGunzip } from 'zlib';

const boardAPI = ref<ChessboardAPI>();

const boardConfig: BoardConfig = {
  events: {
    move: (_) => {
      console.log("!");
      window.open(
        "https://twitter.com/intent/tweet?url=https%3A%2F%2Ffen2image.chessvision.ai%2F" +
        encodeURIComponent(
          encodeURIComponent(boardAPI?.value?.board.getFen() ?? 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1')
        ) +
        "&text=Your%20move.%20Please%20use%20the%20%22a1b2%22%20format.%20This%20is%20a%20majority%20vote.",
        "_blank");
    },
  },
};

function handleCheckmate(isMated: string) {
  alert((isMated === 'w' ? 'Black' : 'White') + 'wins!');
}

</script>
<template>
  <div class="flex flex-col place-content-center flex-initial justify-start place-items-center w-full mb-10 h-screen">
    <div class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold p-3 my-12 m-3">CHESS</div>
    <div class="flex flex-row place-content-center flex-initial justify-start place-items-center mb-10 h-screen">
      <TheChessboard class="p-3" :board-config="boardConfig" @board-created="(api) => (boardAPI = api)"
        @checkmate="handleCheckmate" />
      <div class="p-3">
        <BirdazzoneButton :text="'Check for a new move'" :active="true" />
      </div>
    </div>
  </div>
</template>
