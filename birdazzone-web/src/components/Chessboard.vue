<script setup lang="ts">
import BirdazzoneButton from '../components/BirdazzoneButton.vue';
import { onBeforeMount, ref } from 'vue';
import { TheChessboard } from 'vue3-chessboard';
import 'vue3-chessboard/style.css';
import type { ChessboardAPI, BoardConfig } from 'vue3-chessboard';
import type { User } from '@/api/interfaces/tweet';
import UserInfo from './UserInfo.vue';
import ApiRepository from '@/api/api-repository';

interface IMove {
    from: string
    to: string
}

const WHITE_TURN = "white"
const BLACK_TURN = "black"
const CHESSBOARD_SIZE = 50

const props = defineProps<{user: User}>()

const boardAPI = ref<ChessboardAPI>();
const boardConfig = ref<BoardConfig>({
  viewOnly: false,
  events: {
    move: (from, to, _) => twitterIntent()
  }
})
const boardLocked = ref<boolean>(false)
const turn = ref<number>(1)
const gameId = ref<string>()

const loadLastMove = async () => {
    if (!gameId.value)
        return
    const response = await ApiRepository.getChessMoves(props.user.username, gameId.value, turn.value.toString())
    if (response.statusCode === 200) {
        doOpponentMove(constructMove(response.data!))
        boardLocked.value = false
    } else if (response.statusCode === 204) {
    } else {
 
    } 
}

const constructMove = (apiMove: string): IMove => ({
    from: apiMove.slice(0,2),
    to: apiMove.slice(2, 4)
}) 

const doOpponentMove = (move: IMove) => {
    boardAPI.value?.makeMove(move.from, move.to)
}

const twitterIntent = () => {
    console.log(boardAPI.value?.board.state.turnColor)
    if (boardAPI.value?.board.state.turnColor === BLACK_TURN) {
      /*window.open(
        "https://twitter.com/intent/tweet?url=https%3A%2F%2Ffen2image.chessvision.ai%2F" +
        encodeURIComponent(
          encodeURIComponent(boardAPI?.value?.board.getFen() ?? 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1')
        ) +
        "&text=Your%20move.%20Please%20use%20the%20%22a1b2%22%20format.%20This%20is%20a%20majority%20vote.",
        "_blank");*/

        console.log("intent aperto")
        boardLocked.value = true
      }

    return 1
}

const handleCheckmate = (isMated: string) => {
  alert((isMated === 'w' ? 'Black' : 'White') + 'wins!');
}

const setGameId = () => {
    gameId.value = Date.now().toString()
}

onBeforeMount(()=>{
    setGameId()
})

</script>
<template>
    <div>
        <!-- HEAD -->
        <div class="p-3 flex">
            <BirdazzoneButton :text="'Check for a new move'" :active="boardLocked" @click="loadLastMove" />
            <UserInfo :user="props.user" :turn="turn" :game-id="gameId"/>
        </div>
        <!-- CONTENT -->
        <div class="flex flex-row place-content-center flex-initial justify-start place-items-center mb-10">
        <TheChessboard :style="`height: ${CHESSBOARD_SIZE}rem; width: ${CHESSBOARD_SIZE}rem;`" class="p-3 relative z-0" :board-config="boardConfig" :after-move-cb="twitterIntent" @board-created="(api) => (boardAPI = api)"
            @checkmate="handleCheckmate" />
        <!-- Lock -->
        <div v-show="boardLocked" :style="`height: ${CHESSBOARD_SIZE}rem; width: ${CHESSBOARD_SIZE}rem;`" class="bg-lgreen flex items-center justify-center absolute z-1">
            <img src="/public/icons/lock.svg" class="w-40" alt="lock" />
        </div>
        </div>
    </div>
</template>