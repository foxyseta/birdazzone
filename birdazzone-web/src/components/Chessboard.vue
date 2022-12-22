<script setup lang="ts">
import BirdazzoneButton from '../components/BirdazzoneButton.vue';
import { onBeforeMount, ref } from 'vue';
import { TheChessboard } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';
import type { ChessboardAPI, BoardConfig } from 'vue3-chessboard';
import type { User } from '@/api/interfaces/tweet';
import UserInfo from './UserInfo.vue';
import ApiRepository from '@/api/api-repository';
import { ErrorTypes } from 'vue-router';

export type ChessColor = 'white' | 'black';

interface IMove {
  from: string;
  to: string;
}

const WHITE_TURN = 'white';
const BLACK_TURN = 'black';
const CHESSBOARD_SIZE = 50;

const props = defineProps<{ user: User; startingColor: ChessColor }>();

const boardAPI = ref<ChessboardAPI>();
const boardConfig = ref<BoardConfig>({
  viewOnly: false,
  orientation: props.startingColor,
  events: {
    move: (_) => twitterIntent(),
  },
});
const boardLocked = ref<boolean>(false);
const turn = ref<number>(1);
const gameId = ref<string>();
const isError = ref<boolean>(false);
const error = ref<string>('');

const loadLastMove = async () => {
  if (!gameId.value) return;
  const response = await ApiRepository.getChessMoves(props.user.username, gameId.value, turn.value.toString());
  if (response.statusCode === 200) {
    isError.value = false;
    doOpponentMove(constructMove(response.data!));
    boardLocked.value = false;
  } else if (response.statusCode === 204) {
    isError.value = true;
    error.value = 'Nobody answered yet...be patient';
  } else {
    isError.value = true;
    error.value = 'You should make a new post first!';
  }
};

const constructMove = (apiMove: string): IMove => ({
  from: apiMove.slice(0, 2),
  to: apiMove.slice(2, 4),
});

const doOpponentMove = (move: IMove) => {
  // TODO check correct moves
  boardAPI.value?.makeMove(move.from, move.to);
  turn.value++;
};

const onMoveDone = () => {
  if (boardAPI.value?.board.state.turnColor === BLACK_TURN) {
    twitterIntent();
  }
  return 1;
};

const twitterIntent = () => {
  window.open(
    'https://twitter.com/intent/tweet?url=https%3A%2F%2Ffen2image.chessvision.ai%2F' +
      encodeURIComponent(
        encodeURIComponent(
          boardAPI?.value?.board.getFen() ?? 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1'
        )
      ) +
      '&text=Your%20move.%20Please%20retweet%20using%20the%20%22a1b2%22%20format.%20This%20is%20a%20majority%20vote.',
    '_blank'
  );

  console.log('intent aperto');
  boardLocked.value = true;
};

const handleCheckmate = (isMated: string) => {
  alert((isMated === 'w' ? 'Black' : 'White') + 'wins!');
};

const setGameId = async () => {
  gameId.value = Date.parse((await ApiRepository.getTwitterTimestamp()).toISOString()).toString()
};

onBeforeMount(() => {
  setGameId();
  if (props.startingColor === 'black') {
    console.log('dakdaskds');
    twitterIntent();
  }
});
</script>
<template>
  <div>
    <!-- HEAD -->
    <div class="flex mb-3">
      <UserInfo class="w-full" :user="props.user" :turn="turn" :game-id="gameId" />
      <BirdazzoneButton class="ml-3" :text="'CHECK OPPONENTS'" :active="boardLocked" @click="loadLastMove" />
      <BirdazzoneButton class="ml-3" :text="'TWEET AGAIN'" :active="boardLocked" @click="twitterIntent" />
    </div>
    <!-- CONTENT -->
    <div v-show="isError">
      <h1 class="text-lred font-bold">{{ error }}</h1>
    </div>
    <div class="w-full flex justify-center items-center m-15">
      <TheChessboard
        :style="`height: ${CHESSBOARD_SIZE}rem; width: ${CHESSBOARD_SIZE}rem;`"
        class="relative z-0"
        :board-config="boardConfig"
        :after-move-cb="onMoveDone"
        @board-created="(api) => (boardAPI = api)"
        @checkmate="handleCheckmate"
      />
      <!-- Lock -->
      <div
        v-if="boardLocked"
        :style="`background-color: #1eb980aa`"
        class="flex items-center justify-center absolute z-1 p-20 rounded-full"
      >
        <img src="/icons/lock.svg" class="w-40" alt="lock" />
      </div>
    </div>
  </div>
</template>
