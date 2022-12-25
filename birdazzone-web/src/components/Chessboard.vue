<script setup lang="ts">
  import BirdazzoneSmartButton from '../components/BirdazzoneSmartButton.vue';
  import { onBeforeMount, ref } from 'vue';
  import { TheChessboard } from 'vue3-chessboard';
  import 'vue3-chessboard/style.css';
  import type { ChessboardAPI, BoardConfig } from 'vue3-chessboard';
  import type { User } from '@/api/interfaces/tweet';
  import UserInfo from './UserInfo.vue';
  import ApiRepository from '@/api/api-repository';

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
      move: _ => {
        twitterIntent();
        boardLocked.value = true;
      }
    }
  });
  const boardLocked = ref<boolean>(false);
  const turn = ref<number>(1);
  const gameId = ref<string>();
  const isError = ref<boolean>(false);
  const error = ref<string>('');

  function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  const loadLastMove = async () => {
    if (!gameId.value) return;
    isError.value = false;

    const response = await ApiRepository.getChessMoves(props.user.username, gameId.value, turn.value.toString());
    if (response.statusCode === 200) {
      doOpponentMove(constructMove(response.data!));
      boardLocked.value = false;
    } else if (response.statusCode === 204) {
      isError.value = true;
      error.value = 'Nobody answered yet...be patient';
      await sleep(5000);
      isError.value = false;
    } else {
      isError.value = true;
      error.value = 'You should make a new post first!';
      await sleep(5000);
      isError.value = false;
    }
  };

  const constructMove = (apiMove: string): IMove => ({
    from: apiMove.slice(0, 2),
    to: apiMove.slice(2, 4)
  });

  const doOpponentMove = (move: IMove) => {
    const oldState = boardAPI.value?.board.getFen();
    boardAPI.value?.makeMove(move.from, move.to);
    const newState = boardAPI.value?.board.getFen();
    if (oldState !== newState) {
      turn.value++;
    }
  };

  const onMoveDone = () => {
    if (boardAPI.value?.board.state.turnColor === BLACK_TURN) {
      twitterIntent();
      boardLocked.value = true;
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
  };

  const handleCheckmate = (isMated: string) => {
    alert((isMated === 'w' ? 'Black' : 'White') + 'wins!');
  };

  const setGameId = async () => {
    gameId.value = Date.parse((await ApiRepository.getTwitterTimestamp()).toISOString()).toString();
  };

  onBeforeMount(() => {
    setGameId();
    if (props.startingColor === 'black') {
      twitterIntent();
      boardLocked.value = true;
    }
  });

  const onCheckOpponentClicked = () => {
    if (!boardLocked.value) return;
    loadLastMove();
  };

  const onTweetAgainClicked = () => {
    if (!boardLocked.value) return;
    twitterIntent();
    boardLocked.value = true;
  };
</script>
<template>
  <div class="h-screen flex justify-evenly">
    <!-- HEAD -->
    <div class="ml-10 mt-5 justify-center grid grid-rows-5 grid-flow-col mb-0">
      <div class="mb-6 flex flex-wrap content-center justify-center">
        <UserInfo
          v-if="gameId"
          class="w-1000"
          :user="props.user"
          :turn="turn"
          :game-id="gameId"
        />
      </div>
      <div class="grid grid-cols-2 gap-4 h-10">
        <div class="ml-3">
          <BirdazzoneSmartButton
            :text="'CHECK OPPONENTS'"
            :clickable="boardLocked"
            @click="onCheckOpponentClicked"
          />
        </div>
        <div class="ml-3">
          <BirdazzoneSmartButton
            :text="'TWEET AGAIN'"
            :clickable="boardLocked"
            @click="onTweetAgainClicked"
          />
        </div>
      </div>
      <div
        v-show="error"
        class="flex-col justify-center items-center m-3"
      >
        <div class="flex justify-center items-center">
          <img
            style="height: 12rem; widows: 12rem"
            src="/bchess.png"
            alt="A bird playing chess."
          />
        </div>
        <div class="animate-pulse rounded-2xl p-3">
          <h1 class="text-lred text-center text-xl font-bold">{{ error }}</h1>
        </div>
      </div>
    </div>
    <!-- CONTENT -->
    <div class="flex justify-center items-center mt-0">
      <TheChessboard
        :style="`height: ${CHESSBOARD_SIZE}rem; width: ${CHESSBOARD_SIZE}rem;`"
        class="relative z-0"
        :board-config="boardConfig"
        :after-move-cb="onMoveDone"
        @board-created="api => (boardAPI = api)"
        @checkmate="handleCheckmate"
      />
      <!-- Lock -->
      <div
        v-if="boardLocked"
        :style="`height: ${CHESSBOARD_SIZE}rem; width: ${CHESSBOARD_SIZE}rem;`"
        class="flex absolute z-1 items-center justify-center"
      >
        <div
          :style="`background-color: #1eb980aa`"
          class="flex items-center justify-center absolute z-1 p-20 rounded-full"
        >
          <img
            src="/icons/lock.svg"
            class="w-40"
            alt="lock"
          />
        </div>
      </div>
    </div>
  </div>
</template>
