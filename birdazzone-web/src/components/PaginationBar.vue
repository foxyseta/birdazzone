<script setup lang="ts">
  import { ref } from 'vue';
  const props = defineProps<{ actualPage: number; max: number }>();
</script>

<template>
  <div
    v-if="actualPage >= 1 && actualPage <= max"
    class="flex flex-row justify-center"
    style="max-width: 30rem"
  >
    <div
      v-if="actualPage === 1"
      @click="$emit('changeActual', 1)"
      class="mx-2 p-1 text-lgreen text-center bg-lgreen text-foreground"
      style="height: 2.3rem; width: 2.3rem; border-radius: 50%; border: solid 2px"
    >
      1
    </div>
    <div
      v-else
      id="firstIndex"
      @click="$emit('changeActual', 1)"
      class="mx-2 p-1 text-lgreen text-center hover:bg-lgreen hover:text-foreground"
      style="height: 2.3rem; width: 2.3rem; border-radius: 50%; border: solid 2px"
    >
      1
    </div>

    <div
      v-if="actualPage > 1"
      class="flex flex-row"
    >
      <div v-if="actualPage === 2">
        <div
          v-if="props.max > 2"
          @click="$emit('changeActual', actualPage)"
          class="mx-2 p-1 text-lgreen text-center bg-lgreen text-foreground"
          style="height: 2.3rem; width: 2.3rem; border-radius: 50%; border: solid 2px"
        >
          {{ actualPage }}
        </div>
      </div>
      <div
        v-else
        class="flex flex-row"
      >
        <div
          v-if="actualPage !== 3"
          class="text-lgray mx-2"
          >...</div
        >
        <div
          @click="$emit('changeActual', actualPage - 1)"
          class="mx-2 p-1 text-lgreen text-center hover:bg-lgreen hover:text-foreground"
          style="flex: 1 1 auto; height: 2.3rem; width: 2.3rem; border-radius: 50%; border: solid 2px"
        >
          {{ actualPage - 1 }}
        </div>
        <div
          v-if="actualPage !== max"
          @click="$emit('changeActual', actualPage)"
          class="mx-2 p-1 text-lgreen text-center bg-lgreen text-foreground"
          style="flex: 1 1 auto; height: 2.3rem; width: 2.3rem; border-radius: 50%; border: solid 2px"
        >
          {{ actualPage }}
        </div>
      </div>
    </div>
    <div
      v-if="props.max > actualPage + 1"
      @click="$emit('changeActual', actualPage + 1)"
      class="mx-2 p-1 text-lgreen text-center hover:bg-lgreen hover:text-foreground"
      style="height: 2.3rem; width: 2.3rem; border-radius: 50%; border: solid 2px"
    >
      {{ actualPage + 1 }}
    </div>

    <div
      v-if="actualPage + 2 < max"
      class="text-lgray mx-2"
      >...</div
    >
    <div v-if="props.max !== 1">
      <div
        v-if="actualPage === props.max"
        @click="$emit('changeActual', actualPage)"
        class="mx-2 p-1 text-lgreen text-center bg-lgreen text-foreground"
        style="height: 2.3rem; width: 2.3rem; border-radius: 50%; border: solid 2px"
      >
        {{ props.max }}
      </div>
      <div
        v-else
        class="mx-2 p-1 text-lgreen text-center hover:bg-lgreen hover:text-foreground"
        @click="$emit('changeActual', props.max)"
        style="height: 2.3rem; width: 2.3rem; border-radius: 50%; border: solid 2px"
      >
        {{ props.max }}
      </div>
    </div>
  </div>
</template>
