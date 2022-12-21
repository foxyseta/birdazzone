<script setup lang="ts">
import BirdazzoneButton from '../components/BirdazzoneButton.vue';
import { ref } from 'vue';
import { TheChessboard } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';
import type { ChessboardAPI, BoardConfig } from 'vue3-chessboard';

const boardAPI = ref<ChessboardAPI>();
const boardConfig: BoardConfig = {
  coordinates: false,
  autoCastle: false,
  events: {
    move: updateIntent
  }
};
const intent = ref("");

function updateIntent() {
  intent.value =
      "https://twitter.com/intent/tweet?url=https%3A%2F%2Ffen2image.chessvision.ai%2F" +
      encodeURIComponent(/*boardAPI?.value?.board.getFen() ??*/ "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1") +
      "&text=Your%20move.%20Please%20use%20the%20%22a1b2%22%20format.%20This%20is%20a%20majority%20vote."
}

updateIntent();

function handleCheckmate(isMated: string) {
  if (isMated === 'w') {
    alert('Black wins!');
  } else {
    alert('White wins!');
  }
}

</script>
<template>
  <div class="flex flex-col place-content-center flex-initial justify-start place-items-center w-full mb-10 h-screen">
    <div class="shadow-4xl rounded-lg text-white bg-lgreen text-4xl font-semibold p-3 my-12 m-3">CHESS</div>
    <div>
      <a
        class="flex flex-center"
        :href="intent"
        target="_blank"
        rel="noopener noreferrer"
      >
        <BirdazzoneButton :text="'Ask for some suggestions about a chess match'" :active="true"
          >Ask the scrum for a move</BirdazzoneButton
        >
      </a>
    </div>
    <chessboard />
    <TheChessboard
      :board-config="boardConfig"
      @board-created="(api) => (boardAPI = api)"
      @checkmate="handleCheckmate"
    />
  </div>
</template>
