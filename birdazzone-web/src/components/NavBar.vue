<script setup lang="ts">
  import { ref } from 'vue';
  import { GAMES_PATH, CHESS_PATH, FANTA_PATH } from '../router/index';

  interface NavItem {
    title: string;
    url: string;
    icon: string;
    active: boolean;
  }

  const items = ref<NavItem[]>([
    {
      url: GAMES_PATH,
      title: 'TV GAMES',
      active: true,
      icon: '/tv.svg'
    },
    {
      url: CHESS_PATH,
      title: 'CHESS',
      active: false,
      icon: '/chess.svg'
    },
    {
      url: FANTA_PATH,
      title: 'FANTACITORIO',
      active: false,
      icon: '/montecitorio.svg'
    }
  ]);

  const select = (url: string) => items.value.forEach(x => (x.active = x.url === url));
</script>
<template>
  <aside
    class="w-64 min-w-5xl max-w-5xl h-100 bg-foreground fixed-left"
    aria-label="Sidebar"
  >
    <!-- Title -->
    <div class="py-10 flex flex-col items-center">
      <img
        src="/logo.png"
        alt="Birdazzone's logo"
      />
      <h1 class="text-lgreen font-black text-xl">IL BIRDAZZONE</h1>
    </div>
    <div class="overflow-y-auto py-4 px-3 rounded">
      <ul class="my-5 space-y-2">
        <li>
          <div
            class="my-3"
            v-for="item in items"
          >
            <RouterLink :to="item.url">
              <div
                @click="select(item.url)"
                class="flex text-white items-center p-2 font-normal rounded-lg"
              >
                <div
                  v-if="item.active"
                  class="flex items-center"
                >
                  <img
                    :src="item.icon"
                    alt="An icon for the current category"
                    class="white-filter icon"
                  />
                  <h2 class="ml-3 font-bold text-white">{{ item.title }}</h2>
                </div>
                <div
                  v-else
                  class="flex items-center"
                >
                  <img
                    :src="item.icon"
                    alt="An icon for one of the others categories"
                    class="gray-filter icon"
                  />
                  <h2 class="ml-3 font-bold text-lgray">{{ item.title }}</h2>
                </div>
              </div>
            </RouterLink>
          </div>
        </li>
      </ul>
    </div>
  </aside>
</template>
<style>
  .icon {
    height: 2rem;
    width: 2rem;
  }
  .gray-filter {
    filter: invert(48%) sepia(79%) saturate(0%) hue-rotate(86deg) brightness(118%) contrast(10%);
  }
  .white-filter {
    filter: invert(48%) saturate(100%) brightness(200%) contrast(200%);
  }
</style>
