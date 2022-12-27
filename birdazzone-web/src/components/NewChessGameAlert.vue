<script lang="ts" setup>
  import {ref} from 'vue';
  import ApiRepository from '@/api/api-repository';
  import type {ChessColor} from './Chessboard.vue';
  import BirdazzoneButton from './BirdazzoneButton.vue';
  import type {User} from '@/api/interfaces/tweet';

  const emit = defineEmits(['done']);
  const username = ref<string>('');
  const user = ref<User>();
  const error = ref<string>('');
  const isError = ref<boolean>(false);
  const color = ref<ChessColor>('white');

  const checkUsername = async () => {
    if (!username.value) return false;
    const response = await ApiRepository.getHelloUser(username.value);

    isError.value = !response.esit;

    if (response.statusCode === 200) {
      user.value = response.data!;
      return true;
    }
    error.value = 'User not found :/ Try again!';
    return false;
  };

  const selectWhite = () => {
    color.value = 'white';
  };
  const selectBlack = () => {
    color.value = 'black';
  };

  const verifyAndConfirm = async () => {
    if (await checkUsername()) emit('done', user.value, color.value);
  };
</script>
<template>
  <div
    class="shadow bg-foreground rounded-2xl flex justify-start items-center my-5 mx-20 p-10 flex-col font-semibold"
  >
    <h1 class="font-bold text-4xl text-white"
      >Welcome to the Birdazzone chess game!</h1
    >
    <div class="m-10 flex flex-col items-center justify-evenly">
      <div
        v-show="isError"
        class="p-2 text-center text-lred rounded-2xl border-2 border-lred"
      >
        {{ error }}
      </div>

      <p class="m-3 text-white text-xl text-center">Tell me what's your name</p>
      <input
        v-model="username"
        class="form-control block px-3 py-1.5 text-base font-normal text-white bg-foreground bg-clip-padding border-2 border-solid border-dgreen rounded-xl transition ease-in-out m-0 focus:border-lgreen focus:outline-none"
        type="text"
      />

      <p class="mt-10 text-white text-xl text-center"
        >And which color do you like more</p
      >

      <div class="flex w-full justify-evenly m-5">
        <div>
          <button
            v-if="color !== 'black'"
            @click="selectBlack()"
            class="bg-lblack p-10 text-white rounded-lg"
          >
            black
          </button>
          <button
            v-if="color === 'black'"
            @click="selectBlack()"
            class="bg-lblack p-10 text-white border-lgreen border-2 rounded-lg"
          >
            black
          </button>
        </div>
        <div>
          <button
            v-if="color !== 'white'"
            @click="selectWhite()"
            class="bg-white p-10 text-lblack rounded-lg"
          >
            white
          </button>
          <button
            v-if="color === 'white'"
            @click="selectWhite()"
            class="bg-white p-10 text-lblack border-2 border-lgreen rounded-lg"
          >
            white
          </button>
        </div>
      </div>

      <BirdazzoneButton
        @click="verifyAndConfirm"
        class="h-20 w-40 mt-10 border-2"
        :active="username !== ''"
        :text="'Done!'"
      />
    </div>
  </div>
</template>
